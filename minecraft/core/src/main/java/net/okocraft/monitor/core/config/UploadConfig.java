package net.okocraft.monitor.core.config;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.decoder.Decoder;
import dev.siroshun.codec4j.api.decoder.object.ObjectDecoder;

public record UploadConfig(SignConfig sign, CloudConfig cloud) {

    public static final Decoder<UploadConfig> CODEC = ObjectDecoder.create(
        UploadConfig::new,
        SignConfig.CODEC.toSupplyingFieldDecoder("sign", () -> new SignConfig("")),
        CloudConfig.CODEC.toSupplyingFieldDecoder("cloud", () -> new CloudConfig("", new CloudConfig.R2Config("", "", "", "")))
    );

    public static final UploadConfig EMPTY = new UploadConfig(new SignConfig(""), new CloudConfig("local", new CloudConfig.R2Config("", "", "", "")));

    public record CloudConfig(String type, R2Config r2) {

        public static final Decoder<CloudConfig> CODEC = ObjectDecoder.create(
            CloudConfig::new,
            Codec.STRING.toRequiredFieldDecoder("type"),
            ObjectDecoder.create(
                R2Config::new,
                Codec.STRING.toRequiredFieldDecoder("account-id"),
                Codec.STRING.toRequiredFieldDecoder("access-key"),
                Codec.STRING.toRequiredFieldDecoder("secret-key"),
                Codec.STRING.toRequiredFieldDecoder("bucket-name")
            ).toSupplyingFieldDecoder("r2", () -> new R2Config("", "", "", ""))
        );

        public record R2Config(String accountId, String accessKey, String secretKey, String bucketName) {
        }
    }

    public record SignConfig(String secretKey) {

        public static final Decoder<SignConfig> CODEC = ObjectDecoder.create(
            SignConfig::new,
            Codec.STRING.toRequiredFieldDecoder("secret-key")
        );

    }
}
