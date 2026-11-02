package net.okocraft.monitor.platform.neoforge;

import com.mojang.authlib.GameProfile;
import net.minecraft.core.BlockPos;
import net.minecraft.server.level.ServerChunkCache;
import net.minecraft.server.level.ServerLevel;
import net.minecraft.server.level.ServerPlayer;
import net.minecraft.world.level.GameType;
import net.minecraft.world.level.block.Blocks;
import net.neoforged.neoforge.event.level.block.BreakBlockEvent;
import net.okocraft.monitor.core.config.notification.OreNotification;
import net.okocraft.monitor.core.webhook.discord.DiscordWebhook;
import org.junit.jupiter.api.Test;

import java.util.Map;
import java.util.Set;
import java.util.UUID;

import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.verifyNoMoreInteractions;
import static org.mockito.Mockito.when;

class OreEventsTest {
    @Test
    void nativeOreBreakSendsOneConfiguredNotificationAndCanceledBreakSendsNone() {
        var player = mock(ServerPlayer.class);
        when(player.getGameProfile()).thenReturn(new GameProfile(UUID.randomUUID(), "Alex"));
        when(player.gameMode()).thenReturn(GameType.SURVIVAL);
        var level = mock(ServerLevel.class);
        when(level.dimension()).thenReturn(net.minecraft.world.level.Level.OVERWORLD);
        when(level.getChunkSource()).thenReturn(mock(ServerChunkCache.class));
        var setting = new OreNotification(0, Set.of("minecraft:diamond_ore"), 1,
            "%player_name% found %vein_count% %block_type% at %block_world% %block_location%",
            Map.of("minecraft:diamond_ore", "Diamond"));
        var webhook = mock(DiscordWebhook.class);
        var events = new OreEvents(setting, webhook);

        events.onBreak(new BreakBlockEvent(level, new BlockPos(10, 64, -3), Blocks.DIAMOND_ORE.defaultBlockState(), player));
        var canceled = new BreakBlockEvent(level, new BlockPos(20, 64, -3), Blocks.DIAMOND_ORE.defaultBlockState(), player);
        canceled.setCanceled(true);
        events.onBreak(canceled);

        verify(webhook).send("Alex found 1 Diamond at minecraft:overworld 10 64 -3");
        verifyNoMoreInteractions(webhook);
    }
}
