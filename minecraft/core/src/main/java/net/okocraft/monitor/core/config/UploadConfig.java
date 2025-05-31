package net.okocraft.monitor.core.config;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.codec.object.ObjectCodec;

public record UploadConfig(SignConfig sign, CloudConfig cloud) {

    public static final Codec<UploadConfig> CODEC = ObjectCodec.create(
        UploadConfig::new,
        SignConfig.CODEC.toFieldCodec("sign").defaultValueSupplier(() -> new SignConfig("")).required(UploadConfig::sign),
        CloudConfig.CODEC.toFieldCodec("cloud").defaultValueSupplier(() -> new CloudConfig("", new CloudConfig.R2Config("", "", "", ""))).required(UploadConfig::cloud)
    );

    public static final UploadConfig EMPTY = new UploadConfig(new SignConfig(""), new CloudConfig("local", new CloudConfig.R2Config("", "", "", "")));

    public record CloudConfig(String type, R2Config r2) {

        public static final Codec<CloudConfig> CODEC = ObjectCodec.create(
            CloudConfig::new,
            Codec.STRING.toFieldCodec("type").defaultValue("").required(CloudConfig::type),
            ObjectCodec.create(
                R2Config::new,
                Codec.STRING.toFieldCodec("account-id").defaultValue("").required(R2Config::accountId),
                Codec.STRING.toFieldCodec("access-key").defaultValue("").required(R2Config::accessKey),
                Codec.STRING.toFieldCodec("secret-key").defaultValue("").required(R2Config::secretKey),
                Codec.STRING.toFieldCodec("bucket-name").defaultValue("").required(R2Config::bucketName)
            ).toFieldCodec("r2").defaultValueSupplier(() -> new R2Config("", "", "", "")).required(CloudConfig::r2)
        );

        public record R2Config(String accountId, String accessKey, String secretKey, String bucketName) {
        }
    }

    public record SignConfig(String secretKey) {

        public static final Codec<SignConfig> CODEC = ObjectCodec.create(
            SignConfig::new,
            Codec.STRING.toFieldCodec("secret-key").required(SignConfig::secretKey)
        );

    }
}
