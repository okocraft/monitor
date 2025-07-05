package net.okocraft.monitor.core.database.operator;

import net.okocraft.monitor.core.models.data.PlayerChatLogData;
import net.okocraft.monitor.core.models.data.PlayerConnectLogData;
import net.okocraft.monitor.core.models.logs.PlayerChatLog;
import net.okocraft.monitor.core.models.logs.PlayerConnectLog;
import net.okocraft.monitor.core.models.logs.PlayerEditSignLog;
import net.okocraft.monitor.core.models.logs.PlayerProxyCommandLog;
import net.okocraft.monitor.core.models.logs.PlayerRenameItemLog;
import net.okocraft.monitor.core.models.logs.PlayerWorldCommandLog;
import net.okocraft.monitor.core.models.lookup.PlayerChatLogLookupParams;
import net.okocraft.monitor.core.models.lookup.PlayerConnectLogLookupParams;

import java.sql.Connection;
import java.sql.SQLException;
import java.util.List;
import java.util.function.Consumer;

public interface LogsTableOperator {

    void insertPlayerConnectLogs(Connection connection, List<PlayerConnectLog> logs) throws SQLException;

    long countPlayerConnectLogs(Connection connection, PlayerConnectLogLookupParams params) throws SQLException;

    void selectPlayerConnectLogData(Connection connection, PlayerConnectLogLookupParams params, Consumer<PlayerConnectLogData> consumer) throws SQLException;

    void insertPlayerChatLogs(Connection connection, List<PlayerChatLog> logs) throws SQLException;

    long countPlayerChatLogs(Connection connection, PlayerChatLogLookupParams params) throws SQLException;

    void selectPlayerChatLogData(Connection connection, PlayerChatLogLookupParams params, Consumer<PlayerChatLogData> consumer) throws SQLException;

    void insertPlayerWorldCommandLogs(Connection connection, List<PlayerWorldCommandLog> logs) throws SQLException;

    void insertPlayerProxyCommandLogs(Connection connection, List<PlayerProxyCommandLog> logs) throws SQLException;

    void insertPlayerRenameItemLogs(Connection connection, List<PlayerRenameItemLog> logs) throws SQLException;

    void insertPlayerEditSignLogs(Connection connection, List<PlayerEditSignLog> logs) throws SQLException;

}
