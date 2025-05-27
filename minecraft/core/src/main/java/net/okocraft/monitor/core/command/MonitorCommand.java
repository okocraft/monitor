package net.okocraft.monitor.core.command;

import net.okocraft.monitor.core.command.subcommand.LookupCommand;
import net.okocraft.monitor.core.command.subcommand.VersionCommand;
import net.okocraft.monitor.core.storage.Storage;
import org.jetbrains.annotations.NotNullByDefault;

import java.util.Collections;
import java.util.List;
import java.util.concurrent.CompletableFuture;

@NotNullByDefault
public class MonitorCommand implements Command {

    public static final String LABEL = "monitor";
    public static final String PERMISSION = "monitor.command";

    private final SubCommandMap subCommandMap;

    public MonitorCommand(String pluginVersion, Storage storage) {
        this.subCommandMap = SubCommandMap.builder()
                .add("version", PERMISSION + ".version", (sender, ignored) -> VersionCommand.execute(sender, pluginVersion), "ver", "v")
                .add("lookup", PERMISSION + ".lookup", new LookupCommand(storage.getPlayerLogStorage()))
                .build();
    }

    @Override
    public void execute(CommandSender sender, String[] args) {
        if (!sender.hasPermission(PERMISSION)) {
            sender.sendPlainMessage("You don't have permission to execute this command.");
            return;
        }

        if (args.length == 0) {
            sender.sendPlainMessage("Usage: /" + LABEL + " <subcommand>");
            return;
        }

        Command command = this.subCommandMap.findWithPermission(args[0], sender);
        if (command == null) {
            sender.sendPlainMessage("Unknown subcommand: " + args[0]);
            return;
        }

        command.execute(sender, args);
    }

    @Override
    public CompletableFuture<List<String>> tabComplete(CommandSender sender, String[] args) {
        if (!sender.hasPermission(PERMISSION) || args.length == 0) {
            return CompletableFuture.completedFuture(Collections.emptyList());
        }

        if (args.length == 1) {
            return CompletableFuture.completedFuture(this.subCommandMap.findLabelsWithPermission(args[0], sender));
        }

        Command command = this.subCommandMap.findWithPermission(args[0], sender);
        if (command != null) {
            return command.tabComplete(sender, args);
        }

        return CompletableFuture.completedFuture(Collections.emptyList());
    }
}
