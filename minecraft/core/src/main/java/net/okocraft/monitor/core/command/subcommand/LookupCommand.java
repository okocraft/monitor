package net.okocraft.monitor.core.command.subcommand;

import net.okocraft.monitor.core.command.Command;
import net.okocraft.monitor.core.command.CommandSender;
import net.okocraft.monitor.core.logger.MonitorLogger;
import net.okocraft.monitor.core.storage.PlayerLogStorage;

import java.sql.SQLException;

public class LookupCommand extends AbstractLookupCommand implements Command {

    public LookupCommand(PlayerLogStorage storage) {
        super(storage);
    }

    @Override
    public void execute(CommandSender sender, String[] args) {
        sender.sendPlainMessage("Lookup connect log...");
        try {
            this.lookupConnectLog(log -> sender.sendPlainMessage(log.toString()));
        } catch (SQLException e) {
            sender.sendPlainMessage("Failed to lookup connect log: " + e.getMessage());
            MonitorLogger.logger().error("Failed to lookup connect log", e);
        }
        sender.sendPlainMessage("Done.");
    }
}
