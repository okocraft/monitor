package net.okocraft.monitor.core.database.mysql.operator;

import net.okocraft.monitor.core.database.operator.PlayerNameHistoryTableOperator;
import org.jetbrains.annotations.NotNullByDefault;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.SQLException;

@NotNullByDefault
public class MySQLPlayerNameHistoryTableOperator implements PlayerNameHistoryTableOperator {

    @Override
    public void insertHistory(Connection connection, int playerId, String playerName) throws SQLException {
        try (PreparedStatement statement = connection.prepareStatement("INSERT INTO minecraft_player_name_histories (player_id, name, created_at) VALUE (?, ?, ?)")) {
            statement.setInt(1, playerId);
            statement.setString(2, playerName);
            statement.setTimestamp(3, MySQLDateTime.now());
            statement.execute();
        }
    }
}
