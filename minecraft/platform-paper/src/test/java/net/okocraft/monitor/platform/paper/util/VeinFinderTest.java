package net.okocraft.monitor.platform.paper.util;

import org.jetbrains.annotations.NotNull;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.stream.IntStream;
import java.util.stream.Stream;

import static net.okocraft.monitor.platform.paper.util.VeinFinder.X_BIT_SIZE;
import static net.okocraft.monitor.platform.paper.util.VeinFinder.Y_BIT_SIZE;
import static net.okocraft.monitor.platform.paper.util.VeinFinder.Z_BIT_SIZE;
import static net.okocraft.monitor.platform.paper.util.VeinFinder.canPack;
import static net.okocraft.monitor.platform.paper.util.VeinFinder.packLocation;
import static net.okocraft.monitor.platform.paper.util.VeinFinder.unpackX;
import static net.okocraft.monitor.platform.paper.util.VeinFinder.unpackY;
import static net.okocraft.monitor.platform.paper.util.VeinFinder.unpackZ;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertTrue;

class VeinFinderTest {

    @ParameterizedTest
    @MethodSource("locations")
    void testPacking(@NotNull BlockPos pos) {
        long packed = packLocation(pos.x(), pos.y(), pos.z());
        assertEquals(pos, new BlockPos(unpackX(packed), unpackY(packed), unpackZ(packed)));
    }

    private static @NotNull Stream<BlockPos> locations() {
        return Stream.of(
            new BlockPos(0, 0, 0),
            new BlockPos(10, 10, 10),
            new BlockPos(-10, -10, -10)
        );
    }

    @ParameterizedTest
    @MethodSource("sizes")
    void testCanPack(int size) {
        int max = (1 << size) - 1;
        int min = ~(1 << size);
        assertTrue(canPack(max, size));
        assertFalse(canPack(max + 1, size));
        assertTrue(canPack(min, size));
        assertFalse(canPack(min - 1, size));
    }

    private static IntStream sizes() {
        return IntStream.of(X_BIT_SIZE, Y_BIT_SIZE, Z_BIT_SIZE);
    }

    private record BlockPos(int x, int y, int z) {
    }
}
