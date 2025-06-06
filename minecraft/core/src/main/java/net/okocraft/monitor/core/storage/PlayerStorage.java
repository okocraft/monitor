package net.okocraft.monitor.core.storage;

import net.okocraft.monitor.core.database.Database;
import net.okocraft.monitor.core.database.operator.Operators;
import net.okocraft.monitor.core.models.MonitorPlayer;
import org.jetbrains.annotations.NotNullByDefault;

import java.sql.Connection;
import java.sql.SQLException;
import java.util.UUID;

@NotNullByDefault
public class PlayerStorage {

    private final Database database;
    private final Operators operators;

    public PlayerStorage(Database database) {
        this.database = database;
        this.operators = database.getOperators();
    }

    public MonitorPlayer initializePlayer(UUID uuid, String name) throws SQLException {
        MonitorPlayer player;
        try (var connection = this.database.getConnection()) {
            try {
                connection.setAutoCommit(false);
                player = this.initializePlayer(connection, uuid, name);
                connection.commit();
            } catch (SQLException e) {
                connection.rollback();
                throw e;
            } finally {
                connection.setAutoCommit(true);
            }
        }
        return player;
    }

    private MonitorPlayer initializePlayer(Connection connection, UUID uuid, String name) throws SQLException {
        MonitorPlayer player = this.operators.players().getPlayerByUUID(connection, uuid);

        if (player == null) {
            int playerId = this.operators.players().insertPlayer(connection, uuid, name);
            this.operators.playerNameHistory().insertHistory(connection, playerId, name);
            return new MonitorPlayer(playerId, uuid, name);
        }

        if (player.name().equals(name)) {
            return player;
        }

        MonitorPlayer updatedPlayer = new MonitorPlayer(player.playerId(), uuid, name);

        this.operators.players().updatePlayer(connection, updatedPlayer);
        this.operators.playerNameHistory().insertHistory(connection, player.playerId(), name);

        return updatedPlayer;
    }
}
