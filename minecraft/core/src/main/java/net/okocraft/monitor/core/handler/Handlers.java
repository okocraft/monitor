package net.okocraft.monitor.core.handler;

import net.okocraft.monitor.core.manager.PlayerManager;
import net.okocraft.monitor.core.manager.WorldManager;
import net.okocraft.monitor.core.queue.LoggingQueueHolder;
import net.okocraft.monitor.core.storage.Storage;

public record Handlers(PlayerHandler player, WorldHandler world) {

    public static Handlers initialize(int serverId, Storage storage, PlayerManager playerManager, WorldManager worldManager, LoggingQueueHolder queueHolder) {
        return new Handlers(
            new PlayerHandler(serverId, storage.getPlayerStorage(), playerManager, queueHolder, storage.getPlayerLogStorage()),
            new WorldHandler(serverId, storage.getWorldStorage(), worldManager)
        );
    }

}
