package net.okocraft.monitor.core.database.mysql;

import dev.siroshun.configapi.core.serialization.annotation.DefaultInt;
import dev.siroshun.configapi.core.serialization.annotation.DefaultString;

public record MySQLSetting(
    @DefaultString("address") String address,
    @DefaultInt(3306) int port,
    @DefaultString("monitor_db") String databaseName,
    String username,
    String password
) {
}
