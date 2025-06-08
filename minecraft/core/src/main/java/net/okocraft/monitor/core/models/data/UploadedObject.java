package net.okocraft.monitor.core.models.data;

import java.time.Instant;
import java.time.LocalDateTime;
import java.util.UUID;

public record UploadedObject(UUID id, int type, int version,
                             UUID createdByUUID, String createdByName,
                             LocalDateTime createdTime, Instant expiresAt) {
}
