package net.okocraft.monitor.core.database.operator;

import net.okocraft.monitor.core.models.MonitorWorld;
import org.jetbrains.annotations.Nullable;

import java.sql.Connection;
import java.sql.SQLException;
import java.util.UUID;

public interface WorldsTableOperator {
    @Nullable MonitorWorld getWorldByUID(Connection connection, int serverId,  UUID uid) throws SQLException;

    int insertWorld(Connection connection, int serverId, UUID uid, String name) throws SQLException;

    void updateWorld(Connection connection, MonitorWorld world) throws SQLException;
}
