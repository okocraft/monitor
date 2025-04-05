package net.okocraft.monitor.core.database.mysql.operator;

import net.okocraft.monitor.core.database.operator.PlayerConnectLogsTableOperator;
import net.okocraft.monitor.core.models.PlayerConnectLog;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.SQLException;
import java.util.List;

public class MySQLPlayerConnectLogsTableOperator implements PlayerConnectLogsTableOperator {

    private static final MySQLBulkInserter INSERTER = MySQLBulkInserter.create("minecraft_player_connect_logs", List.of("player_id", "server_id", "action", "address", "reason", "created_at"));

    @Override
    public void insertLogs(Connection connection, List<PlayerConnectLog> logs) throws SQLException {
        try (PreparedStatement statement = connection.prepareStatement(INSERTER.createQuery(logs.size()))) {
            int parameterIndex = 1;
            for (PlayerConnectLog log : logs) {
                statement.setInt(parameterIndex++, log.playerId());
                statement.setInt(parameterIndex++, log.serverId());
                statement.setInt(parameterIndex++, log.action().id());
                statement.setString(parameterIndex++, log.address().substring(0, Math.min(64, log.address().length())));
                statement.setString(parameterIndex++, log.reason());
                statement.setTimestamp(parameterIndex++, MySQLDateTime.now());
            }
            statement.executeUpdate();
        }
    }
}
