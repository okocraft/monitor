package net.okocraft.monitor.core.cloud.storage;

import net.okocraft.monitor.core.cloud.storage.cloudflare.CloudflareR2Storage;
import net.okocraft.monitor.core.cloud.storage.local.LocalStorage;
import net.okocraft.monitor.core.config.UploadConfig;

import java.nio.file.Path;

public final class CloudStorageFactory {

    public static CloudStorage create(Path directory, UploadConfig.CloudConfig config) {
        if (config.type().equals("r2")) {
            UploadConfig.CloudConfig.R2Config r2Config = config.r2();
            return CloudflareR2Storage.create(r2Config.accountId(), r2Config.accessKey(), r2Config.secretKey(), r2Config.bucketName(), directory.resolve(".cloud.tmp"));
        }
        return new LocalStorage(directory.resolve("local"));
    }

    private CloudStorageFactory() {
        throw new UnsupportedOperationException();
    }
}
