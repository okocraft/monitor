package net.okocraft.monitor.core.cloud.data;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.encoder.Encoder;
import dev.siroshun.codec4j.api.encoder.object.FieldEncoder;
import dev.siroshun.codec4j.api.encoder.object.ObjectEncoder;

import java.util.HexFormat;

public record SignedData<T>(T original, String jsonData, byte[] signature) {

    public static final Encoder<SignedData<?>> ENCODER_WITHOUT_META = ObjectEncoder.create(
        FieldEncoder.create("data", Codec.STRING, SignedData::jsonData, null),
        FieldEncoder.create("signature", Codec.STRING.comap(b -> HexFormat.of().formatHex(b)), SignedData::signature, null)
    );

}
