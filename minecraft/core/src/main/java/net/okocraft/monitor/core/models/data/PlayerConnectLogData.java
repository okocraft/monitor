package net.okocraft.monitor.core.models.data;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.codec.UUIDCodec;
import dev.siroshun.codec4j.api.encoder.Encoder;
import dev.siroshun.codec4j.api.encoder.object.FieldEncoder;
import dev.siroshun.codec4j.api.encoder.object.ObjectEncoder;
import net.okocraft.monitor.core.models.logs.PlayerConnectLog;
import org.jetbrains.annotations.NotNullByDefault;

import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.UUID;

@NotNullByDefault
public record PlayerConnectLogData(UUID uuid, String name, String serverName, PlayerConnectLog.Action action,
                                   String address, String reason, LocalDateTime time) {

    public static final Encoder<PlayerConnectLogData> ENCODER = ObjectEncoder.create(
        FieldEncoder.create("uuid", UUIDCodec.UUID_AS_STRING, PlayerConnectLogData::uuid, null),
        FieldEncoder.create("name", Codec.STRING, PlayerConnectLogData::name, null),
        FieldEncoder.create("server_name", Codec.STRING, PlayerConnectLogData::serverName, null),
        FieldEncoder.create("action", Codec.STRING.comap(PlayerConnectLog.Action::name), PlayerConnectLogData::action, null),
        FieldEncoder.create("address", Codec.STRING, PlayerConnectLogData::address, null),
        FieldEncoder.create("reason", Codec.STRING, PlayerConnectLogData::reason, null),
        FieldEncoder.create("time", Codec.STRING.comap(DateTimeFormatter.ISO_LOCAL_DATE_TIME::format), PlayerConnectLogData::time, null)
    );

}
