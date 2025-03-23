package net.okocraft.monitor.core.storage;

import net.okocraft.monitor.core.database.Database;
import org.jetbrains.annotations.NotNullByDefault;

import java.nio.file.Path;

@NotNullByDefault
public class Storage {

    private final ServerInfoStorage serverInfoStorage;
    private final PlayerStorage playerStorage;

    public Storage(Path dataDirectory, Database database) {
        this.serverInfoStorage = new ServerInfoStorage(database);
        this.playerStorage = new PlayerStorage(database);
    }

    public ServerInfoStorage getServerInfoStorage() {
        return this.serverInfoStorage;
    }

    public PlayerStorage getPlayerStorage() {
        return this.playerStorage;
    }
}
