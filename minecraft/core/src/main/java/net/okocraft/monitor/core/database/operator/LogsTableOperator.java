package net.okocraft.monitor.core.database.operator;

import net.okocraft.monitor.core.models.logs.PlayerChatLog;
import net.okocraft.monitor.core.models.logs.PlayerConnectLog;
import net.okocraft.monitor.core.models.logs.PlayerProxyCommandLog;
import net.okocraft.monitor.core.models.logs.PlayerWorldCommandLog;

import java.sql.Connection;
import java.sql.SQLException;
import java.util.List;

public interface LogsTableOperator {

    void insertPlayerConnectLogs(Connection connection, List<PlayerConnectLog> logs) throws SQLException;

    void insertPlayerChatLogs(Connection connection, List<PlayerChatLog> logs) throws SQLException;

    void insertPlayerWorldCommandLogs(Connection connection, List<PlayerWorldCommandLog> logs) throws SQLException;

    void insertPlayerProxyCommandLogs(Connection connection, List<PlayerProxyCommandLog> logs) throws SQLException;
}
