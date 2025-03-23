package net.okocraft.monitor.core.storage;

import net.okocraft.monitor.core.database.Database;
import org.jetbrains.annotations.NotNullByDefault;

import java.sql.Connection;
import java.sql.SQLException;
import java.util.OptionalInt;

@NotNullByDefault
public class ServerInfoStorage {

    private final Database database;

    public ServerInfoStorage(Database database) {
        this.database = database;
    }

    public int initializeServerId(String serverName) throws SQLException {
        try (Connection connection = this.database.getConnection()) {
            OptionalInt id = this.database.getOperators().servers().getServerIdByName(connection, serverName);
            if (id.isPresent()) {
                this.database.getOperators().servers().updateServerInfo(connection, id.getAsInt(), serverName);
                return id.getAsInt();
            }
            return this.database.getOperators().servers().insertNewServer(connection, serverName);
        }
    }
}
