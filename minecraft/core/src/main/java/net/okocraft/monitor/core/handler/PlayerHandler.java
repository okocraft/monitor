package net.okocraft.monitor.core.handler;

import net.kyori.adventure.text.Component;
import net.kyori.adventure.text.serializer.plain.PlainTextComponentSerializer;
import net.okocraft.monitor.core.logger.MonitorLogger;
import net.okocraft.monitor.core.manager.PlayerManager;
import net.okocraft.monitor.core.models.MonitorPlayer;
import net.okocraft.monitor.core.models.logs.PlayerConnectLog;
import net.okocraft.monitor.core.queue.LoggingQueue;
import net.okocraft.monitor.core.queue.LoggingQueueHolder;
import net.okocraft.monitor.core.storage.PlayerLogStorage;
import net.okocraft.monitor.core.storage.PlayerStorage;
import org.jetbrains.annotations.Nullable;

import java.net.SocketAddress;
import java.util.Objects;
import java.util.UUID;

public class PlayerHandler {

    private final int serverId;
    private final PlayerStorage playerStorage;
    private final PlayerManager playerManager;
    private final LoggingQueue<PlayerConnectLog> queue;

    public PlayerHandler(int serverId, PlayerStorage playerStorage, PlayerManager playerManager,
                         LoggingQueueHolder queueHolder, PlayerLogStorage playerLogStorage) {
        this.serverId = serverId;
        this.playerStorage = playerStorage;
        this.playerManager = playerManager;
        this.queue = queueHolder.createQueue(playerLogStorage::saveConnectLogs, 100);
    }

    public void onJoin(UUID uuid, String name, @Nullable SocketAddress address) {
        MonitorPlayer player;
        try {
            player = this.playerStorage.initializePlayer(uuid, name);
            this.playerManager.putPlayer(player);
        } catch (Exception e) {
            MonitorLogger.logger().error("Failed to initialize player", e);
            return;
        }

        this.queue.push(new PlayerConnectLog(player.playerId(), this.serverId, PlayerConnectLog.Action.CONNECT, Objects.toString(address), ""));
    }

    public void onLeave(UUID uuid, PlayerConnectLog.Action action, @Nullable SocketAddress address, Component reason) {
        MonitorPlayer player = this.playerManager.getPlayerByUUID(uuid);
        if (player == null) {
            return;
        }
        this.queue.push(new PlayerConnectLog(player.playerId(), this.serverId, action, Objects.toString(address), PlainTextComponentSerializer.plainText().serialize(reason)));
    }
}
