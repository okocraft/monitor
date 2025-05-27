package net.okocraft.monitor.core.cloud.data;

import dev.siroshun.codec4j.api.codec.Base64Codec;
import dev.siroshun.codec4j.api.encoder.Encoder;
import dev.siroshun.codec4j.api.encoder.object.ObjectEncoder;

public record SignedData<T>(T original, byte[] jsonData, byte[] signature) {

    public static final Encoder<SignedData<?>> ENCODER_WITHOUT_META = ObjectEncoder.create(
        Base64Codec.CODEC.toFieldEncoder("data", SignedData::jsonData),
        Base64Codec.CODEC.toFieldEncoder("signature", SignedData::signature)
    );

}
