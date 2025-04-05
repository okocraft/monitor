package net.okocraft.monitor.core.storage;

import net.okocraft.monitor.core.database.Database;
import net.okocraft.monitor.core.database.operator.Operators;
import net.okocraft.monitor.core.models.MonitorWorld;

import java.sql.SQLException;
import java.util.UUID;

public class WorldStorage {

    private final Database database;
    private final Operators operators;

    public WorldStorage(Database database) {
        this.database = database;
        this.operators = database.getOperators();
    }

    public MonitorWorld initializeWorld(int serverId, UUID uid, String name) throws SQLException {
        try (var connection = this.database.getConnection()) {
            MonitorWorld world = this.operators.worlds().getWorldByUID(connection, serverId, uid);

            if (world == null) {
                int worldId = this.operators.worlds().insertWorld(connection, serverId, uid, name);
                return new MonitorWorld(worldId, serverId, uid, name);
            }

            if (world.name().equals(name)) {
                return world;
            }

            MonitorWorld updatedWorld = new MonitorWorld(world.worldId(), serverId, uid, name);

            this.operators.worlds().updateWorld(connection, updatedWorld);

            return updatedWorld;
        }
    }
}
