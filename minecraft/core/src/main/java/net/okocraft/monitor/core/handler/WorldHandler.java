package net.okocraft.monitor.core.handler;

import net.okocraft.monitor.core.logger.MonitorLogger;
import net.okocraft.monitor.core.manager.WorldManager;
import net.okocraft.monitor.core.models.MonitorWorld;
import net.okocraft.monitor.core.storage.WorldStorage;

import java.util.UUID;

public class WorldHandler {

    private final int serverId;
    private final WorldStorage worldStorage;
    private final WorldManager worldManager;

    public WorldHandler(int serverId, WorldStorage worldStorage, WorldManager worldManager) {
        this.serverId = serverId;
        this.worldStorage = worldStorage;
        this.worldManager = worldManager;
    }

    public void onLoad(UUID uid, String name) {
        MonitorWorld loadedWorld = this.worldManager.getWorldByUUID(uid);
        if (loadedWorld != null && loadedWorld.name().equals(name)) {
            return;
        }
        try {
            MonitorWorld world = this.worldStorage.initializeWorld(this.serverId, uid, name);
            this.worldManager.putWorld(world);
        } catch (Exception e) {
            MonitorLogger.logger().error("Failed to initialize world: {}", name, e);
        }
    }
}
