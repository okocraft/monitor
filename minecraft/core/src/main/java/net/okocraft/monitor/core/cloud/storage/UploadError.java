package net.okocraft.monitor.core.cloud.storage;

public sealed interface UploadError permits UploadError.EncodeError, UploadError.UploadException {

    record EncodeError(dev.siroshun.codec4j.api.error.EncodeError error) implements UploadError {
    }

    record UploadException(Exception exception) implements UploadError {
    }
}
