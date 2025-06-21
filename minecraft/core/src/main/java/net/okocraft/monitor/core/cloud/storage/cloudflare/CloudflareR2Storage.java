package net.okocraft.monitor.core.cloud.storage.cloudflare;

import dev.siroshun.codec4j.api.encoder.Encoder;
import dev.siroshun.codec4j.api.error.EncodeError;
import dev.siroshun.codec4j.io.gson.GsonIO;
import dev.siroshun.codec4j.io.gzip.GzipIO;
import dev.siroshun.jfun.result.Result;
import io.minio.MinioClient;
import io.minio.UploadObjectArgs;
import net.okocraft.monitor.core.cloud.storage.CloudStorage;
import net.okocraft.monitor.core.cloud.storage.UploadError;
import net.okocraft.monitor.core.logger.MonitorLogger;
import org.jetbrains.annotations.NotNullByDefault;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;

@NotNullByDefault
public class CloudflareR2Storage implements CloudStorage {

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

        Result<Void, EncodeError> result = GzipIO.bestCompression(GsonIO.DEFAULT).encodeTo(tmpFile, encoder, object);
        if (result.isFailure()) {
            try {
                Files.deleteIfExists(tmpFile);
            } catch (IOException e) {
                MonitorLogger.logger().error("Failed to delete temporary file: {}", tmpFile, e);
            }
            return result.mapError(UploadError.EncodeError::new);
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
        } finally {
            try {
                Files.deleteIfExists(tmpFile);
            } catch (IOException e) {
                MonitorLogger.logger().error("Failed to delete temporary file: {}", tmpFile, e);
            }
        }

        return Result.success();
    }
}
