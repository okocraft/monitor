package net.okocraft.monitor.core.cloud.data;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.codec.EnumCodec;
import dev.siroshun.codec4j.api.codec.UUIDCodec;
import dev.siroshun.codec4j.api.encoder.Encoder;
import dev.siroshun.codec4j.api.encoder.object.FieldEncoder;
import dev.siroshun.codec4j.api.encoder.object.ObjectEncoder;

import java.time.Instant;
import java.util.UUID;

public record ObjectMeta(UUID id, ObjectType type, int version, Instant expiresAt) {

    public static final int CURRENT_VERSION = 1;

    public static Encoder<ObjectMeta> ENCODER = ObjectEncoder.create(
        FieldEncoder.create("id", UUIDCodec.UUID_AS_STRING, ObjectMeta::id, null),
        FieldEncoder.create("type", EnumCodec.byOrdinal(ObjectType.class), ObjectMeta::type, null),
        FieldEncoder.create("version", Codec.INT, ObjectMeta::version, null),
        FieldEncoder.create("expires_at", Codec.LONG.comap(Instant::toEpochMilli), ObjectMeta::expiresAt, null)
    );

    public enum ObjectType {
        PLAYER_CONNECT_LOG,
        PLAYER_CHAT_LOG
    }
}
