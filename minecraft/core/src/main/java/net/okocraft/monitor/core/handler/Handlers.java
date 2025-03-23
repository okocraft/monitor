package net.okocraft.monitor.core.handler;

import net.okocraft.monitor.core.manager.PlayerManager;
import net.okocraft.monitor.core.storage.Storage;

public record Handlers(PlayerHandler player) {

    public static Handlers initialize(int serverId, Storage storage, PlayerManager playerManager) {
        return new Handlers(
            new PlayerHandler(serverId, storage.getPlayerStorage(), playerManager)
        );
    }

}
