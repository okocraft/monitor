package net.okocraft.monitor.core.config;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.codec.object.ObjectCodec;

public record UploadConfig(SignConfig sign) {

    public static final Codec<UploadConfig> CODEC = ObjectCodec.create(
        UploadConfig::new,
        SignConfig.CODEC.toFieldCodec("sign").defaultValueSupplier(() -> new SignConfig("")).required(UploadConfig::sign)
    );

    public static final UploadConfig EMPTY = new UploadConfig(new SignConfig(""));

    public record SignConfig(String secretKey) {

        public static final Codec<SignConfig> CODEC = ObjectCodec.create(
            SignConfig::new,
            Codec.STRING.toFieldCodec("secret-key").required(SignConfig::secretKey)
        );

    }
}
