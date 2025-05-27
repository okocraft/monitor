package net.okocraft.monitor.core.models.data;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.codec.UUIDCodec;
import dev.siroshun.codec4j.api.encoder.Encoder;
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
        UUIDCodec.UUID_AS_STRING.toFieldEncoder("uuid", PlayerConnectLogData::uuid),
        Codec.STRING.toFieldEncoder("name", PlayerConnectLogData::name),
        Codec.STRING.toFieldEncoder("server_name", PlayerConnectLogData::serverName),
        Codec.STRING.comap(PlayerConnectLog.Action::name).toFieldEncoder("action", PlayerConnectLogData::action),
        Codec.STRING.toFieldEncoder("address", PlayerConnectLogData::address),
        Codec.STRING.toFieldEncoder("reason", PlayerConnectLogData::reason),
        Codec.STRING.comap(DateTimeFormatter.ISO_LOCAL_DATE_TIME::format).toFieldEncoder("time", PlayerConnectLogData::time)
    );

    public record LookupParams(LocalDateTime start, LocalDateTime end) {
    }

}
