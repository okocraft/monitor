package net.okocraft.monitor.core;

import net.okocraft.monitor.core.config.MonitorConfig;
import net.okocraft.monitor.core.database.Database;
import net.okocraft.monitor.core.database.mysql.MySQLDatabase;
import net.okocraft.monitor.core.handler.Handlers;
import net.okocraft.monitor.core.logger.MonitorLogger;
import net.okocraft.monitor.core.manager.PlayerManager;
import net.okocraft.monitor.core.manager.WorldManager;
import net.okocraft.monitor.core.platform.CancellableTask;
import net.okocraft.monitor.core.platform.PlatformAdapter;
import net.okocraft.monitor.core.queue.LoggingQueueHolder;
import net.okocraft.monitor.core.storage.Storage;
import org.jetbrains.annotations.NotNullByDefault;
import org.jetbrains.annotations.Nullable;

import java.nio.file.Path;
import java.util.concurrent.TimeUnit;

@NotNullByDefault
public final class Monitor {

    private final Path dataDirectory;
    private final MonitorConfig.Holder configHolder;
    private final Database database;

    private @Nullable CancellableTask saveLogTask;
    private @Nullable LoggingQueueHolder loggingQueueHolder;

    public Monitor(Path dataDirectory, MonitorConfig.Holder configHolder) {
        this.dataDirectory = dataDirectory;
        this.configHolder = configHolder;
        this.database = new MySQLDatabase(configHolder.get().database().mysql());
    }

    public void start(PlatformAdapter adapter) {
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

        Storage storage = new Storage(this.dataDirectory, this.database);

        String serverName = this.configHolder.get().server().getServerName();
        int serverId;
        try {
            serverId = storage.getServerInfoStorage().initializeServerId(serverName);
        } catch (Exception e) {
            MonitorLogger.logger().error("Failed to initialize server id", e);
            return;
        }

        MonitorLogger.logger().info("{}'s server id: {}", serverName, serverId);

        PlayerManager playerManager = new PlayerManager();
        WorldManager worldManager = new WorldManager();
        this.loggingQueueHolder = new LoggingQueueHolder();
        Handlers handlers = Handlers.initialize(serverId, storage, playerManager, worldManager, this.loggingQueueHolder);

        MonitorLogger.logger().info("Registering event listeners...");
        adapter.registerEventListeners(handlers);

        MonitorLogger.logger().info("Scheduling logging task...");
        this.saveLogTask = adapter.scheduleTask( this.loggingQueueHolder::handleLimited, 10, 5, TimeUnit.SECONDS);

        MonitorLogger.logger().info("Successfully started Monitor!");
    }

    public void shutdown(PlatformAdapter adapter) {
        MonitorLogger.logger().info("Shutting down Monitor...");

        MonitorLogger.logger().info("Unregistering event listeners...");
        adapter.unregisterEventListeners();

        MonitorLogger.logger().info("Cancelling logging task...");
        if (this.saveLogTask != null) {
            this.saveLogTask.cancel();
        }

        if (this.loggingQueueHolder != null) {
            this.loggingQueueHolder.handleAll();
        }

        try {
            this.database.shutdown();
        } catch (Exception e) {
            MonitorLogger.logger().error("Failed to shutdown database", e);
        }

        MonitorLogger.logger().info("Successfully shutdown Monitor!");
    }
}
