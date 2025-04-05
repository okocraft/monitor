package net.okocraft.monitor.core.database.mysql.operator;

import net.okocraft.monitor.core.database.operator.WorldsTableOperator;
import net.okocraft.monitor.core.models.MonitorWorld;
import org.jetbrains.annotations.NotNullByDefault;
import org.jetbrains.annotations.Nullable;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;
import java.util.UUID;

@NotNullByDefault
public class MySQLWorldsTableOperator implements WorldsTableOperator {

    @Override
    public @Nullable MonitorWorld getWorldByUID(Connection connection, int serverId, UUID uid) throws SQLException {
        try (PreparedStatement statement = connection.prepareStatement("SELECT id, name FROM minecraft_worlds WHERE server_id = ? AND uid = ?")) {
            statement.setInt(1, serverId);
            statement.setBytes(2, MySQLUUID.uuidToBytes(uid));
            try (ResultSet resultSet = statement.executeQuery()) {
                if (resultSet.next()) {
                    return new MonitorWorld(
                        resultSet.getInt("id"),
                        serverId,
                        uid,
                        resultSet.getString("name")
                    );
                }
            }
        }
        return null;
    }

    @Override
    public int insertWorld(Connection connection, int serverId, UUID uuid, String name) throws SQLException {
        try (PreparedStatement statement = connection.prepareStatement("INSERT INTO minecraft_worlds (server_id, uid, name) VALUES (?, ?, ?)", Statement.RETURN_GENERATED_KEYS)) {
            statement.setInt(1, serverId);
            statement.setBytes(2, MySQLUUID.uuidToBytes(uuid));
            statement.setString(3, name);
            statement.executeUpdate();
            try (ResultSet resultSet = statement.getGeneratedKeys()) {
                if (resultSet.next()) {
                    return resultSet.getInt(1);
                }
            }
        }
        throw new IllegalStateException("database does not return world id");
    }

    @Override
    public void updateWorld(Connection connection, MonitorWorld world) throws SQLException {
        try (PreparedStatement statement = connection.prepareStatement("UPDATE minecraft_worlds SET name = ? WHERE id = ?")) {
            statement.setString(1, world.name());
            statement.setInt(2, world.worldId());
            statement.executeUpdate();
        }
    }
}
