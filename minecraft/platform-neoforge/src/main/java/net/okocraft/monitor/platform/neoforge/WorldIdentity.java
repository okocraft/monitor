package net.okocraft.monitor.platform.neoforge;

import java.nio.charset.StandardCharsets;
import java.util.UUID;

final class WorldIdentity {
    static UUID uuid(String dimensionKey) {
        return UUID.nameUUIDFromBytes(("monitor:dimension:" + dimensionKey).getBytes(StandardCharsets.UTF_8));
    }

    private WorldIdentity() {
    }
}
