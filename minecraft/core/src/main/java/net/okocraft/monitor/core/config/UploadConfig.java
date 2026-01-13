package net.okocraft.monitor.core.config;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.decoder.Decoder;
import dev.siroshun.codec4j.api.decoder.object.FieldDecoder;
import dev.siroshun.codec4j.api.decoder.object.ObjectDecoder;

public record UploadConfig(SignConfig sign, CloudConfig cloud) {

    public static final Decoder<UploadConfig> CODEC = ObjectDecoder.create(
        UploadConfig::new,
        FieldDecoder.supplying("sign", SignConfig.CODEC, () -> new SignConfig("")),
        FieldDecoder.supplying("cloud", CloudConfig.CODEC, () -> new CloudConfig("", new CloudConfig.R2Config("", "", "", "")))
    );

    public static final UploadConfig EMPTY = new UploadConfig(new SignConfig(""), new CloudConfig("local", new CloudConfig.R2Config("", "", "", "")));

    public record CloudConfig(String type, R2Config r2) {

        public static final Decoder<CloudConfig> CODEC = ObjectDecoder.create(
            CloudConfig::new,
            FieldDecoder.required("type", Codec.STRING),
            FieldDecoder.supplying(
                "r2",
                ObjectDecoder.create(
                    R2Config::new,
                    FieldDecoder.required("account-id", Codec.STRING),
                    FieldDecoder.required("access-key", Codec.STRING),
                    FieldDecoder.required("secret-key", Codec.STRING),
                    FieldDecoder.required("bucket-name", Codec.STRING)
                ),
                () -> new R2Config("", "", "", "")
            )
        );

        public record R2Config(String accountId, String accessKey, String secretKey, String bucketName) {
        }
    }

    public record SignConfig(String secretKey) {

        public static final Decoder<SignConfig> CODEC = ObjectDecoder.create(
            SignConfig::new,
            FieldDecoder.required("secret-key", Codec.STRING)
        );

    }
}
