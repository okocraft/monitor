package net.okocraft.monitor.core.command.subcommand;

import dev.siroshun.jfun.result.Result;
import net.okocraft.monitor.core.command.ArgumentList;
import net.okocraft.monitor.core.command.Command;
import net.okocraft.monitor.core.command.subcommand.lookup.LogLookup;
import net.okocraft.monitor.core.command.subcommand.lookup.ParamContext;
import net.okocraft.monitor.core.command.subcommand.lookup.ParamParseError;
import net.okocraft.monitor.core.command.subcommand.lookup.ParamParser;
import net.okocraft.monitor.core.models.lookup.PlayerChatLogLookupParams;
import net.okocraft.monitor.core.models.data.PlayerChatLogData;
import net.okocraft.monitor.core.models.data.PlayerConnectLogData;
import net.okocraft.monitor.core.models.lookup.CommonLookupParams;
import net.okocraft.monitor.core.models.lookup.PlayerConnectLogLookupParams;
import net.okocraft.monitor.core.storage.PlayerLogStorage;

import java.time.LocalDateTime;
import java.util.Arrays;

public abstract class AbstractLookupCommand implements Command {

    protected final LogLookup<PlayerConnectLogLookupParams, PlayerConnectLogData> connectLogLookup;
    protected final LogLookup<PlayerChatLogLookupParams, PlayerChatLogData> chatLogLookup;

    protected AbstractLookupCommand(PlayerLogStorage logStorage, int logsPerPage) {
        this.connectLogLookup = new LogLookup<>(
            "connect",
            ParamParser.builder().withStartAndEnd().withDuration().withPage().build(context -> {
                Result<CommonLookupParams.Record, ParamParseError> common = createCommonParamsFromContext(context, logsPerPage);
                if (common.isFailure()) {
                    return Result.failure(common.unwrapError());
                }

                return Result.success(new PlayerConnectLogLookupParams(
                    common.unwrap().start(), common.unwrap().end(), common.unwrap().limit(), common.unwrap().offset()
                ));
            }),
            logStorage::countConnectLogs,
            logStorage::lookupConnectLogData
        );

        this.chatLogLookup = new LogLookup<>(
            "chat",
            ParamParser.builder().withStartAndEnd().withDuration().withPage().build(context -> {
                Result<CommonLookupParams.Record, ParamParseError> common = createCommonParamsFromContext(context, logsPerPage);
                if (common.isFailure()) {
                    return Result.failure(common.unwrapError());
                }

                return Result.success(new PlayerChatLogLookupParams(
                    common.unwrap().start(), common.unwrap().end(), common.unwrap().limit(), common.unwrap().offset()
                ));
            }),
            logStorage::countChatLogs,
            logStorage::lookupChatLogData
        );
    }

    protected <P extends CommonLookupParams> Result<P, ParamParseError> parseAsLookupParams(String[] args, LogLookup<P, ?> lookup) {
        ArgumentList argList = new ArgumentList(Arrays.copyOfRange(args, 2, args.length));
        return lookup.parser().parse(argList);
    }

    private static Result<CommonLookupParams.Record, ParamParseError> createCommonParamsFromContext(ParamContext context, long logsPerPage) {
        LocalDateTime start = context.get(ParamContext.KEY_START);
        LocalDateTime end = context.get(ParamContext.KEY_END);
        long page = context.getOrDefault(ParamContext.KEY_PAGE, 1L);

        if (start == null && end == null) {
            start = LocalDateTime.now().minusDays(1);
            end = LocalDateTime.now();
        } else if (start == null) {
            start = end.minusDays(1);
        } else if (end == null) {
            end = LocalDateTime.now();
        } else if (start.isAfter(end)) {
            return Result.failure(new ParamParseError("start datetime is after end datetime"));
        }

        long offset;
        try {
            offset = Math.max(0, Math.multiplyExact(page - 1, logsPerPage));
        } catch (ArithmeticException e) {
            return Result.failure(new ParamParseError("page number is too large"));
        }

        return Result.success(new CommonLookupParams.Record(start, end, logsPerPage, offset));
    }
}
