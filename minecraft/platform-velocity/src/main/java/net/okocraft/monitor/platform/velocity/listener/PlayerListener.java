package net.okocraft.monitor.platform.velocity.listener;

import com.velocitypowered.api.event.Subscribe;
import com.velocitypowered.api.event.command.CommandExecuteEvent;
import com.velocitypowered.api.event.connection.DisconnectEvent;
import com.velocitypowered.api.event.connection.PostLoginEvent;
import com.velocitypowered.api.event.player.KickedFromServerEvent;
import com.velocitypowered.api.proxy.Player;
import com.velocitypowered.api.proxy.ProxyServer;
import net.kyori.adventure.text.Component;
import net.okocraft.monitor.core.handler.PlayerHandler;
import net.okocraft.monitor.core.models.logs.PlayerConnectLog;

public class PlayerListener {

    private final ProxyServer server;
    private final PlayerHandler handler;

    public PlayerListener(ProxyServer server, PlayerHandler handler) {
        this.server = server;
        this.handler = handler;
    }

    @Subscribe
    public void onJoin(PostLoginEvent event) {
        Player player = event.getPlayer();
        this.handler.onJoin(player.getUniqueId(), player.getUsername(), player.getRemoteAddress());
    }

    @Subscribe
    public void onKick(KickedFromServerEvent event) {
        Player player = event.getPlayer();
        this.handler.onLeave(player.getUniqueId(), PlayerConnectLog.Action.KICKED, player.getRemoteAddress(), event.getServerKickReason().orElse(Component.empty()));
    }

    @Subscribe
    public void onQuit(DisconnectEvent event) {
        Player player = event.getPlayer();
        this.handler.onLeave(player.getUniqueId(), PlayerConnectLog.Action.DISCONNECT, player.getRemoteAddress(), Component.empty());
    }

    @Subscribe
    public void onCommand(CommandExecuteEvent event) {
        if (!(event.getCommandSource() instanceof Player player)) {
            return;
        }

        var command = event.getCommand();
        int firstSpace = command.indexOf(' ');
        var label = firstSpace == -1 ? command : command.substring(0, firstSpace);
        if (this.server.getCommandManager().hasCommand(label)) {
            this.handler.onProxyCommand(player.getUniqueId(), event.getCommand());
        }
    }
}
