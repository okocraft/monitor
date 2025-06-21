package net.okocraft.monitor.platform.velocity.listener;

import com.velocitypowered.api.command.CommandManager;
import com.velocitypowered.api.command.CommandSource;
import com.velocitypowered.api.event.Subscribe;
import com.velocitypowered.api.event.command.CommandExecuteEvent;
import com.velocitypowered.api.proxy.ConsoleCommandSource;
import com.velocitypowered.api.proxy.Player;
import net.okocraft.monitor.core.handler.PlayerHandler;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class CommandListener {

    private static final Logger COMMAND_EXECUTION_LOGGER = LoggerFactory.getLogger("command execution");
    private final PlayerHandler handler;
    private final CommandManager commandManager;

    public CommandListener(PlayerHandler handler, CommandManager commandManager) {
        this.handler = handler;
        this.commandManager = commandManager;
    }

    @Subscribe(priority = Short.MIN_VALUE) // last
    public void onCommand(CommandExecuteEvent event) {
        if (!event.getResult().isAllowed() || event.getResult().isForwardToServer()) {
            return;
        }

        CommandSource source = event.getCommandSource();
        String senderName;
        String commandline = event.getCommand();

        if (source instanceof Player player) {
            senderName = player.getUsername();

            int firstSpace = commandline.indexOf(' ');
            String label = commandline.substring(0, firstSpace != -1 ? firstSpace : commandline.length());

            if (!this.commandManager.hasCommand(label)) {
                return;
            }

            this.handler.onProxyCommand(player.getUniqueId(), event.getCommand());
        } else if (source instanceof ConsoleCommandSource) {
            senderName = "Console";
        } else {
            senderName = source.toString();
        }

        COMMAND_EXECUTION_LOGGER.info("{} issued proxy command: /{}", senderName, commandline);
    }
}
