package net.okocraft.monitor.core.models.data;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.codec.UUIDCodec;
import dev.siroshun.codec4j.api.encoder.Encoder;
import dev.siroshun.codec4j.api.encoder.object.FieldEncoder;
import dev.siroshun.codec4j.api.encoder.object.ObjectEncoder;
import net.okocraft.monitor.core.models.BlockPosition;
import org.jetbrains.annotations.NotNullByDefault;

import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.UUID;

@NotNullByDefault
public record PlayerChatLogData(UUID uuid, String name,
                                String serverName, String worldName,
                                BlockPosition position, String message, LocalDateTime time) {

    public static final Encoder<PlayerChatLogData> ENCODER = ObjectEncoder.create(
        FieldEncoder.create("uuid", UUIDCodec.UUID_AS_STRING, PlayerChatLogData::uuid, null),
        FieldEncoder.create("name", Codec.STRING, PlayerChatLogData::name, null),
        FieldEncoder.create("server_name", Codec.STRING, PlayerChatLogData::serverName, null),
        FieldEncoder.create("world_name", Codec.STRING, PlayerChatLogData::worldName, null),
        FieldEncoder.create("position", BlockPosition.ENCODER, PlayerChatLogData::position, null),
        FieldEncoder.create("message", Codec.STRING, PlayerChatLogData::message, null),
        FieldEncoder.create("time", Codec.STRING.comap(DateTimeFormatter.ISO_LOCAL_DATE_TIME::format), PlayerChatLogData::time, null)
    );

}
