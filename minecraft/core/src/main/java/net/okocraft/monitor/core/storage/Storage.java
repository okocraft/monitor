package net.okocraft.monitor.core.storage;

import net.okocraft.monitor.core.database.Database;
import org.jetbrains.annotations.NotNullByDefault;

import java.nio.file.Path;

@NotNullByDefault
public class Storage {

    private final ServerInfoStorage serverInfoStorage;
    private final PlayerStorage playerStorage;
    private final WorldStorage worldStorage;
    private final PlayerLogStorage playerLogStorage;

    public Storage(Path dataDirectory, Database database) {
        this.serverInfoStorage = new ServerInfoStorage(database);
        this.playerStorage = new PlayerStorage(database);
        this.worldStorage = new WorldStorage(database);
        this.playerLogStorage = new PlayerLogStorage(database);
    }

    public ServerInfoStorage getServerInfoStorage() {
        return this.serverInfoStorage;
    }

    public PlayerStorage getPlayerStorage() {
        return this.playerStorage;
    }

    public WorldStorage getWorldStorage() {
        return this.worldStorage;
    }

    public PlayerLogStorage getPlayerLogStorage() {
        return this.playerLogStorage;
    }
}
