package net.okocraft.monitor.core;

import net.okocraft.monitor.core.config.MonitorConfig;
import net.okocraft.monitor.core.logger.MonitorLogger;

import java.nio.file.Path;

public final class Monitor {

    private final Path dataDirectory;
    private final MonitorConfig config;

    public Monitor(Path dataDirectory, MonitorConfig config) {
        this.dataDirectory = dataDirectory;
        this.config = config;
    }

    public void start() throws Exception {
        MonitorLogger.logger().info("Starting Monitor...");

        MonitorLogger.logger().info("Successfully started Monitor!");
    }

    public void shutdown() throws Exception {
        MonitorLogger.logger().info("Shutting down Monitor...");

        MonitorLogger.logger().info("Successfully shutdown Monitor!");
    }
}
