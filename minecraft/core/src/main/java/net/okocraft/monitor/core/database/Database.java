package net.okocraft.monitor.core.database;

import org.jetbrains.annotations.NotNull;

import java.sql.Connection;
import java.sql.SQLException;

public interface Database {

    void prepare() throws Exception;

    void shutdown() throws Exception;

    @NotNull
    Connection getConnection() throws SQLException;
}
