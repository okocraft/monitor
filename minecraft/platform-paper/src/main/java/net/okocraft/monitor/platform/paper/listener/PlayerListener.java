package net.okocraft.monitor.platform.paper.listener;

import io.papermc.paper.event.player.AsyncChatEvent;
import net.kyori.adventure.text.Component;
import net.okocraft.monitor.core.handler.PlayerHandler;
import net.okocraft.monitor.core.models.logs.PlayerConnectLog;
import net.okocraft.monitor.platform.paper.adapter.PositionAdapter;
import org.bukkit.entity.Player;
import org.bukkit.event.EventHandler;
import org.bukkit.event.EventPriority;
import org.bukkit.event.Listener;
import org.bukkit.event.player.PlayerCommandPreprocessEvent;
import org.bukkit.event.player.PlayerJoinEvent;
import org.bukkit.event.player.PlayerKickEvent;
import org.bukkit.event.player.PlayerQuitEvent;

public class PlayerListener implements Listener {

    private final PlayerHandler handler;

    public PlayerListener(PlayerHandler handler) {
        this.handler = handler;
    }

    @EventHandler
    public void onJoin(PlayerJoinEvent event) {
        Player player = event.getPlayer();
        this.handler.onJoin(player.getUniqueId(), player.getName(), player.getAddress());
    }

    @EventHandler
    public void onKick(PlayerKickEvent event) {
        Player player = event.getPlayer();
        this.handler.onLeave(player.getUniqueId(), PlayerConnectLog.Action.KICKED, player.getAddress(), event.reason());
    }

    @EventHandler
    public void onQuit(PlayerQuitEvent event) {
        if (event.getReason() == PlayerQuitEvent.QuitReason.KICKED) {
            return;
        }

        Player player = event.getPlayer();
        PlayerConnectLog.Action action = switch (event.getReason()) {
            case DISCONNECTED -> PlayerConnectLog.Action.DISCONNECT;
            case KICKED -> throw new AssertionError();
            case TIMED_OUT -> PlayerConnectLog.Action.TIMED_OUT;
            case ERRONEOUS_STATE -> PlayerConnectLog.Action.ERRONEOUS_STATE;
        };
        this.handler.onLeave(player.getUniqueId(), action, player.getAddress(), Component.empty());
    }

    @EventHandler(priority = EventPriority.MONITOR)
    public void onChat(AsyncChatEvent event) {
        Player player = event.getPlayer();
        this.handler.onChat(player.getUniqueId(), player.getWorld().getUID(), PositionAdapter.fromLocation(player.getLocation()), event.originalMessage());
    }

    @EventHandler(priority = EventPriority.MONITOR)
    public void onCommand(PlayerCommandPreprocessEvent event) {
        Player player = event.getPlayer();
        this.handler.onWorldCommand(player.getUniqueId(), player.getWorld().getUID(), PositionAdapter.fromLocation(player.getLocation()), event.getMessage());
    }
}
