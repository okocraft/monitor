package net.okocraft.monitor.platform.paper.listener;

import net.okocraft.monitor.core.handler.WorldHandler;
import org.bukkit.Bukkit;
import org.bukkit.World;
import org.bukkit.event.EventHandler;
import org.bukkit.event.Listener;
import org.bukkit.event.server.ServerLoadEvent;
import org.bukkit.event.world.WorldLoadEvent;

public class WorldListener implements Listener {

    private final WorldHandler handler;

    public WorldListener(WorldHandler handler) {
        this.handler = handler;
    }

    @EventHandler
    public void onServerLoad(ServerLoadEvent event) {
        for (World world : Bukkit.getWorlds()) {
            this.handler.onLoad(world.getUID(), world.getName());
        }
    }

    @EventHandler
    public void onWorldLoad(WorldLoadEvent event) {
        this.handler.onLoad(event.getWorld().getUID(), event.getWorld().getName());
    }
}
