package net.okocraft.monitor.core.models;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.encoder.Encoder;
import dev.siroshun.codec4j.api.encoder.object.FieldEncoder;
import dev.siroshun.codec4j.api.encoder.object.ObjectEncoder;

public record BlockPosition(int x, int y, int z) {

    public static final Encoder<BlockPosition> ENCODER = ObjectEncoder.create(
        FieldEncoder.create("x", Codec.INT, BlockPosition::x, null),
        FieldEncoder.create("y", Codec.INT, BlockPosition::y, null),
        FieldEncoder.create("z", Codec.INT, BlockPosition::z, null)
    );
}
