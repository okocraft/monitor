package net.okocraft.monitor.core.bootstrap;

import net.okocraft.monitor.core.Monitor;
import net.okocraft.monitor.core.config.MonitorConfig;
import net.okocraft.monitor.core.logger.MonitorLogger;
import org.jetbrains.annotations.NotNullByDefault;
import org.slf4j.Logger;
import org.slf4j.helpers.SubstituteLogger;

import java.nio.file.Path;
import java.util.Optional;

@NotNullByDefault
public final class MonitorBootstrap {

    public static Optional<Monitor> load(Logger logger, Path dataDirectory) {
        ((SubstituteLogger) MonitorLogger.logger()).setDelegate(logger);
        var config = loadConfig(dataDirectory.resolve("config.yml"));
        if (config.isEmpty()) {
            return Optional.empty();
        }

        MonitorLogger.logger().info("Successfully loaded Monitor!");
        return Optional.of(new Monitor(dataDirectory, config.get()));
    }

    private static Optional<MonitorConfig> loadConfig(Path filepath) {
        MonitorLogger.logger().info("Loading configuration from {}", Path.of(".").relativize(filepath));
        return Optional.of(new MonitorConfig());
    }

    private MonitorBootstrap() {
        throw new UnsupportedOperationException();
    }
}
