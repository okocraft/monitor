package net.okocraft.monitor.core.cloud.sign;

import dev.siroshun.codec4j.api.encoder.Encoder;
import dev.siroshun.codec4j.api.error.EncodeError;
import dev.siroshun.codec4j.io.gson.GsonIO;
import dev.siroshun.jfun.result.Result;
import net.okocraft.monitor.core.cloud.data.SignedData;

import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.security.Key;

public class HmacSigner {

    private static final String ALGORITHM = "HmacSHA256";

    public static HmacSigner create(String secretKey) {
        Key key = new SecretKeySpec(secretKey.getBytes(StandardCharsets.UTF_8), ALGORITHM);
        return new HmacSigner(key);
    }

    private final Key key;

    private HmacSigner(Key key) {
        this.key = key;
    }

    public <T> Result<SignedData<T>, EncodeError> sign(T data, Encoder<T> encoder) {
        Result<byte[], EncodeError> encodeResult = this.encode(data, encoder);
        if (encodeResult.isFailure()) {
            return encodeResult.asFailure();
        }

        Result<byte[], EncodeError> hmacResult = this.generateHmac(encodeResult.unwrap());
        if (hmacResult.isFailure()) {
            return hmacResult.asFailure();
        }

        return Result.success(new SignedData<>(data, encodeResult.unwrap(), hmacResult.unwrap()));
    }

    private <T> Result<byte[], EncodeError> encode(T data, Encoder<T> encoder) {
        try (ByteArrayOutputStream out = new ByteArrayOutputStream()) {
            Result<Void, EncodeError> encodeResult = GsonIO.DEFAULT.encodeTo(out, encoder, data);
            if (encodeResult.isFailure()) {
                return encodeResult.asFailure();
            }
            return Result.success(out.toByteArray());
        } catch (IOException e) {
            return EncodeError.fatalError(e).asFailure();
        }
    }

    private Result<byte[], EncodeError> generateHmac(byte[] data) {
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
