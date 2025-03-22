package net.okocraft.monitor.core.database.mysql;

public record MySQLSetting(
    String address,
    int port,
    String databaseName,
    String username,
    String password
) {
}
