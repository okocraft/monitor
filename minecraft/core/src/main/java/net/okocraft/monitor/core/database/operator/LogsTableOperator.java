package net.okocraft.monitor.core.database.operator;

import net.okocraft.monitor.core.models.logs.PlayerConnectLog;

import java.sql.Connection;
import java.sql.SQLException;
import java.util.List;

public interface LogsTableOperator {

    void insertPlayerConnectLogs(Connection connection, List<PlayerConnectLog> logs) throws SQLException;

}
