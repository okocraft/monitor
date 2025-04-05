package net.okocraft.monitor.core.manager;

import net.okocraft.monitor.core.models.MonitorWorld;
import org.jetbrains.annotations.NotNullByDefault;
import org.jetbrains.annotations.Nullable;

import java.util.Map;
import java.util.UUID;
import java.util.concurrent.ConcurrentHashMap;

@NotNullByDefault
public class WorldManager {

    private final Map<UUID, MonitorWorld> worldMap = new ConcurrentHashMap<>();

    public @Nullable MonitorWorld getWorldByUUID(UUID uuid)  {
        return this.worldMap.get(uuid);
    }

    public void putWorld(MonitorWorld world) {
        this.worldMap.putIfAbsent(world.uid(), world);
    }

    public void removeWorld(UUID uid) {
        this.worldMap.remove(uid);
    }
}
