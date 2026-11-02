package net.okocraft.monitor.platform.neoforge;

import net.okocraft.monitor.core.config.MonitorConfig;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.io.TempDir;

import java.nio.file.Files;
import java.nio.file.Path;

import static org.junit.jupiter.api.Assertions.assertEquals;

class ConfigCompatibilityTest {
    @TempDir
    Path directory;

    @Test
    void decodesSharedMysqlConfigurationWithSnakeYamlTwo() throws Exception {
        Path config = directory.resolve("config.yml");
        Files.writeString(config, """
            database:
              mysql:
                host: localhost
                port: 3306
                database-name: monitor
                username: fixture
                password: fixture
            server:
              name: dqm
            """);

        var loaded = MonitorConfig.load(config).get();
        assertEquals("dqm", loaded.server().name());
        assertEquals("monitor", loaded.database().mysql().databaseName());
        assertEquals(3306, loaded.database().mysql().port());
    }
}
