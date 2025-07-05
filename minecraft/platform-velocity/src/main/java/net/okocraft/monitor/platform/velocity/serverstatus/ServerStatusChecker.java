package net.okocraft.monitor.platform.velocity.serverstatus;

import com.velocitypowered.api.proxy.ProxyServer;
import com.velocitypowered.api.proxy.server.RegisteredServer;
import com.velocitypowered.api.proxy.server.ServerInfo;
import com.velocitypowered.api.proxy.server.ServerPing;
import net.okocraft.monitor.core.config.notification.ServerStatusNotification;
import net.okocraft.monitor.core.webhook.discord.DiscordWebhook;
import org.jetbrains.annotations.NotNull;

import java.util.HashMap;
import java.util.Map;
import java.util.function.Function;

public class ServerStatusChecker implements Runnable {

    private final Map<ServerInfo, ServerStatus> statusMap = new HashMap<>();
    private final ProxyServer proxy;
    private final ServerStatusNotification notification;
    private final DiscordWebhook webhook;

    public ServerStatusChecker(@NotNull ProxyServer proxy, @NotNull ServerStatusNotification notification,
                               @NotNull DiscordWebhook webhook) {
        this.proxy = proxy;
        this.notification = notification;
        this.webhook = webhook;
    }

    @Override
    public void run() {
        for (RegisteredServer server : this.proxy.getAllServers()) {
            if (!this.notification.enabledServerNames().contains(server.getServerInfo().getName())) {
                continue;
            }

            ServerStatus status = this.statusMap.computeIfAbsent(server.getServerInfo(), ServerStatus::new);
            status.updateStatus(server);
        }
    }

    private class ServerStatus {

        private final ServerInfo serverInfo;
        private Status status = Status.UNKNOWN;

        private ServerStatus(@NotNull ServerInfo serverInfo) {
            this.serverInfo = serverInfo;
        }

        private void updateStatus(@NotNull RegisteredServer server) {
            server.ping().thenAcceptAsync(this::success).exceptionallyAsync(this::failure);
        }

        private void success(@NotNull ServerPing ignored) {
            switch (this.status) {
                case RUNNING, UNKNOWN -> this.sendNotification(ServerStatusNotification::currentStatus, this.serverInfo);
                case FIRST_FAILURE, STOPPED -> this.sendNotification(ServerStatusNotification::serverStarted, this.serverInfo);
            }
            this.status = Status.RUNNING;
        }

        private Void failure(@NotNull Throwable e) {
            switch (this.status) {
                case RUNNING -> {
                    this.status = Status.FIRST_FAILURE;
                    this.sendNotification(ServerStatusNotification::firstPingFailure, this.serverInfo);
                }
                case FIRST_FAILURE -> {
                    this.status = Status.STOPPED;
                    this.sendNotification(ServerStatusNotification::serverStopped, this.serverInfo);
                }
                case UNKNOWN -> {
                    this.status = Status.STOPPED;
                    this.sendNotification(ServerStatusNotification::serverNotStarted, this.serverInfo);
                }
            }
            return null;
        }

        private void sendNotification(@NotNull Function<ServerStatusNotification, ServerStatusNotification.Setting> settingFunction, @NotNull ServerInfo serverInfo) {
            ServerStatusNotification.Setting setting = settingFunction.apply(ServerStatusChecker.this.notification);
            if (setting.enabled() && !setting.message().isEmpty()) {
                ServerStatusChecker.this.webhook.send(setting.message().replace("%server_name%", serverInfo.getName()));
            }
        }

        private enum Status {
            RUNNING,
            FIRST_FAILURE,
            STOPPED,
            UNKNOWN
        }
    }
}
