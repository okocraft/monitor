package net.okocraft.monitor.core.cloud.sign;

import dev.siroshun.codec4j.api.encoder.Encoder;
import dev.siroshun.codec4j.api.error.EncodeError;
import dev.siroshun.codec4j.io.gson.GsonIO;
import dev.siroshun.jfun.result.Result;
import net.okocraft.monitor.core.cloud.data.SignedData;
import org.jetbrains.annotations.Nullable;

import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.nio.charset.StandardCharsets;
import java.security.Key;

public class HmacSigner {

    private static final String ALGORITHM = "HmacSHA256";

    public static HmacSigner create(String secretKey) {
        Key key = secretKey.isEmpty() ? null : new SecretKeySpec(secretKey.getBytes(StandardCharsets.UTF_8), ALGORITHM);
        return new HmacSigner(key);
    }

    private final @Nullable Key key;

    private HmacSigner(@Nullable Key key) {
        this.key = key;
    }

    public boolean hasKey() {
        return this.key != null;
    }

    public <T> Result<SignedData<T>, EncodeError> sign(T data, Encoder<T> encoder) {
        Result<String, EncodeError> encodeResult = GsonIO.DEFAULT.encodeToString(encoder, data);

        Result<byte[], EncodeError> hmacResult = this.generateHmac(encodeResult.unwrap().getBytes(StandardCharsets.UTF_8));
        if (hmacResult.isFailure()) {
            return hmacResult.asFailure();
        }

        return Result.success(new SignedData<>(data, encodeResult.unwrap(), hmacResult.unwrap()));
    }

    private Result<byte[], EncodeError> generateHmac(byte[] data) {
        if (this.key == null) {
            return Result.failure();
        }

        Mac mac;
        try {
            mac = Mac.getInstance(ALGORITHM);
            mac.init(this.key);
        } catch (Exception e) {
            return EncodeError.fatalError(e).asFailure();
        }

        return Result.success(mac.doFinal(data));
    }
}
