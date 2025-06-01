package net.okocraft.monitor.core.config;

import dev.siroshun.codec4j.api.decoder.Decoder;
import dev.siroshun.codec4j.api.decoder.object.ObjectDecoder;
import net.okocraft.monitor.core.database.mysql.MySQLSetting;

public record DatabaseConfig(MySQLSetting mysql) {

    public static final Decoder<DatabaseConfig> CODEC = ObjectDecoder.create(
        DatabaseConfig::new,
        MySQLSetting.CODEC.toSupplyingFieldDecoder("mysql", () -> new MySQLSetting("", 0, "", "", ""))
    );

}
