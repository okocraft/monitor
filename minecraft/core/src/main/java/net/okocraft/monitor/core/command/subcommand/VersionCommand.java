package net.okocraft.monitor.core.command.subcommand;

import net.okocraft.monitor.core.command.CommandSender;

public class VersionCommand {

    public static void execute(CommandSender sender, String version) {
        sender.sendPlainMessage("Monitor version: " + version);
    }

}
