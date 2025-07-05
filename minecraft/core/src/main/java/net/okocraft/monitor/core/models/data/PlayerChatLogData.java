package net.okocraft.monitor.core.models.data;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.codec.UUIDCodec;
import dev.siroshun.codec4j.api.encoder.Encoder;
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
        UUIDCodec.UUID_AS_STRING.toFieldEncoder("uuid", PlayerChatLogData::uuid),
        Codec.STRING.toFieldEncoder("name", PlayerChatLogData::name),
        Codec.STRING.toFieldEncoder("server_name", PlayerChatLogData::serverName),
        Codec.STRING.toFieldEncoder("world_name", PlayerChatLogData::worldName),
        BlockPosition.ENCODER.toFieldEncoder("position", PlayerChatLogData::position),
        Codec.STRING.toFieldEncoder("message", PlayerChatLogData::message),
        Codec.STRING.comap(DateTimeFormatter.ISO_LOCAL_DATE_TIME::format).toFieldEncoder("time", PlayerChatLogData::time)
    );

}
