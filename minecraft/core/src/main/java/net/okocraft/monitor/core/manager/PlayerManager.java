package net.okocraft.monitor.core.manager;

import net.okocraft.monitor.core.models.MonitorPlayer;
import org.jetbrains.annotations.NotNullByDefault;
import org.jetbrains.annotations.Nullable;

import java.util.Map;
import java.util.UUID;
import java.util.concurrent.ConcurrentHashMap;

@NotNullByDefault
public class PlayerManager {

    private final Map<UUID, MonitorPlayer> playerMap = new ConcurrentHashMap<>();

    public @Nullable MonitorPlayer getPlayerByUUID(UUID uuid) {
        return this.playerMap.get(uuid);
    }

    public void putPlayer(MonitorPlayer player) {
        this.playerMap.putIfAbsent(player.uuid(), player);
    }

    public @Nullable MonitorPlayer removePlayer(UUID uuid) {
        return this.playerMap.remove(uuid);
    }
}
