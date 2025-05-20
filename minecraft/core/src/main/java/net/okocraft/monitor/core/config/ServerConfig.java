package net.okocraft.monitor.core.config;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.codec.object.ObjectCodec;
import org.jetbrains.annotations.NotNullByDefault;

import java.nio.file.Path;

@NotNullByDefault
public record ServerConfig(String name) {

    public static final Codec<ServerConfig> CODEC = ObjectCodec.create(
        ServerConfig::new,
        Codec.STRING.toFieldCodec("name").defaultValueSupplier(() -> "").required(ServerConfig::name)
    );

    public String getServerName() {
        var env = System.getenv("MONITOR_SERVER_NAME");
        if (env != null && !env.isEmpty()) {
            return env;
        }

        var property = System.getProperty("monitor.server.name");
        if (property != null && !property.isEmpty()) {
            return property;
        }

        if (!this.name.isEmpty()) {
            return this.name;
        }

        return Path.of(".").toAbsolutePath().normalize().getFileName().toString();
    }
}
