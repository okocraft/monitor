package net.okocraft.monitor.core.config.notification;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.decoder.Decoder;
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
        Codec.STRING.toSetDecoder().toRequiredFieldDecoder("enabled-server-names"),
        Codec.LONG.toOptionalFieldDecoder("thread-id", 0L),
        Setting.CODEC.toRequiredFieldDecoder("current-status"),
        Setting.CODEC.toRequiredFieldDecoder("server-started"),
        Setting.CODEC.toRequiredFieldDecoder("server-stopped"),
        Setting.CODEC.toRequiredFieldDecoder("server-not-started"),
        Setting.CODEC.toRequiredFieldDecoder("first-ping-failure")
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
            Codec.BOOLEAN.toRequiredFieldDecoder("enabled"),
            Codec.STRING.toRequiredFieldDecoder("message")
        );
    }
}
