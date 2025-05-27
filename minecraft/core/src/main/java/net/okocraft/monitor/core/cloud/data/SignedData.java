package net.okocraft.monitor.core.cloud.data;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.encoder.Encoder;
import dev.siroshun.codec4j.api.encoder.object.ObjectEncoder;

import java.util.HexFormat;

public record SignedData<T>(T original, String jsonData, byte[] signature) {

    public static final Encoder<SignedData<?>> ENCODER_WITHOUT_META = ObjectEncoder.create(
        Codec.STRING.toFieldEncoder("data", SignedData::jsonData),
        Codec.STRING.<byte[]>comap(b -> HexFormat.of().formatHex(b)).toFieldEncoder("signature", SignedData::signature)
    );

}
