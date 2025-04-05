package net.okocraft.monitor.core.storage;

import net.okocraft.monitor.core.database.Database;
import net.okocraft.monitor.core.database.operator.Operators;
import net.okocraft.monitor.core.models.logs.PlayerConnectLog;

import java.sql.Connection;
import java.sql.SQLException;
import java.util.List;

public class PlayerLogStorage {

    private final Database database;
    private final Operators operators;

    public PlayerLogStorage(Database database) {
        this.database = database;
        this.operators = database.getOperators();
    }

    public void saveConnectLogs(List<PlayerConnectLog> list) throws SQLException {
        try (Connection connection = this.database.getConnection()) {
            this.operators.logs().insertPlayerConnectLogs(connection, list);
        }
    }
}
