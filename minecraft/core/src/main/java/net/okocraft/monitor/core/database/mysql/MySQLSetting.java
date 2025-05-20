package net.okocraft.monitor.core.database.mysql;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.codec.object.ObjectCodec;

public record MySQLSetting(
    String host, int port, String databaseName,
    String username, String password
) {

    public static final Codec<MySQLSetting> CODEC = ObjectCodec.create(
        MySQLSetting::new,
        Codec.STRING.toFieldCodec("host").required(MySQLSetting::host),
        Codec.INT.toFieldCodec("port").required(MySQLSetting::port),
        Codec.STRING.toFieldCodec("database-name").defaultValue("monitor_db").required(MySQLSetting::databaseName),
        Codec.STRING.toFieldCodec("username").required(MySQLSetting::username),
        Codec.STRING.toFieldCodec("password").required(MySQLSetting::password)
    );

}
