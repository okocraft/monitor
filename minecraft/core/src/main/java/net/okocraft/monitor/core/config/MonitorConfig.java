package net.okocraft.monitor.core.config;

import dev.siroshun.codec4j.api.decoder.Decoder;
import dev.siroshun.codec4j.api.decoder.object.FieldDecoder;
import dev.siroshun.codec4j.api.decoder.object.ObjectDecoder;
import dev.siroshun.codec4j.io.yaml.YamlIO;
import org.jetbrains.annotations.NotNullByDefault;

import java.nio.file.Path;
import java.util.concurrent.atomic.AtomicReference;

@NotNullByDefault
public record MonitorConfig(DatabaseConfig database, ServerConfig server, UploadConfig upload, CommandConfig command, DiscordWebhookConfig discordWebhook) {

    public static final Decoder<MonitorConfig> CODEC = ObjectDecoder.create(
        MonitorConfig::new,
        FieldDecoder.required("database", DatabaseConfig.CODEC),
        FieldDecoder.required("server", ServerConfig.CODEC),
        FieldDecoder.optional("upload", UploadConfig.CODEC, UploadConfig.EMPTY),
        FieldDecoder.optional("command", CommandConfig.CODEC, CommandConfig.EMPTY),
        FieldDecoder.optional("discord-webhook", DiscordWebhookConfig.DECODER, DiscordWebhookConfig.EMPTY)
    );

    public static Holder load(Path filepath) throws Exception {
        var result = YamlIO.DEFAULT.decodeFrom(filepath, CODEC);
        if (result.isFailure()) {
            throw new Exception("Unable to load config.yml: " + result);
        }
        return new Holder(filepath, result.unwrap());
    }

    public static final class Holder {

        private final Path filepath;
        private final AtomicReference<MonitorConfig> ref;

        public Holder(Path filepath, MonitorConfig initial) {
            this.filepath = filepath;
            this.ref = new AtomicReference<>(initial);
        }

        public MonitorConfig get() {
            return this.ref.get();
        }

        public void reload() throws Exception {
            var result = YamlIO.DEFAULT.decodeFrom(this.filepath, CODEC);
            if (result.isFailure()) {
                throw new Exception("Unable to load config.yml: " + result);
            }
            this.ref.set(result.unwrap());
        }
    }
}
