package net.okocraft.monitor.core.cloud.storage.local;

import dev.siroshun.codec4j.api.encoder.Encoder;
import dev.siroshun.codec4j.io.gson.GsonIO;
import dev.siroshun.codec4j.io.gzip.GzipIO;
import dev.siroshun.jfun.result.Result;
import net.okocraft.monitor.core.cloud.storage.CloudStorage;
import net.okocraft.monitor.core.cloud.storage.UploadError;

import java.nio.file.Files;
import java.nio.file.Path;

public class LocalStorage implements CloudStorage {

    private static final String FILE_EXTENSION = ".json.gz";

    private final Path directory;

    public LocalStorage(Path directory) {
        this.directory = directory;
    }

    @Override
    public String name() {
        return "local";
    }

    @Override
    public void prepare() throws Exception {
        Files.createDirectories(this.directory);
    }

    @Override
    public void shutdown() {
    }

    @Override
    public <T> Result<Void, UploadError> upload(String key, Encoder<T> encoder, T object) {
        Path filepath = this.directory.resolve(key + FILE_EXTENSION);
        return GzipIO.bestCompression(GsonIO.DEFAULT).encodeTo(filepath, encoder, object).mapError(UploadError.EncodeError::new);
    }

}
