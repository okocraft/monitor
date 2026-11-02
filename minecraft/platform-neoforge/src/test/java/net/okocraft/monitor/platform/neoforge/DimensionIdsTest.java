package net.okocraft.monitor.platform.neoforge;

import org.junit.jupiter.api.Test;

import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertEquals;

class DimensionIdsTest {
    @Test
    void dimensionIdentityIsStableAcrossRestartAndDistinctAcrossDimensions() {
        assertEquals(UUID.fromString("462cf3b4-c73a-31fa-8e13-89e00f110626"),
            WorldIdentity.uuid("minecraft:overworld"));
        assertEquals(UUID.fromString("e4d86236-4a26-34c7-8585-2c59482d56a0"),
            WorldIdentity.uuid("minecraft:the_nether"));
    }
}
