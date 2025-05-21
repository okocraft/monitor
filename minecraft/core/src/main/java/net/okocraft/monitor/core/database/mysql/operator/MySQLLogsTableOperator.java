package net.okocraft.monitor.core.database.mysql.operator;

import net.kyori.adventure.text.serializer.gson.GsonComponentSerializer;
import net.kyori.adventure.text.serializer.plain.PlainTextComponentSerializer;
import net.okocraft.monitor.core.database.operator.LogsTableOperator;
import net.okocraft.monitor.core.models.logs.PlayerChatLog;
import net.okocraft.monitor.core.models.logs.PlayerConnectLog;
import net.okocraft.monitor.core.models.logs.PlayerProxyCommandLog;
import net.okocraft.monitor.core.models.logs.PlayerRenameItemLog;
import net.okocraft.monitor.core.models.logs.PlayerWorldCommandLog;

import java.io.ByteArrayInputStream;
import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.SQLException;
import java.util.List;

public class MySQLLogsTableOperator implements LogsTableOperator {

    private static final MySQLBulkInserter PLAYER_CONNECT_LOG_INSERTER = MySQLBulkInserter.create("minecraft_player_connect_logs", List.of("player_id", "server_id", "action", "address", "reason", "created_at"));
    private static final MySQLBulkInserter PLAYER_CHAT_LOG_INSERTER = MySQLBulkInserter.create("minecraft_player_chat_logs", List.of("player_id", "world_id", "position_x", "position_y", "position_z", "message", "created_at"));
    private static final MySQLBulkInserter PLAYER_WORLD_COMMAND_LOG_INSERTER = MySQLBulkInserter.create("minecraft_player_world_command_logs", List.of("player_id", "world_id", "position_x", "position_y", "position_z", "command", "created_at"));
    private static final MySQLBulkInserter PLAYER_PROXY_COMMAND_LOG_INSERTER = MySQLBulkInserter.create("minecraft_player_proxy_command_logs", List.of("player_id", "server_id", "command", "created_at"));
    private static final MySQLBulkInserter PLAYER_RENAME_ITEM_LOG_INSERTER = MySQLBulkInserter.create("minecraft_player_rename_item_logs", List.of("player_id", "world_id", "position_x", "position_y", "position_z", "item_type", "item_name", "item_name_component", "amount", "created_at"));

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
                statement.setTimestamp(parameterIndex++, MySQLDateTime.from(log.time()));
            }
            statement.executeUpdate();
        }
    }

    @Override
    public void insertPlayerChatLogs(Connection connection, List<PlayerChatLog> logs) throws SQLException {
        try (PreparedStatement statement = connection.prepareStatement(PLAYER_CHAT_LOG_INSERTER.createQuery(logs.size()))) {
            int parameterIndex = 1;
            for (PlayerChatLog log : logs) {
                statement.setInt(parameterIndex++, log.playerId());
                statement.setInt(parameterIndex++, log.worldId());
                statement.setInt(parameterIndex++, log.position().x());
                statement.setInt(parameterIndex++, log.position().y());
                statement.setInt(parameterIndex++, log.position().z());
                statement.setString(parameterIndex++, log.message().substring(0, Math.min(65535, log.message().length())));
                statement.setTimestamp(parameterIndex++, MySQLDateTime.from(log.time()));
            }
            statement.executeUpdate();
        }
    }

    @Override
    public void insertPlayerWorldCommandLogs(Connection connection, List<PlayerWorldCommandLog> logs) throws SQLException {
        try (PreparedStatement statement = connection.prepareStatement(PLAYER_WORLD_COMMAND_LOG_INSERTER.createQuery(logs.size()))) {
            int parameterIndex = 1;
            for (PlayerWorldCommandLog log : logs) {
                statement.setInt(parameterIndex++, log.playerId());
                statement.setInt(parameterIndex++, log.worldId());
                statement.setInt(parameterIndex++, log.position().x());
                statement.setInt(parameterIndex++, log.position().y());
                statement.setInt(parameterIndex++, log.position().z());
                statement.setString(parameterIndex++, log.command().substring(0, Math.min(65535, log.command().length())));
                statement.setTimestamp(parameterIndex++, MySQLDateTime.from(log.time()));
            }
            statement.executeUpdate();
        }
    }

    @Override
    public void insertPlayerProxyCommandLogs(Connection connection, List<PlayerProxyCommandLog> logs) throws SQLException {
        try (PreparedStatement statement = connection.prepareStatement(PLAYER_PROXY_COMMAND_LOG_INSERTER.createQuery(logs.size()))) {
            int parameterIndex = 1;
            for (PlayerProxyCommandLog log : logs) {
                statement.setInt(parameterIndex++, log.playerId());
                statement.setInt(parameterIndex++, log.serverId());
                statement.setString(parameterIndex++, log.command().substring(0, Math.min(65535, log.command().length())));
                statement.setTimestamp(parameterIndex++, MySQLDateTime.from(log.time()));
            }
            statement.executeUpdate();
        }
    }

    @Override
    public void insertPlayerRenameItemLogs(Connection connection, List<PlayerRenameItemLog> logs) throws SQLException {
        try (PreparedStatement statement = connection.prepareStatement(PLAYER_RENAME_ITEM_LOG_INSERTER.createQuery(logs.size()))) {
            int parameterIndex = 1;
            for (PlayerRenameItemLog log : logs) {
                statement.setInt(parameterIndex++, log.playerId());
                statement.setInt(parameterIndex++, log.worldId());
                statement.setInt(parameterIndex++, log.position().x());
                statement.setInt(parameterIndex++, log.position().y());
                statement.setInt(parameterIndex++, log.position().z());
                statement.setString(parameterIndex++, log.itemType().asMinimalString());
                statement.setString(parameterIndex++, PlainTextComponentSerializer.plainText().serialize(log.itemName()));
                String itemNameJson = GsonComponentSerializer.gson().serialize(log.itemName());
                try (ByteArrayInputStream in = new ByteArrayInputStream(itemNameJson.getBytes(StandardCharsets.UTF_8))) {
                    statement.setBinaryStream(parameterIndex++, in);
                } catch (IOException ignored) {
                    // ByteArrayInputStream never throws IOException
                }
                statement.setInt(parameterIndex++, log.amount());
                statement.setTimestamp(parameterIndex++, MySQLDateTime.from(log.time()));
            }
            statement.executeUpdate();
        }
    }
}
