package net.okocraft.monitor.core.config;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.codec.object.ObjectCodec;
import net.okocraft.monitor.core.database.mysql.MySQLSetting;

public record DatabaseConfig(MySQLSetting mysql) {

    public static final Codec<DatabaseConfig> CODEC = ObjectCodec.create(
        DatabaseConfig::new,
        MySQLSetting.CODEC.toFieldCodec("mysql").defaultValueSupplier(() -> new MySQLSetting("", 0, "", "", "")).required(DatabaseConfig::mysql)
    );

}
