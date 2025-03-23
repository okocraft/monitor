package net.okocraft.monitor.core.database.operator;

import org.jetbrains.annotations.NotNullByDefault;

import java.sql.Connection;
import java.sql.SQLException;

@NotNullByDefault
public interface PlayerNameHistoryTableOperator {

    void insertHistory(Connection connection, int playerId, String playerName) throws SQLException;

}
