package net.okocraft.monitor.platform.paper.listener;

import net.okocraft.monitor.core.handler.PlayerHandler;
import org.bukkit.entity.Player;
import org.bukkit.event.EventHandler;
import org.bukkit.event.Listener;
import org.bukkit.event.player.PlayerJoinEvent;

public class PlayerListener implements Listener {

    private final PlayerHandler handler;

    public PlayerListener(PlayerHandler handler) {
        this.handler = handler;
    }

    @EventHandler
    public void onJoin(PlayerJoinEvent event) {
        Player player = event.getPlayer();
        this.handler.onJoin(player.getUniqueId(), player.getName());
    }
}
