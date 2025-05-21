package net.okocraft.monitor.core.storage;

import net.okocraft.monitor.core.database.Database;
import net.okocraft.monitor.core.database.operator.Operators;
import net.okocraft.monitor.core.models.logs.PlayerChatLog;
import net.okocraft.monitor.core.models.logs.PlayerConnectLog;
import net.okocraft.monitor.core.models.logs.PlayerProxyCommandLog;
import net.okocraft.monitor.core.models.logs.PlayerRenameItemLog;
import net.okocraft.monitor.core.models.logs.PlayerWorldCommandLog;

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

    public void saveChatLogs(List<PlayerChatLog> list) throws SQLException {
        try (Connection connection = this.database.getConnection()) {
            this.operators.logs().insertPlayerChatLogs(connection, list);
        }
    }

    public void saveWorldCommandLogs(List<PlayerWorldCommandLog> list) throws SQLException {
        try (Connection connection = this.database.getConnection()) {
            this.operators.logs().insertPlayerWorldCommandLogs(connection, list);
        }
    }

    public void saveProxyCommandLogs(List<PlayerProxyCommandLog> list) throws SQLException {
        try (Connection connection = this.database.getConnection()) {
            this.operators.logs().insertPlayerProxyCommandLogs(connection, list);
        }
    }

    public void saveRenameItemLogs(List<PlayerRenameItemLog> list) throws SQLException {
        try (Connection connection = this.database.getConnection()) {
            this.operators.logs().insertPlayerRenameItemLogs(connection, list);
        }
    }
}
