package net.okocraft.monitor.core.database.mysql.operator;

import net.okocraft.monitor.core.database.operator.ServersTableOperator;
import org.jetbrains.annotations.NotNullByDefault;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;
import java.sql.Timestamp;
import java.util.OptionalInt;

@NotNullByDefault
public class MySQLServersTableOperator implements ServersTableOperator {

    @Override
    public OptionalInt getServerIdByName(Connection connection, String serverName) throws SQLException {
        try (PreparedStatement statement = connection.prepareStatement("SELECT id FROM minecraft_servers WHERE name = ?")) {
            statement.setString(1, serverName);
            try (ResultSet resultSet = statement.executeQuery()) {
                if (resultSet.next()) {
                    return OptionalInt.of(resultSet.getInt("id"));
                }
            }
            return OptionalInt.empty();
        }
    }

    @Override
    public int insertNewServer(Connection connection, String serverName) throws SQLException {
        try (PreparedStatement statement = connection.prepareStatement("INSERT INTO minecraft_servers (name, created_at, updated_at) VALUES (?, ?, ?)", Statement.RETURN_GENERATED_KEYS)) {
            statement.setString(1, serverName);
            Timestamp now = MySQLDateTime.now();
            statement.setTimestamp(2, now);
            statement.setTimestamp(3, now);
            statement.executeUpdate();
            try (ResultSet resultSet = statement.getGeneratedKeys()) {
                if (resultSet.next()) {
                    return resultSet.getInt(1);
                }
            }
        }
        throw new IllegalStateException("database does not return server id");
    }

    @Override
    public void updateServerInfo(Connection connection, int serverId, String serverName) throws SQLException {
        try (PreparedStatement statement = connection.prepareStatement("UPDATE minecraft_servers SET name = ?, updated_at = ? WHERE id = ?")) {
            statement.setString(1, serverName);
            statement.setTimestamp(2, MySQLDateTime.now());
            statement.setInt(3, serverId);
            statement.execute();
        }
    }
}
