package net.okocraft.monitor.core.database.operator;

import net.okocraft.monitor.core.models.PlayerConnectLog;

import java.sql.Connection;
import java.sql.SQLException;
import java.util.List;

public interface PlayerConnectLogsTableOperator {

    void insertLogs(Connection connection, List<PlayerConnectLog> logs) throws SQLException;

}
