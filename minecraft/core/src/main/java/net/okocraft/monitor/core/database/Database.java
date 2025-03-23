package net.okocraft.monitor.core.database;

import net.okocraft.monitor.core.database.operator.Operators;
import org.jetbrains.annotations.NotNullByDefault;

import java.sql.Connection;
import java.sql.SQLException;

@NotNullByDefault
public interface Database {

    void prepare() throws Exception;

    void shutdown() throws Exception;

    Connection getConnection() throws SQLException;

    Operators getOperators();
}
