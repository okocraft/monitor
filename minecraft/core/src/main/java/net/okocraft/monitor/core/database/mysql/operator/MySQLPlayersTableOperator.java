package net.okocraft.monitor.core.database.mysql.operator;

import net.okocraft.monitor.core.database.operator.PlayersTableOperator;
import net.okocraft.monitor.core.models.MonitorPlayer;
import org.jetbrains.annotations.NotNullByDefault;
import org.jetbrains.annotations.Nullable;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;
import java.util.UUID;

@NotNullByDefault
public class MySQLPlayersTableOperator implements PlayersTableOperator {

    @Override
    public @Nullable MonitorPlayer getPlayerByUUID(Connection connection, UUID uuid) throws SQLException {
        try (PreparedStatement statement = connection.prepareStatement("SELECT id, name FROM minecraft_players WHERE uuid = ?")) {
            statement.setBytes(1, MySQLUUID.uuidToBytes(uuid));
            try (ResultSet resultSet = statement.executeQuery()) {
                if (resultSet.next()) {
                    return new MonitorPlayer(
                        resultSet.getInt("id"),
                        uuid,
                        resultSet.getString("name")
                    );
                }
            }
        }
        return null;
    }

    @Override
    public int insertPlayer(Connection connection, UUID uuid, String name) throws SQLException {
        try (PreparedStatement statement = connection.prepareStatement("INSERT INTO minecraft_players (uuid, name) VALUES (?, ?)", Statement.RETURN_GENERATED_KEYS)) {
            statement.setBytes(1, MySQLUUID.uuidToBytes(uuid));
            statement.setString(2, name);
            statement.executeUpdate();
            try (ResultSet resultSet = statement.getGeneratedKeys()) {
                if (resultSet.next()) {
                    return resultSet.getInt(1);
                }
            }
        }
        throw new IllegalStateException("database does not return player id");
    }

    @Override
    public void updatePlayer(Connection connection, MonitorPlayer player) throws SQLException {
        try (PreparedStatement statement = connection.prepareStatement("UPDATE minecraft_players SET name = ? WHERE id = ?")) {
            statement.setString(1, player.name());
            statement.setInt(2, player.playerId());
            statement.executeUpdate();
        }
    }
}
