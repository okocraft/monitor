package net.okocraft.monitor.core.cloud.storage;

import dev.siroshun.codec4j.api.encoder.Encoder;
import dev.siroshun.jfun.result.Result;

import java.util.UUID;

public interface CloudStorage {

    String name();

    void prepare() throws Exception;

    <T> Result<Void, UploadError> upload(UUID uuid, Encoder<T> encoder, T object);

}
