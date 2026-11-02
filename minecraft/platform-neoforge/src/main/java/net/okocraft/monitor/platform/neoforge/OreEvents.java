package net.okocraft.monitor.platform.neoforge;

import net.minecraft.core.BlockPos;
import net.minecraft.core.registries.BuiltInRegistries;
import net.minecraft.server.level.ServerLevel;
import net.minecraft.server.level.ServerPlayer;
import net.minecraft.world.level.GameType;
import net.minecraft.world.level.block.Block;
import net.neoforged.bus.api.SubscribeEvent;
import net.neoforged.bus.api.EventPriority;
import net.neoforged.neoforge.event.level.BlockEvent;
import net.neoforged.neoforge.event.level.block.BreakBlockEvent;
import net.okocraft.monitor.core.config.notification.OreNotification;
import net.okocraft.monitor.core.webhook.discord.DiscordWebhook;

import java.util.ArrayDeque;
import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

final class OreEvents {
    private final OreNotification setting;
    private final DiscordWebhook webhook;
    private final Map<ServerLevel, Set<BlockPos>> placed = new HashMap<>();
    private final Map<ServerLevel, Set<BlockPos>> seen = new HashMap<>();

    OreEvents(OreNotification setting, DiscordWebhook webhook) {
        this.setting = setting;
        this.webhook = webhook;
    }

    @SubscribeEvent(priority = EventPriority.LOWEST)
    public void onPlace(BlockEvent.EntityPlaceEvent event) {
        if (event.isCanceled() || !(event.getLevel() instanceof ServerLevel level)) {
            return;
        }
        String blockKey = BuiltInRegistries.BLOCK.getKey(event.getPlacedBlock().getBlock()).toString();
        if (setting.enabledOres().contains(blockKey)) {
            placed.computeIfAbsent(level, ignored -> new HashSet<>()).add(event.getPos().immutable());
        }
    }

    @SubscribeEvent(priority = EventPriority.LOWEST)
    public void onBreak(BreakBlockEvent event) {
        if (event.isCanceled() || !(event.getLevel() instanceof ServerLevel level) || !(event.getPlayer() instanceof ServerPlayer player)) {
            return;
        }
        if (player.gameMode() == GameType.CREATIVE || player.gameMode() == GameType.SPECTATOR) {
            return;
        }
        Block block = event.getState().getBlock();
        String blockKey = BuiltInRegistries.BLOCK.getKey(block).toString();
        if (!setting.enabledOres().contains(blockKey)) {
            return;
        }
        BlockPos center = event.getPos();
        if (placed.getOrDefault(level, Set.of()).contains(center)) {
            return;
        }
        Set<BlockPos> seenInLevel = seen.computeIfAbsent(level, ignored -> new HashSet<>());
        if (!seenInLevel.add(center.immutable())) {
            return;
        }
        int count = countVein(level, center, block, seenInLevel);
        webhook.send(OreMessage.format(setting, player.getGameProfile().name(), blockKey, DimensionIds.name(level),
            center.getX(), center.getY(), center.getZ(), count));
    }

    private int countVein(ServerLevel level, BlockPos center, Block block, Set<BlockPos> seenInLevel) {
        ArrayDeque<BlockPos> queue = new ArrayDeque<>();
        queue.add(center.immutable());
        int count = 0;
        int searched = 0;
        while (!queue.isEmpty() && searched < setting.maxSearchCount()) {
            BlockPos current = queue.removeFirst();
            count++;
            for (int x = -1; x <= 1 && searched < setting.maxSearchCount(); x++) {
                for (int y = -1; y <= 1 && searched < setting.maxSearchCount(); y++) {
                    for (int z = -1; z <= 1 && searched < setting.maxSearchCount(); z++) {
                        if (x == 0 && y == 0 && z == 0) {
                            continue;
                        }
                        BlockPos next = current.offset(x, y, z);
                        if (!level.getChunkSource().hasChunk(next.getX() >> 4, next.getZ() >> 4)) {
                            continue;
                        }
                        searched++;
                        if (!seenInLevel.contains(next) && !placed.getOrDefault(level, Set.of()).contains(next) && level.getBlockState(next).is(block)) {
                            BlockPos stored = next.immutable();
                            seenInLevel.add(stored);
                            queue.add(stored);
                        }
                    }
                }
            }
        }
        return count;
    }
}
