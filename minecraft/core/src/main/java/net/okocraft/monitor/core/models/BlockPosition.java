package net.okocraft.monitor.core.models;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.encoder.Encoder;
import dev.siroshun.codec4j.api.encoder.object.ObjectEncoder;

public record BlockPosition(int x, int y, int z) {

    public static final Encoder<BlockPosition> ENCODER = ObjectEncoder.create(
        Codec.INT.toFieldEncoder("x", BlockPosition::x),
        Codec.INT.toFieldEncoder("y", BlockPosition::y),
        Codec.INT.toFieldEncoder("z", BlockPosition::z)
    );
}
