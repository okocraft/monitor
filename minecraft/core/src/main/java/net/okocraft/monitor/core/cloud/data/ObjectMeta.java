package net.okocraft.monitor.core.cloud.data;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.codec.EnumCodec;
import dev.siroshun.codec4j.api.codec.UUIDCodec;
import dev.siroshun.codec4j.api.encoder.Encoder;
import dev.siroshun.codec4j.api.encoder.object.ObjectEncoder;

import java.time.Instant;
import java.util.UUID;

public record ObjectMeta(UUID id, ObjectType type, int version, Instant expiresAt) {

    public static final int CURRENT_VERSION = 1;

    public static Encoder<ObjectMeta> ENCODER = ObjectEncoder.create(
        UUIDCodec.UUID_AS_STRING.toFieldEncoder("id", ObjectMeta::id),
        EnumCodec.byOrdinal(ObjectType.class).toFieldEncoder("type", ObjectMeta::type),
        Codec.INT.toFieldEncoder("version", ObjectMeta::version),
        Codec.LONG.comap(Instant::toEpochMilli).toFieldEncoder("expires_at", ObjectMeta::expiresAt)
    );

    public enum ObjectType {
        PLAYER_CONNECT_LOG,
        PLAYER_CHAT_LOG
    }
}
