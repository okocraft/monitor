package net.okocraft.monitor.core.cloud.storage.cloudflare;

import dev.siroshun.codec4j.api.encoder.Encoder;
import dev.siroshun.codec4j.io.gson.GsonIO;
import dev.siroshun.jfun.result.Result;
import io.minio.MinioClient;
import io.minio.UploadObjectArgs;
import net.okocraft.monitor.core.cloud.storage.CloudStorage;
import net.okocraft.monitor.core.cloud.storage.UploadError;
import org.jetbrains.annotations.NotNullByDefault;

import java.io.BufferedOutputStream;
import java.io.IOException;
import java.io.OutputStream;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.zip.Deflater;
import java.util.zip.GZIPOutputStream;

@NotNullByDefault
public class CloudflareR2Storage implements CloudStorage {

    private static final int BUFFER_SIZE = 1 << 18;

    public static CloudflareR2Storage create(String accountId, String accessKeyId, String secretAccessKey, String bucketName, Path tmpDirectory) {
        String endpoint = String.format("https://%s.r2.cloudflarestorage.com", accountId);
        MinioClient client = MinioClient.builder()
            .endpoint(endpoint)
            .credentials(accessKeyId, secretAccessKey)
            .build();
        return new CloudflareR2Storage(client, bucketName, tmpDirectory);
    }

    private final MinioClient client;
    private final String bucketName;
    private final Path tmpDirectory;

    public CloudflareR2Storage(MinioClient client, String bucketName, Path tmpDirectory) {
        this.client = client;
        this.bucketName = bucketName;
        this.tmpDirectory = tmpDirectory;
    }

    @Override
    public String name() {
        return "cloudflare-r2";
    }

    @Override
    public void prepare() throws Exception {
        Files.createDirectories(this.tmpDirectory);
    }

    @Override
    public void shutdown() throws Exception {
        this.client.close();
    }

    @Override
    public <T> Result<Void, UploadError> upload(String key, Encoder<T> encoder, T object) {
        Path tmpFile;
        try {
            tmpFile = Files.createTempFile(this.tmpDirectory, null, null);
        } catch (IOException e) {
            return Result.failure(new UploadError.UploadException(e));
        }

        try (OutputStream out = Files.newOutputStream(tmpFile);
             BufferedOutputStream bufferedOut = new BufferedOutputStream(out, BUFFER_SIZE);
             GZIPOutputStream gzipOut = new GZIPOutputStream(bufferedOut) {
                 {
                     this.def.setLevel(Deflater.BEST_COMPRESSION);
                 }
             }
        ) {
            Result<Void, UploadError.EncodeError> result = GsonIO.DEFAULT.encodeTo(gzipOut, encoder, object).mapError(UploadError.EncodeError::new);
            if (result.isFailure()) {
                try {
                    Files.deleteIfExists(tmpFile);
                } catch (IOException ignored) {
                }
                return Result.failure(result.unwrapError());
            }
        } catch (IOException e) {
            try {
                Files.deleteIfExists(tmpFile);
            } catch (IOException e1) {
                e.addSuppressed(e1);
            }
            return Result.failure(new UploadError.UploadException(e));
        }

        try {
            this.client.uploadObject(
                UploadObjectArgs.builder()
                    .contentType("application/gzip")
                    .filename(tmpFile.toAbsolutePath().toString())
                    .bucket(this.bucketName)
                    .object(key)
                    .build()
            );
        } catch (Exception e) {
            return Result.failure(new UploadError.UploadException(e));
        }

        try {
            Files.deleteIfExists(tmpFile);
        } catch (IOException e) {
            return Result.failure(new UploadError.UploadException(e));
        }

        return Result.success();
    }
}
