package net.okocraft.monitor.core.handler;

import net.kyori.adventure.key.Key;
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
import net.okocraft.monitor.core.models.logs.PlayerProxyCommandLog;
import net.okocraft.monitor.core.models.logs.PlayerRenameItemLog;
import net.okocraft.monitor.core.models.logs.PlayerWorldCommandLog;
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
    private final LoggingQueue<PlayerWorldCommandLog> worldCommandLogQueue;
    private final LoggingQueue<PlayerProxyCommandLog> proxyCommandLogQueue;
    private final LoggingQueue<PlayerRenameItemLog> renameItemLogQueue;

    public PlayerHandler(int serverId, PlayerStorage playerStorage,
                         PlayerManager playerManager, WorldManager worldManager,
                         LoggingQueueHolder queueHolder, PlayerLogStorage playerLogStorage) {
        this.serverId = serverId;
        this.playerStorage = playerStorage;
        this.playerManager = playerManager;
        this.worldManager = worldManager;
        this.connectLogQueue = queueHolder.createQueue(playerLogStorage::saveConnectLogs, 100);
        this.chatLogQueue = queueHolder.createQueue(playerLogStorage::saveChatLogs, 100);
        this.worldCommandLogQueue = queueHolder.createQueue(playerLogStorage::saveWorldCommandLogs, 100);
        this.proxyCommandLogQueue = queueHolder.createQueue(playerLogStorage::saveProxyCommandLogs, 100);
        this.renameItemLogQueue = queueHolder.createQueue(playerLogStorage::saveRenameItemLogs, 15);
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
        if (world == null) {
            return;
        }
        this.chatLogQueue.push(new PlayerChatLog(player.playerId(), world.worldId(), position, PlainTextComponentSerializer.plainText().serialize(message), LocalDateTime.now()));
    }

    public void onWorldCommand(UUID uuid, UUID worldUid, BlockPosition position, String command) {
        MonitorPlayer player = this.playerManager.getPlayerByUUID(uuid);
        if (player == null) {
            return;
        }
        MonitorWorld world = this.worldManager.getWorldByUUID(worldUid);
        if (world == null) {
            return;
        }
        this.worldCommandLogQueue.push(new PlayerWorldCommandLog(player.playerId(), world.worldId(), position, command, LocalDateTime.now()));
    }

    public void onProxyCommand(UUID uuid, String command) {
        MonitorPlayer player = this.playerManager.getPlayerByUUID(uuid);
        if (player == null) {
            return;
        }
        this.proxyCommandLogQueue.push(new PlayerProxyCommandLog(player.playerId(), this.serverId, command, LocalDateTime.now()));
    }

    public void onRenameItem(UUID uuid, UUID worldUid, BlockPosition position, Key itemType, Component itemName, int amount) {
        MonitorPlayer player = this.playerManager.getPlayerByUUID(uuid);
        if (player == null) {
            return;
        }
        MonitorWorld world = this.worldManager.getWorldByUUID(worldUid);
        if (world == null) {
            return;
        }
        this.renameItemLogQueue.push(new PlayerRenameItemLog(player.playerId(), world.worldId(), position, itemType, itemName, amount, LocalDateTime.now()));
    }
}
