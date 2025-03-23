package net.okocraft.monitor.core.database.operator;

import net.okocraft.monitor.core.models.MonitorPlayer;
import org.jetbrains.annotations.NotNullByDefault;
import org.jetbrains.annotations.Nullable;

import java.sql.Connection;
import java.sql.SQLException;
import java.util.UUID;

@NotNullByDefault
public interface PlayersTableOperator {

    @Nullable MonitorPlayer getPlayerByUUID(Connection connection, UUID uuid) throws SQLException;

    int insertPlayer(Connection connection, UUID uuid, String name) throws SQLException;

    void updatePlayer(Connection connection, MonitorPlayer player) throws SQLException;
}
