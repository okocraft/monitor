package net.okocraft.monitor.core.command.subcommand.lookup;

import dev.siroshun.jfun.result.Result;
import net.okocraft.monitor.core.command.ArgumentList;

import java.util.HashMap;
import java.util.Map;
import java.util.function.Function;

public final class ParamParser<P> {

    public static <P> Builder builder() {
        return new Builder();
    }

    private final Map<String, SingleParamParser> parsers;
    private final Function<ParamContext, Result<P, ParamParseError>> build;

    private ParamParser(Map<String, SingleParamParser> parsers, Function<ParamContext, Result<P, ParamParseError>> build) {
        this.parsers = parsers;
        this.build = build;
    }

    public Result<P, ParamParseError> parse(ArgumentList args) {
        ParamContext context = ParamContext.create();
        while (args.hasNext()) {
            String key = args.next();
            SingleParamParser parser = this.parsers.get(key);
            if (parser == null) {
                return Result.failure(new ParamParseError("unknown parameter: " + key));
            }
            Result<Void, ParamParseError> result = parser.parse(context, args);
            if (result.isFailure()) {
                return result.asFailure();
            }
        }

        return this.build.apply(context);
    }

    public static final class Builder {

        private final Map<String, SingleParamParser> parsers = new HashMap<>();

        public Builder param(SingleParamParser parser, String... flags) {
            if (flags == null || flags.length == 0) {
                throw new IllegalArgumentException("flags must not be null or empty");
            }

            for (String flag : flags) {
                if (this.parsers.containsKey(flag)) {
                    throw new IllegalArgumentException("flag is already registered: " + flag);
                }
                this.parsers.put(flag, parser);
            }

            return this;
        }

        public Builder withStartAndEnd() {
            return this.param(SingleParamParser.START_PARSER, "-s", "--start").param(SingleParamParser.END_PARSER, "-e", "--end");
        }

        public Builder withDuration() {
            return this.param(SingleParamParser.DURATION_PARSER, "-d", "--duration");
        }

        public Builder withPage() {
            return this.param(SingleParamParser.PAGE_PARSER, "-p", "--page");
        }

        public <P> ParamParser<P> build(Function<ParamContext, Result<P, ParamParseError>> build) {
            if (build == null) {
                throw new IllegalArgumentException("build must not be null");
            }

            return new ParamParser<>(Map.copyOf(this.parsers), build);
        }
    }
}
