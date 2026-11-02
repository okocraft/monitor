package net.okocraft.monitor.platform.neoforge;

import net.kyori.adventure.key.Key;
import net.kyori.adventure.text.Component;
import net.minecraft.core.BlockPos;
import net.minecraft.core.RegistryAccess;
import net.minecraft.server.level.ServerLevel;
import net.minecraft.server.level.ServerPlayer;
import net.minecraft.server.network.FilteredText;
import net.minecraft.world.level.Level;
import net.minecraft.world.level.block.Blocks;
import net.minecraft.world.level.block.entity.SignBlockEntity;
import net.neoforged.neoforge.common.NeoForge;
import net.okocraft.monitor.core.handler.PlayerHandler;
import net.okocraft.monitor.core.handler.WorldHandler;
import net.okocraft.monitor.core.models.BlockPosition;
import net.okocraft.monitor.core.models.logs.PlayerEditSignLog;
import org.junit.jupiter.api.Test;

import java.util.List;
import java.util.UUID;

import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.verifyNoMoreInteractions;
import static org.mockito.Mockito.when;

class SignEditNativeTest {
    @Test
    void acceptedSignUpdateThroughMinecraftMixinReachesCoreOnce() {
        var playerId = UUID.fromString("cc376fb4-4c73-4cba-889c-4f6da8eaa0c9");
        var player = mock(ServerPlayer.class);
        when(player.getUUID()).thenReturn(playerId);
        var level = mock(ServerLevel.class);
        when(level.dimension()).thenReturn(Level.OVERWORLD);
        when(level.registryAccess()).thenReturn(RegistryAccess.EMPTY);
        var pos = new BlockPos(5, 70, -2);
        var sign = new SignBlockEntity(pos, Blocks.OAK_SIGN.defaultBlockState());
        sign.setLevel(level);
        sign.setAllowedPlayerEditor(playerId);
        var handler = mock(PlayerHandler.class);
        var events = new MonitorEvents(handler, mock(WorldHandler.class));
        NeoForge.EVENT_BUS.register(events);
        try {
            var lines = List.of(FilteredText.passThrough("Welcome"), FilteredText.passThrough(""),
                FilteredText.passThrough(""), FilteredText.passThrough(""));
            sign.updateSignText(player, true, lines);
            sign.setAllowedPlayerEditor(playerId);
            sign.updateSignText(player, true, lines);

            verify(handler).onEditSign(playerId, UUID.fromString("462cf3b4-c73a-31fa-8e13-89e00f110626"),
                new BlockPosition(5, 70, -2), Key.key("minecraft:oak_sign"), PlayerEditSignLog.Side.FRONT,
                List.of(Component.text("Welcome"), Component.empty(), Component.empty(), Component.empty()));
            verifyNoMoreInteractions(handler);
        } finally {
            NeoForge.EVENT_BUS.unregister(events);
        }
    }
}
