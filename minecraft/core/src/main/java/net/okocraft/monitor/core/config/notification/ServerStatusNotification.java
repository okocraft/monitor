package net.okocraft.monitor.core.config.notification;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.decoder.Decoder;
import dev.siroshun.codec4j.api.decoder.collection.SetDecoder;
import dev.siroshun.codec4j.api.decoder.object.FieldDecoder;
import dev.siroshun.codec4j.api.decoder.object.ObjectDecoder;

import java.util.Set;

public record ServerStatusNotification(
    Set<String> enabledServerNames,
    long threadId,
    Setting currentStatus,
    Setting serverStarted,
    Setting serverStopped,
    Setting serverNotStarted,
    Setting firstPingFailure
) {

    public static final Decoder<ServerStatusNotification> DECODER = ObjectDecoder.create(
        ServerStatusNotification::new,
        FieldDecoder.required("enabled-server-names", SetDecoder.create(Codec.STRING)),
        FieldDecoder.optional("thread-id", Codec.LONG, 0L),
        FieldDecoder.required("current-status", Setting.CODEC),
        FieldDecoder.required("server-started", Setting.CODEC),
        FieldDecoder.required("server-stopped", Setting.CODEC),
        FieldDecoder.required("server-not-started", Setting.CODEC),
        FieldDecoder.required("first-ping-failure", Setting.CODEC)
    );

    public static final ServerStatusNotification EMPTY =             new ServerStatusNotification(
        Set.of(),
        0L,
        new ServerStatusNotification.Setting(false, ""),
        new ServerStatusNotification.Setting(false, ""),
        new ServerStatusNotification.Setting(false, ""),
        new ServerStatusNotification.Setting(false, ""),
        new ServerStatusNotification.Setting(false, "")
    );

    public record Setting(boolean enabled, String message) {
        public static final Decoder<Setting> CODEC = ObjectDecoder.create(
            Setting::new,
            FieldDecoder.required("enabled", Codec.BOOLEAN),
            FieldDecoder.required("message", Codec.STRING)
        );
    }
}
