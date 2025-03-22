package net.okocraft.monitor.platform.paper;

import net.okocraft.monitor.core.Monitor;
import net.okocraft.monitor.core.bootstrap.MonitorBootstrap;
import net.okocraft.monitor.core.logger.MonitorLogger;
import org.bukkit.plugin.java.JavaPlugin;

public class MonitorPaperPlugin extends JavaPlugin {

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

        try {
            this.monitor.start();
        } catch (Exception e) {
            MonitorLogger.logger().error("Failed to start Monitor", e);
        }
    }

    @Override
    public void onDisable() {
        if (this.monitor == null) {
            return;
        }

        try {
            this.monitor.shutdown();
        } catch (Exception e) {
            MonitorLogger.logger().error("Failed to shutdown Monitor", e);
        }
    }
}
