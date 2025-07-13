package net.okocraft.monitor.core.config;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.decoder.Decoder;
import dev.siroshun.codec4j.api.decoder.object.ObjectDecoder;
import net.okocraft.monitor.core.config.notification.OreNotification;
import net.okocraft.monitor.core.config.notification.ServerStatusNotification;

public record DiscordWebhookConfig(String url, Notifications notifications) {

    public static final DiscordWebhookConfig EMPTY = new DiscordWebhookConfig(
        "",
        new Notifications(
            ServerStatusNotification.EMPTY,
            OreNotification.EMPTY
        )
    );

    public static final Decoder<DiscordWebhookConfig> DECODER = ObjectDecoder.create(
        DiscordWebhookConfig::new,
        Codec.STRING.toRequiredFieldDecoder("url"),
        Notifications.DECODER.toRequiredFieldDecoder("notifications")
    );

    public record Notifications(
        ServerStatusNotification serverStatus,
        OreNotification ore
    ) {

        private static final Decoder<Notifications> DECODER = ObjectDecoder.create(
            Notifications::new,
            ServerStatusNotification.DECODER.toOptionalFieldDecoder("server-status", ServerStatusNotification.EMPTY),
            OreNotification.DECODER.toOptionalFieldDecoder("ore", OreNotification.EMPTY)
        );
    }

}
