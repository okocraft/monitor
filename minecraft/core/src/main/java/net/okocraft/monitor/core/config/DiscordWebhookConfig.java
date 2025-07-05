package net.okocraft.monitor.core.config;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.decoder.Decoder;
import dev.siroshun.codec4j.api.decoder.object.ObjectDecoder;
import net.okocraft.monitor.core.config.notification.ServerStatusNotification;

import java.util.Set;

public record DiscordWebhookConfig(String url, Notifications notifications) {

    public static final DiscordWebhookConfig EMPTY = new DiscordWebhookConfig(
        "",
        new Notifications(
            new ServerStatusNotification(
                Set.of(),
                0L,
                new ServerStatusNotification.Setting(false, ""),
                new ServerStatusNotification.Setting(false, ""),
                new ServerStatusNotification.Setting(false, ""),
                new ServerStatusNotification.Setting(false, ""),
                new ServerStatusNotification.Setting(false, "")
            )
        )
    );

    public static final Decoder<DiscordWebhookConfig> DECODER = ObjectDecoder.create(
        DiscordWebhookConfig::new,
        Codec.STRING.toRequiredFieldDecoder("url"),
        Notifications.DECODER.toRequiredFieldDecoder("notifications")
    );

    public record Notifications(
        ServerStatusNotification serverStatus
    ) {

        private static final Decoder<Notifications> DECODER = ObjectDecoder.create(
            Notifications::new,
            ServerStatusNotification.CODEC.toRequiredFieldDecoder("server-status")
        );

    }

}
