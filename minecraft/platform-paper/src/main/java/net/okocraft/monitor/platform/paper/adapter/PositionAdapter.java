package net.okocraft.monitor.platform.paper.adapter;

import net.okocraft.monitor.core.models.BlockPosition;
import org.bukkit.Location;
import org.jetbrains.annotations.NotNullByDefault;

@NotNullByDefault
public final class PositionAdapter {

    public static BlockPosition fromLocation(Location location) {
        return new BlockPosition(location.getBlockX(), location.getBlockY(), location.getBlockZ());
    }

    private PositionAdapter() {
        throw new UnsupportedOperationException();
    }
}
