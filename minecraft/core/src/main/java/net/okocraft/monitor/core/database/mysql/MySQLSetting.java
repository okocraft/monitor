package net.okocraft.monitor.core.database.mysql;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.decoder.Decoder;
import dev.siroshun.codec4j.api.decoder.object.FieldDecoder;
import dev.siroshun.codec4j.api.decoder.object.ObjectDecoder;

public record MySQLSetting(
    String host, int port, String databaseName,
    String username, String password
) {

    public static final Decoder<MySQLSetting> CODEC = ObjectDecoder.create(
        MySQLSetting::new,
        FieldDecoder.required("host", Codec.STRING),
        FieldDecoder.required("port", Codec.INT),
        FieldDecoder.required("database-name", Codec.STRING),
        FieldDecoder.required("username", Codec.STRING),
        FieldDecoder.required("password", Codec.STRING)
    );

}
