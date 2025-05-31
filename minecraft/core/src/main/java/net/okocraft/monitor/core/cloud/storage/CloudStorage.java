package net.okocraft.monitor.core.cloud.storage;

import dev.siroshun.codec4j.api.encoder.Encoder;
import dev.siroshun.jfun.result.Result;

public interface CloudStorage {

    String name();

    void prepare() throws Exception;

    void shutdown() throws Exception;

    <T> Result<Void, UploadError> upload(String key, Encoder<T> encoder, T object);

}
