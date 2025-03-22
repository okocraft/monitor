package net.okocraft.monitor.core;

import net.okocraft.monitor.core.config.MonitorConfig;
import net.okocraft.monitor.core.database.Database;
import net.okocraft.monitor.core.database.mysql.MySQLDatabase;
import net.okocraft.monitor.core.logger.MonitorLogger;
import org.jetbrains.annotations.NotNullByDefault;

import java.nio.file.Path;

@NotNullByDefault
public final class Monitor {

    private final Path dataDirectory;
    private final MonitorConfig config;
    private final Database database;

    public Monitor(Path dataDirectory, MonitorConfig config) {
        this.dataDirectory = dataDirectory;
        this.config = config;
        this.database = new MySQLDatabase(null); // TODO: load from config
    }

    public void start() {
        MonitorLogger.logger().info("Starting Monitor...");

        MonitorLogger.logger().info("Connecting to database...");
        try {
            this.database.prepare();
        } catch (Exception e) {
            MonitorLogger.logger().error("Failed to connect to database.", e);

            try {
                this.database.shutdown();
            } catch (Exception e2) {
                MonitorLogger.logger().error("Failed to shutdown database", e2);
            }

            return;
        }

        MonitorLogger.logger().info("Successfully started Monitor!");
    }

    public void shutdown() {
        MonitorLogger.logger().info("Shutting down Monitor...");

        try {
            this.database.shutdown();
        } catch (Exception e) {
            MonitorLogger.logger().error("Failed to shutdown database", e);
        }

        MonitorLogger.logger().info("Successfully shutdown Monitor!");
    }
}
