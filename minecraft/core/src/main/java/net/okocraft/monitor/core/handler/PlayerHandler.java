package net.okocraft.monitor.core.handler;

import net.okocraft.monitor.core.logger.MonitorLogger;
import net.okocraft.monitor.core.manager.PlayerManager;
import net.okocraft.monitor.core.models.MonitorPlayer;
import net.okocraft.monitor.core.storage.PlayerStorage;

import java.util.UUID;

public class PlayerHandler {

    private final int serverId;
    private final PlayerStorage playerStorage;
    private final PlayerManager playerManager;

    public PlayerHandler(int serverId, PlayerStorage playerStorage, PlayerManager playerManager) {
        this.serverId = serverId;
        this.playerStorage = playerStorage;
        this.playerManager = playerManager;
    }

    public void onJoin(UUID uuid, String name) {
        try {
            MonitorPlayer player = this.playerStorage.initializePlayer(uuid, name);
            this.playerManager.putPlayer(player);
        } catch (Exception e) {
            MonitorLogger.logger().error("Failed to initialize player", e);
        }
    }
}
