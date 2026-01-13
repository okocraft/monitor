package net.okocraft.monitor.core.config;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.decoder.Decoder;
import dev.siroshun.codec4j.api.decoder.object.FieldDecoder;
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
        FieldDecoder.required("url", Codec.STRING),
        FieldDecoder.required("notifications", Notifications.DECODER)
    );

    public record Notifications(
        ServerStatusNotification serverStatus,
        OreNotification ore
    ) {

        private static final Decoder<Notifications> DECODER = ObjectDecoder.create(
            Notifications::new,
            FieldDecoder.optional("server-status", ServerStatusNotification.DECODER, ServerStatusNotification.EMPTY),
            FieldDecoder.optional("ore", OreNotification.DECODER, OreNotification.EMPTY)
        );
    }

}
