package net.okocraft.monitor.platform.neoforge;

import net.okocraft.monitor.core.config.notification.OreNotification;
import org.junit.jupiter.api.Test;

import java.util.Map;
import java.util.Set;

import static org.junit.jupiter.api.Assertions.assertEquals;

class OreMessageTest {
    @Test
    void notificationIncludesConfiguredDisplayNameAndLocation() {
        var setting = new OreNotification(0, Set.of("minecraft:diamond_ore"), 500,
            "%player_name% found %vein_count% %block_type% at %block_world% %block_location%",
            Map.of("minecraft:diamond_ore", "Diamond"));
        assertEquals("Alex found 4 Diamond at minecraft:overworld 10 64 -3",
            OreMessage.format(setting, "Alex", "minecraft:diamond_ore", "minecraft:overworld", 10, 64, -3, 4));
    }
}
