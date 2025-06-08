package net.okocraft.monitor.core.database.mysql;

import com.zaxxer.hikari.HikariConfig;
import com.zaxxer.hikari.HikariDataSource;
import net.okocraft.monitor.core.database.Database;
import net.okocraft.monitor.core.database.mysql.operator.MySQLLogsTableOperator;
import net.okocraft.monitor.core.database.mysql.operator.MySQLPlayerNameHistoryTableOperator;
import net.okocraft.monitor.core.database.mysql.operator.MySQLPlayersTableOperator;
import net.okocraft.monitor.core.database.mysql.operator.MySQLServersTableOperator;
import net.okocraft.monitor.core.database.mysql.operator.MySQLUploadedObjectTableOperator;
import net.okocraft.monitor.core.database.mysql.operator.MySQLWorldsTableOperator;
import net.okocraft.monitor.core.database.operator.Operators;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.NotNullByDefault;
import org.jetbrains.annotations.Nullable;

import java.sql.Connection;
import java.sql.SQLException;
import java.util.Properties;
import java.util.concurrent.TimeUnit;

@NotNullByDefault
public class MySQLDatabase implements Database {

    private final MySQLSetting mySQLSetting;
    private final Operators operators;
    private @Nullable HikariDataSource hikariDataSource;

    public MySQLDatabase(MySQLSetting setting) {
        this.mySQLSetting = setting;
        this.operators = new Operators(
            new MySQLServersTableOperator(),
            new MySQLPlayersTableOperator(),
            new MySQLWorldsTableOperator(),
            new MySQLPlayerNameHistoryTableOperator(),
            new MySQLLogsTableOperator(),
            new MySQLUploadedObjectTableOperator()
        );
    }

    @Override
    public void prepare() throws Exception {
        if (this.mySQLSetting.username().isEmpty() || this.mySQLSetting.password().isEmpty()) {
            throw new Exception("MySQL database requires a username and password");
        }

        Class.forName("com.mysql.cj.jdbc.Driver"); // checks if the driver exists
        this.hikariDataSource = new HikariDataSource(this.createHikariConfig());
    }

    private @NotNull HikariConfig createHikariConfig() {
        var config = new HikariConfig();

        config.setJdbcUrl("jdbc:mysql://" + this.mySQLSetting.host() + ":" + this.mySQLSetting.port() + "/" + this.mySQLSetting.databaseName());
        config.setUsername(this.mySQLSetting.username());
        config.setPassword(this.mySQLSetting.password());

        config.setPoolName("MonitorMySQLPool");
        config.setDriverClassName("com.mysql.cj.jdbc.Driver");
        config.setConnectionTestQuery("SELECT 1");
        config.setMaxLifetime(TimeUnit.MINUTES.toMillis(30));
        config.setMaximumPoolSize(20);
        config.setMinimumIdle(3);
        config.setIdleTimeout(TimeUnit.SECONDS.toMillis(60));

        this.configureDataSourceProperties(config.getDataSourceProperties());

        return config;
    }

    @Override
    public void shutdown() {
        if (this.hikariDataSource != null) {
            this.hikariDataSource.close();
        }
    }

    @Override
    public @NotNull Connection getConnection() throws SQLException {
        if (this.hikariDataSource == null) {
            throw new IllegalStateException("HikariDataSource is not initialized.");
        }

        return this.hikariDataSource.getConnection();
    }

    @Override
    public Operators getOperators() {
        return this.operators;
    }

    private void configureDataSourceProperties(@NotNull Properties properties) {
        // https://github.com/brettwooldridge/HikariCP/wiki/Rapid-Recovery
        properties.putIfAbsent("socketTimeout", String.valueOf(TimeUnit.SECONDS.toMillis(30)));

        // https://github.com/brettwooldridge/HikariCP/wiki/MySQL-Configuration
        properties.putIfAbsent("cachePrepStmts", "true");
        properties.putIfAbsent("prepStmtCacheSize", "250");
        properties.putIfAbsent("prepStmtCacheSqlLimit", "2048");
        properties.putIfAbsent("useServerPrepStmts", "true");
        properties.putIfAbsent("useLocalSessionState", "true");
        properties.putIfAbsent("rewriteBatchedStatements", "true");
        properties.putIfAbsent("cacheResultSetMetadata", "true");
        properties.putIfAbsent("cacheServerConfiguration", "true");
        properties.putIfAbsent("elideSetAutoCommits", "true");
        properties.putIfAbsent("maintainTimeStats", "false");
        properties.putIfAbsent("alwaysSendSetIsolation", "false");
        properties.putIfAbsent("cacheCallableStmts", "true");
    }
}
