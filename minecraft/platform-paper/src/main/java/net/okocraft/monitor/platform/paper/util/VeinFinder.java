package net.okocraft.monitor.platform.paper.util;

import it.unimi.dsi.fastutil.longs.LongOpenHashSet;
import it.unimi.dsi.fastutil.longs.LongSet;
import net.okocraft.monitor.core.logger.MonitorLogger;
import org.bukkit.Bukkit;
import org.bukkit.Location;
import org.bukkit.Material;
import org.bukkit.World;
import org.bukkit.block.Block;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.VisibleForTesting;

import java.util.concurrent.locks.StampedLock;

public class VeinFinder {

    private final LongSet placedLocation = new LongOpenHashSet();
    private final LongSet placedOrFoundLocations = new LongOpenHashSet();
    private final StampedLock lock = new StampedLock();

    private final int maxSearchCount;

    public VeinFinder(int maxSearchCount) {
        this.maxSearchCount = maxSearchCount;
    }

    public boolean isPlacedLocation(@NotNull Location location) {
        return this.checkContains(packLocation(location), this.placedLocation);
    }

    public void recordPlaced(@NotNull Location location) {
        long packedLocation = packLocation(location);
        long stamp = this.lock.writeLock();

        try {
            this.placedLocation.add(packedLocation);
            this.placedOrFoundLocations.add(packedLocation);
        } finally {
            this.lock.unlockWrite(stamp);
        }
    }

    public @NotNull SearchResult findSameBlock(@NotNull Block center) {
        var result = new SearchResult();
        long packedCenterLocation = packLocation(center.getLocation());

        if (this.checkContains(packedCenterLocation, this.placedOrFoundLocations)) {
            return result;
        }

        var foundLocations = new LongOpenHashSet();
        foundLocations.add(packedCenterLocation);

        this.findSameBlockType(new LongOpenHashSet(), foundLocations, center.getWorld(), center.getX(), center.getY(), center.getZ(), center.getType());

        long stamp = this.lock.writeLock();

        try {
            this.placedOrFoundLocations.add(packedCenterLocation);
            this.placedOrFoundLocations.addAll(foundLocations);
        } finally {
            this.lock.unlockWrite(stamp);
        }

        foundLocations.forEach(packedLocation -> result.increaseCount());
        return result;
    }

    private void findSameBlockType(@NotNull LongSet searchedLocations, @NotNull LongSet foundLocations,
                                   @NotNull World world, int centerX, int centerY, int centerZ, @NotNull Material type) {
        for (int x = -1; x < 2; x++) {
            for (int y = -1; y < 2; y++) {
                for (int z = -1; z < 2; z++) {
                    if (x == 0 && y == 0 && z == 0) {
                        continue;
                    }

                    if (this.maxSearchCount <= searchedLocations.size()) {
                        break;
                    }

                    int blockX = centerX + x;
                    int blockY = centerY + y;
                    int blockZ = centerZ + z;

                    if (!canPack(blockX, blockY, blockZ)) {
                        MonitorLogger.logger().error("Cannot handle the location ({}, {}, {})", blockX, blockY, blockZ);
                        break;
                    }

                    long packedLocation = packLocation(blockX, blockY, blockZ);

                    if (!searchedLocations.add(packedLocation) ||
                        this.checkContains(packedLocation, this.placedOrFoundLocations)) {
                        // already handled or cannot access due to outside of Folia's regions
                        continue;
                    }

                    if (Bukkit.isOwnedByCurrentRegion(world, blockX >> 4, blockZ >> 4) && type == world.getType(blockX, blockY, blockZ)) {
                        foundLocations.add(packedLocation);
                        this.findSameBlockType(searchedLocations, foundLocations, world, blockX, blockY, blockZ, type);
                    }
                }
            }
        }
    }

    private boolean checkContains(long packedLocation, @NotNull LongSet set) {
        {
            long stamp = this.lock.tryOptimisticRead();
            boolean result = set.contains(packedLocation);

            if (this.lock.validate(stamp)) {
                return result;
            }
        }

        long stamp = this.lock.readLock();

        try {
            return set.contains(packedLocation);
        } finally {
            this.lock.unlockRead(stamp);
        }
    }

    public static class SearchResult {

        private int count;

        public int count() {
            return this.count;
        }

        private void increaseCount() {
            this.count++;
        }
    }

    /* Utility for packing block locations */

    @VisibleForTesting
    static final int X_BIT_SIZE = 27;
    @VisibleForTesting
    static final int Y_BIT_SIZE = 10;
    @VisibleForTesting
    static final int Z_BIT_SIZE = 27;

    private static final int X_SHIFT = Y_BIT_SIZE + Z_BIT_SIZE;
    private static final int Y_LEFT_SHIFT = Z_BIT_SIZE;
    private static final int Y_RIGHT_SHIFT = X_BIT_SIZE + Z_BIT_SIZE;
    private static final int Z_SHIFT = X_BIT_SIZE + Z_BIT_SIZE;

    private static final int Y_MASK = (1 << Y_BIT_SIZE) - 1;
    private static final int Z_MASK = (1 << Z_BIT_SIZE) - 1;

    private static long packLocation(@NotNull Location location) {
        return packLocation(location.getBlockX(), location.getBlockY(), location.getBlockZ());
    }

    private static boolean canPack(int x, int y, int z) {
        return canPack(x, X_BIT_SIZE) && canPack(y, Y_BIT_SIZE) && canPack(z, Z_BIT_SIZE);
    }

    @VisibleForTesting
    static long packLocation(int x, int y, int z) {
        long value = 0;
        value |= (long) x << X_SHIFT;
        value |= (long) (y & Y_MASK) << Y_LEFT_SHIFT;
        value |= (long) z & Z_MASK;
        return value;
    }

    @VisibleForTesting
    static int unpackX(long value) {
        return (int) (value >> X_SHIFT);
    }

    @VisibleForTesting
    static int unpackY(long value) {
        return (int) ((value << Y_LEFT_SHIFT) >> Y_RIGHT_SHIFT);
    }

    @VisibleForTesting
    static int unpackZ(long value) {
        return (int) ((value << Z_SHIFT) >> Z_SHIFT);
    }

    @VisibleForTesting
    static boolean canPack(int value, int size) {
        return ~(1 << size) <= value && value <= (1 << size) - 1;
    }
}
