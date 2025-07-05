package net.okocraft.monitor.core.storage;

import net.okocraft.monitor.core.database.Database;
import net.okocraft.monitor.core.database.operator.Operators;
import net.okocraft.monitor.core.models.lookup.PlayerChatLogLookupParams;
import net.okocraft.monitor.core.models.data.PlayerChatLogData;
import net.okocraft.monitor.core.models.data.PlayerConnectLogData;
import net.okocraft.monitor.core.models.logs.PlayerChatLog;
import net.okocraft.monitor.core.models.logs.PlayerConnectLog;
import net.okocraft.monitor.core.models.logs.PlayerEditSignLog;
import net.okocraft.monitor.core.models.logs.PlayerProxyCommandLog;
import net.okocraft.monitor.core.models.logs.PlayerRenameItemLog;
import net.okocraft.monitor.core.models.logs.PlayerWorldCommandLog;
import net.okocraft.monitor.core.models.lookup.PlayerConnectLogLookupParams;

import java.sql.Connection;
import java.sql.SQLException;
import java.util.List;
import java.util.function.Consumer;

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

    public long countConnectLogs(PlayerConnectLogLookupParams params) throws SQLException {
        try (Connection connection = this.database.getConnection()) {
            return this.operators.logs().countPlayerConnectLogs(connection, params);
        }
    }

    public void lookupConnectLogData(PlayerConnectLogLookupParams params, Consumer<PlayerConnectLogData> consumer) throws SQLException {
        try (Connection connection = this.database.getConnection()) {
            this.operators.logs().selectPlayerConnectLogData(connection, params, consumer);
        }
    }

    public void saveChatLogs(List<PlayerChatLog> list) throws SQLException {
        try (Connection connection = this.database.getConnection()) {
            this.operators.logs().insertPlayerChatLogs(connection, list);
        }
    }

    public long countChatLogs(PlayerChatLogLookupParams params) throws SQLException {
        try (Connection connection = this.database.getConnection()) {
            return this.operators.logs().countPlayerChatLogs(connection, params);
        }
    }

    public void lookupChatLogData(PlayerChatLogLookupParams params, Consumer<PlayerChatLogData> consumer) throws SQLException {
        try (Connection connection = this.database.getConnection()) {
            this.operators.logs().selectPlayerChatLogData(connection, params, consumer);
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

    public void saveEditSignLog(List<PlayerEditSignLog> list) throws SQLException {
        try (Connection connection = this.database.getConnection()) {
            this.operators.logs().insertPlayerEditSignLogs(connection, list);
        }
    }
}
