package net.okocraft.monitor.platform.neoforge;

import net.okocraft.monitor.core.config.notification.OreNotification;

final class OreMessage {
    static String format(OreNotification setting, String player, String blockKey, String world, int x, int y, int z, int count) {
        return setting.format()
            .replace("%player_name%", player)
            .replace("%block_type%", setting.displayNameMap().getOrDefault(blockKey, blockKey))
            .replace("%block_world%", world)
            .replace("%block_location%", x + " " + y + " " + z)
            .replace("%vein_count%", Integer.toString(count));
    }

    private OreMessage() {
    }
}
