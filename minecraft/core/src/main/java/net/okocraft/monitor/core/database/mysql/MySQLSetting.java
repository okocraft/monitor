package net.okocraft.monitor.core.database.mysql;

import dev.siroshun.configapi.core.serialization.annotation.DefaultInt;
import dev.siroshun.configapi.core.serialization.annotation.DefaultString;

public record MySQLSetting(
    @DefaultString("localhost") String host,
    @DefaultInt(3306) int port,
    @DefaultString("monitor_db") String databaseName,
    String username,
    String password
) {
}
