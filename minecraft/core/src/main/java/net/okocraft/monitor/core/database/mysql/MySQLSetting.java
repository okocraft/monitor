package net.okocraft.monitor.core.database.mysql;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.decoder.Decoder;
import dev.siroshun.codec4j.api.decoder.object.ObjectDecoder;

public record MySQLSetting(
    String host, int port, String databaseName,
    String username, String password
) {

    public static final Decoder<MySQLSetting> CODEC = ObjectDecoder.create(
        MySQLSetting::new,
        Codec.STRING.toRequiredFieldDecoder("host"),
        Codec.INT.toRequiredFieldDecoder("port"),
        Codec.STRING.toRequiredFieldDecoder("database-name"),
        Codec.STRING.toRequiredFieldDecoder("username"),
        Codec.STRING.toRequiredFieldDecoder("password")
    );

}
