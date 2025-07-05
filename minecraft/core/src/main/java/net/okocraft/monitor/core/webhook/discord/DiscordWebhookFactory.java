package net.okocraft.monitor.core.webhook.discord;

import club.minnced.discord.webhook.WebhookClientBuilder;
import net.okocraft.monitor.core.util.NamedThreadFactory;
import org.jetbrains.annotations.Nullable;

public class DiscordWebhookFactory {

    private final String url;

    public DiscordWebhookFactory(String url) {
        this.url = url;
    }

    public @Nullable DiscordWebhook create(long threadId) {
        if (this.url.isEmpty()) {
            return null;
        }

        return new DiscordWebhook(
            new WebhookClientBuilder(this.url)
                .setThreadId(threadId)
                .setThreadFactory(NamedThreadFactory.DEFAULT)
                .build()
        );
    }
}
