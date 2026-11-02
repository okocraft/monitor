package net.okocraft.monitor.platform.neoforge;

import net.minecraft.server.level.ServerLevel;
import net.minecraft.world.level.Level;
import net.okocraft.monitor.core.models.BlockPosition;

import java.util.UUID;

final class DimensionIds {
    static String name(ServerLevel level) {
        return level.dimension().identifier().toString();
    }

    static UUID uuid(Level level) {
        return WorldIdentity.uuid(level.dimension().identifier().toString());
    }

    static BlockPosition position(net.minecraft.core.BlockPos pos) {
        return new BlockPosition(pos.getX(), pos.getY(), pos.getZ());
    }

    private DimensionIds() {
    }
}
