package net.okocraft.monitor.core.handler;

import net.kyori.adventure.text.Component;
import net.kyori.adventure.text.serializer.plain.PlainTextComponentSerializer;
import net.okocraft.monitor.core.logger.MonitorLogger;
import net.okocraft.monitor.core.manager.PlayerManager;
import net.okocraft.monitor.core.manager.WorldManager;
import net.okocraft.monitor.core.models.BlockPosition;
import net.okocraft.monitor.core.models.MonitorPlayer;
import net.okocraft.monitor.core.models.MonitorWorld;
import net.okocraft.monitor.core.models.logs.PlayerChatLog;
import net.okocraft.monitor.core.models.logs.PlayerConnectLog;
import net.okocraft.monitor.core.queue.LoggingQueue;
import net.okocraft.monitor.core.queue.LoggingQueueHolder;
import net.okocraft.monitor.core.storage.PlayerLogStorage;
import net.okocraft.monitor.core.storage.PlayerStorage;
import org.jetbrains.annotations.Nullable;

import java.net.SocketAddress;
import java.time.LocalDateTime;
import java.util.Objects;
import java.util.UUID;

public class PlayerHandler {

    private final int serverId;
    private final PlayerStorage playerStorage;
    private final PlayerManager playerManager;
    private final WorldManager worldManager;
    private final LoggingQueue<PlayerConnectLog> connectLogQueue;
    private final LoggingQueue<PlayerChatLog> chatLogQueue;

    public PlayerHandler(int serverId, PlayerStorage playerStorage,
                         PlayerManager playerManager, WorldManager worldManager,
                         LoggingQueueHolder queueHolder, PlayerLogStorage playerLogStorage) {
        this.serverId = serverId;
        this.playerStorage = playerStorage;
        this.playerManager = playerManager;
        this.worldManager = worldManager;
        this.connectLogQueue = queueHolder.createQueue(playerLogStorage::saveConnectLogs, 100);
        this.chatLogQueue = queueHolder.createQueue(playerLogStorage::saveChatLogs, 100);
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

        this.connectLogQueue.push(new PlayerConnectLog(player.playerId(), this.serverId, PlayerConnectLog.Action.CONNECT, Objects.toString(address), "", LocalDateTime.now()));
    }

    public void onLeave(UUID uuid, PlayerConnectLog.Action action, @Nullable SocketAddress address, Component reason) {
        MonitorPlayer player = this.playerManager.getPlayerByUUID(uuid);
        if (player == null) {
            return;
        }
        this.connectLogQueue.push(new PlayerConnectLog(player.playerId(), this.serverId, action, Objects.toString(address), PlainTextComponentSerializer.plainText().serialize(reason), LocalDateTime.now()));
    }

    public void onChat(UUID uuid, UUID worldUid, BlockPosition position, Component message) {
        MonitorPlayer player = this.playerManager.getPlayerByUUID(uuid);
        if (player == null) {
            return;
        }
        MonitorWorld world = this.worldManager.getWorldByUUID(worldUid);
        int worldId = world != null ? world.worldId() : 0;
        this.chatLogQueue.push(new PlayerChatLog(player.playerId(), worldId, position, PlainTextComponentSerializer.plainText().serialize(message), LocalDateTime.now()));
    }
}
