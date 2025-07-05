package net.okocraft.monitor.core.webhook.discord;

import club.minnced.discord.webhook.WebhookClient;

public final class DiscordWebhook {

    private final WebhookClient client;

    DiscordWebhook(WebhookClient client) {
        this.client = client;
    }

    public void send(String message) {
        this.client.send(message);
    }
}
