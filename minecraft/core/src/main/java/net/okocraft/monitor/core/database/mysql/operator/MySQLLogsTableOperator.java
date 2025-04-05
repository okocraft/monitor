package net.okocraft.monitor.core.database.mysql.operator;

import net.okocraft.monitor.core.database.operator.LogsTableOperator;
import net.okocraft.monitor.core.models.logs.PlayerConnectLog;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.SQLException;
import java.util.List;

public class MySQLLogsTableOperator implements LogsTableOperator {

    private static final MySQLBulkInserter PLAYER_CONNECT_LOG_INSERTER = MySQLBulkInserter.create("minecraft_player_connect_logs", List.of("player_id", "server_id", "action", "address", "reason", "created_at"));

    @Override
    public void insertPlayerConnectLogs(Connection connection, List<PlayerConnectLog> logs) throws SQLException {
        try (PreparedStatement statement = connection.prepareStatement(PLAYER_CONNECT_LOG_INSERTER.createQuery(logs.size()))) {
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
