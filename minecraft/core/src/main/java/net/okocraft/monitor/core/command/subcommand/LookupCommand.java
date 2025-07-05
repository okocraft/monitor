package net.okocraft.monitor.core.command.subcommand;

import dev.siroshun.jfun.result.Result;
import net.okocraft.monitor.core.command.CommandSender;
import net.okocraft.monitor.core.command.subcommand.lookup.LogLookup;
import net.okocraft.monitor.core.command.subcommand.lookup.ParamParseError;
import net.okocraft.monitor.core.logger.MonitorLogger;
import net.okocraft.monitor.core.models.lookup.CommonLookupParams;
import net.okocraft.monitor.core.storage.PlayerLogStorage;

import java.sql.SQLException;
import java.time.format.DateTimeFormatter;

public class LookupCommand extends AbstractLookupCommand {

    public LookupCommand(PlayerLogStorage storage) {
        super(storage, 15);
    }

    @Override
    public void execute(CommandSender sender, String[] args) {
        if (args.length < 2) {
            sender.sendPlainMessage("Usage: /monitor lookup <type> {params}");
            return;
        }

        switch (args[1].toLowerCase()) {
            case "connect" -> this.lookupAndSendResult(sender, args, this.connectLogLookup);
            case "chat" -> this.lookupAndSendResult(sender, args, this.chatLogLookup);
            default -> sender.sendPlainMessage("Unknown type: " + args[1]);
        }
    }

    private <P extends CommonLookupParams, T> void lookupAndSendResult(CommandSender sender, String[] args, LogLookup<P, T> lookup) {
        Result<P, ParamParseError> paramResult = this.parseAsLookupParams(args, lookup);
        if (paramResult.isFailure()) {
            sender.sendPlainMessage(paramResult.unwrapError().message());
            return;
        }

        P params = paramResult.unwrap();
        try {
            long count = lookup.countByParam().count(params);
            if (count == 0) {
                sender.sendPlainMessage("No " + lookup.type() + " logs found.");
                return;
            }

            long start = params.offset() + 1;
            long end = Math.min(start + params.limit() - 1, count);

            sender.sendPlainMessage("Showing " + start + "-" + end + " of " + count + " logs.");
            sender.sendPlainMessage("(" + DateTimeFormatter.ISO_LOCAL_DATE_TIME.format(params.start()) + " ~ " + DateTimeFormatter.ISO_LOCAL_DATE_TIME.format(params.end()) + ")");
            lookup.lookupByParam().lookup(params, log -> sender.sendPlainMessage(log.toString()));
        } catch (SQLException e) {
            sender.sendPlainMessage("Failed to lookup " + lookup.type() + " log: " + e.getMessage());
            MonitorLogger.logger().error("Failed to lookup {} log", lookup.type(), e);
        }
    }
}
