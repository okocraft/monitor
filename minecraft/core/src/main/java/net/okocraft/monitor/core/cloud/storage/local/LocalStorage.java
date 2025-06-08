package net.okocraft.monitor.core.cloud.storage.local;

import dev.siroshun.codec4j.api.encoder.Encoder;
import dev.siroshun.codec4j.api.file.DefaultOpenOptions;
import dev.siroshun.codec4j.io.gson.GsonIO;
import dev.siroshun.jfun.result.Result;
import net.okocraft.monitor.core.cloud.storage.CloudStorage;
import net.okocraft.monitor.core.cloud.storage.UploadError;

import java.io.BufferedOutputStream;
import java.io.IOException;
import java.io.OutputStream;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.zip.Deflater;
import java.util.zip.GZIPOutputStream;

public class LocalStorage implements CloudStorage {

    private static final int BUFFER_SIZE = 1 << 18;
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

        Path parent = filepath.getParent();
        if (parent != null && !Files.isDirectory(parent)) {
            try {
                Files.createDirectories(parent);
            } catch (IOException e) {
                return Result.failure(new UploadError.UploadException(e));
            }
        }

        try (OutputStream out = Files.newOutputStream(filepath, DefaultOpenOptions.fileOpenOptions());
             BufferedOutputStream bufferedOut = new BufferedOutputStream(out, BUFFER_SIZE);
             GZIPOutputStream gzipOut = new GZIPOutputStream(bufferedOut) {
                 {
                     this.def.setLevel(Deflater.BEST_COMPRESSION);
                 }
             }
        ) {
            return GsonIO.DEFAULT.encodeTo(gzipOut, encoder, object).mapError(UploadError.EncodeError::new);
        } catch (IOException e) {
            return Result.failure(new UploadError.UploadException(e));
        }
    }

}
