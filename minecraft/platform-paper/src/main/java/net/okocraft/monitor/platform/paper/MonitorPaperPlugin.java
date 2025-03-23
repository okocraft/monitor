package net.okocraft.monitor.platform.paper;

import net.okocraft.monitor.core.Monitor;
import net.okocraft.monitor.core.bootstrap.MonitorBootstrap;
import net.okocraft.monitor.core.handler.Handlers;
import net.okocraft.monitor.core.logger.MonitorLogger;
import net.okocraft.monitor.core.platform.PlatformAdapter;
import net.okocraft.monitor.platform.paper.listener.PlayerListener;
import org.bukkit.event.HandlerList;
import org.bukkit.event.Listener;
import org.bukkit.plugin.java.JavaPlugin;

import java.util.stream.Stream;

public class MonitorPaperPlugin extends JavaPlugin implements PlatformAdapter {

    private Monitor monitor;

    @Override
    public void onLoad() {
        var monitor = MonitorBootstrap.load(this.getSLF4JLogger(), this.getDataPath());
        if (monitor.isEmpty()) {
            MonitorLogger.logger().error("Failed to load Monitor");
            return;
        }
        this.monitor = monitor.get();
    }

    @Override
    public void onEnable() {
        if (this.monitor == null) {
            return;
        }

        this.monitor.start(this);
    }

    @Override
    public void onDisable() {
        if (this.monitor == null) {
            return;
        }

        this.monitor.shutdown(this);
    }

    @Override
    public void registerEventListeners(Handlers handlers) {
        Stream.<Listener>of(
            new PlayerListener(handlers.player())
        ).forEach(listener -> this.getServer().getPluginManager().registerEvents(listener, this));
    }

    @Override
    public void unregisterEventListeners() {
        HandlerList.unregisterAll(this);
    }
}
