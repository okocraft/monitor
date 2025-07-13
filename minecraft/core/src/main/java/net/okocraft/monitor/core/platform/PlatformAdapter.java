package net.okocraft.monitor.core.platform;

import net.okocraft.monitor.core.command.Command;
import net.okocraft.monitor.core.config.notification.OreNotification;
import net.okocraft.monitor.core.config.notification.ServerStatusNotification;
import net.okocraft.monitor.core.handler.Handlers;
import net.okocraft.monitor.core.webhook.discord.DiscordWebhook;
import org.jetbrains.annotations.NotNullByDefault;
import org.jetbrains.annotations.Nullable;

import java.util.concurrent.TimeUnit;

@NotNullByDefault
public interface PlatformAdapter {

    String pluginVersion();

    void registerEventListeners(Handlers handlers);

    void registerVeinFindListener(OreNotification setting, @Nullable DiscordWebhook webhook);

    void unregisterEventListeners();

    void registerCommand(String label, Command command);

    CancellableTask startServerStatusChecker(ServerStatusNotification notification, @Nullable DiscordWebhook webhook);

    CancellableTask scheduleTask(Runnable task, long delay, long interval, TimeUnit unit);

}
