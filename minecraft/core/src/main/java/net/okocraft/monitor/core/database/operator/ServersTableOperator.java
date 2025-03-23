package net.okocraft.monitor.core.database.operator;

import org.jetbrains.annotations.NotNullByDefault;

import java.sql.Connection;
import java.sql.SQLException;
import java.util.OptionalInt;

@NotNullByDefault
public interface ServersTableOperator {

    OptionalInt getServerIdByName(Connection connection, String serverName) throws SQLException;

    int insertNewServer(Connection connection, String serverName) throws SQLException;

    void updateServerInfo(Connection connection, int serverId, String serverName) throws SQLException;

}
