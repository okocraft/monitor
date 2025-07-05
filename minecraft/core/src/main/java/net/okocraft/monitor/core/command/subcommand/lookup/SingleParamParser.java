package net.okocraft.monitor.core.command.subcommand.lookup;

import dev.siroshun.jfun.result.Result;
import net.okocraft.monitor.core.command.ArgumentList;

import java.time.Duration;
import java.time.LocalDateTime;

public interface SingleParamParser {

    SingleParamParser START_PARSER = (context, args) -> {
        if (context.has(ParamContext.KEY_START)) {
            return Result.failure(new ParamParseError("start datetime is already specified"));
        }

        if (!args.hasNext()) {
            return Result.failure(new ParamParseError("--start is specified, but no datetime supplied"));
        }

        Result<LocalDateTime, ParamParseError> result = ParseUtils.parseAsDateTime(args.next(), false);
        if (result.isFailure()) {
            return result.asFailure();
        }

        context.put(ParamContext.KEY_START, result.unwrap());
        return Result.success(null);
    };

    SingleParamParser END_PARSER = (context, args) -> {
        if (context.has(ParamContext.KEY_END)) {
            return Result.failure(new ParamParseError("end datetime is already specified"));
        }

        if (!args.hasNext()) {
            return Result.failure(new ParamParseError("--end is specified, but no datetime supplied"));
        }

        Result<LocalDateTime, ParamParseError> result = ParseUtils.parseAsDateTime(args.next(), true);
        if (result.isFailure()) {
            return result.asFailure();
        }

        context.put(ParamContext.KEY_END, result.unwrap());
        return Result.success(null);
    };

    SingleParamParser DURATION_PARSER = (context, args) -> {
        if (!args.hasNext()) {
            return Result.failure(new ParamParseError("--duration is specified, but no duration supplied"));
        }

        Result<Duration, ParamParseError> result = ParseUtils.parseAsDuration(args.next());
        if (result.isFailure()) {
            return result.asFailure();
        }

        LocalDateTime start = context.get(ParamContext.KEY_START);
        LocalDateTime end = context.get(ParamContext.KEY_END);

        if (start == null && end == null) {
            context.put(ParamContext.KEY_START, LocalDateTime.now().minus(result.unwrap()));
            context.put(ParamContext.KEY_END, LocalDateTime.now());
        } else if (start == null) {
            context.put(ParamContext.KEY_START, end.minus(result.unwrap()));
        } else if (end == null) {
            context.put(ParamContext.KEY_END, start.plus(result.unwrap()));
        } else {
            return Result.failure(new ParamParseError("both start and end datetime are specified, or duration is already specified"));
        }

        return Result.success(null);
    };

    SingleParamParser PAGE_PARSER = (context, args) -> {
        if (context.has(ParamContext.KEY_PAGE)) {
            return Result.failure(new ParamParseError("page number is already specified"));
        }

        if (!args.hasNext()) {
            return Result.failure(new ParamParseError("--page is specified, but no page number supplied"));
        }

        long page;
        try {
            page = Long.parseLong(args.next());
        } catch (NumberFormatException e) {
            return Result.failure(new ParamParseError("invalid page number: " + args.next()));
        }

        if (page < 1) {
            return Result.failure(new ParamParseError("page number must be greater than or equal to 1"));
        }

        context.put(ParamContext.KEY_PAGE, page);
        return Result.success(null);
    };

    Result<Void, ParamParseError> parse(ParamContext context, ArgumentList args);

}
