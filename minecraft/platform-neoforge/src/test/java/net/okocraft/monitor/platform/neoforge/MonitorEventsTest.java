package net.okocraft.monitor.platform.neoforge;

import com.mojang.authlib.GameProfile;
import com.mojang.brigadier.ParseResults;
import com.mojang.brigadier.StringReader;
import com.mojang.brigadier.context.CommandContextBuilder;
import net.kyori.adventure.text.Component;
import net.minecraft.core.BlockPos;
import net.minecraft.core.RegistryAccess;
import net.minecraft.commands.CommandSourceStack;
import net.minecraft.server.level.ServerPlayer;
import net.minecraft.server.level.ServerLevel;
import net.minecraft.server.network.ServerGamePacketListenerImpl;
import net.minecraft.network.Connection;
import net.minecraft.network.DisconnectionDetails;
import net.minecraft.world.level.Level;
import net.minecraft.world.item.ItemStack;
import net.minecraft.world.item.Items;
import net.minecraft.world.inventory.AnvilMenu;
import net.minecraft.world.inventory.ContainerLevelAccess;
import net.okocraft.monitor.platform.neoforge.mixin.ItemCombinerMenuAccess;
import net.neoforged.neoforge.event.ServerChatEvent;
import net.neoforged.neoforge.event.CommandEvent;
import net.neoforged.neoforge.event.entity.player.AnvilCraftEvent;
import net.neoforged.neoforge.common.NeoForge;
import net.neoforged.neoforge.event.level.LevelEvent;
import net.neoforged.neoforge.event.entity.player.PlayerEvent;
import net.okocraft.monitor.core.handler.PlayerHandler;
import net.okocraft.monitor.core.handler.WorldHandler;
import net.okocraft.monitor.core.models.BlockPosition;
import net.okocraft.monitor.core.models.logs.PlayerConnectLog;
import org.junit.jupiter.api.Test;

import java.net.InetSocketAddress;
import java.util.UUID;

import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

class MonitorEventsTest {
    @Test
    void nativeJoinAndLeaveCarryPlayerIdentityAndRemoteAddressToCore() {
        var playerId = UUID.fromString("cc376fb4-4c73-4cba-889c-4f6da8eaa0c9");
        var address = new InetSocketAddress("127.0.0.1", 25565);
        var player = mock(ServerPlayer.class);
        player.connection = mock(ServerGamePacketListenerImpl.class);
        when(player.connection.getConnection()).thenReturn(mock(Connection.class));
        when(player.getUUID()).thenReturn(playerId);
        when(player.getGameProfile()).thenReturn(new GameProfile(playerId, "Alex"));
        when(player.connection.getRemoteAddress()).thenReturn(address);
        var handler = mock(PlayerHandler.class);
        var events = new MonitorEvents(handler, mock(WorldHandler.class));

        events.onLogin(new PlayerEvent.PlayerLoggedInEvent(player));
        events.onLogout(new PlayerEvent.PlayerLoggedOutEvent(player));

        verify(handler).onJoin(playerId, "Alex", address);
        verify(handler).onLeave(playerId, PlayerConnectLog.Action.DISCONNECT, address, Component.empty());
    }

    @Test
    void nativeChatAuditsRawTextWorldAndPositionEvenWhenAnotherModCancelsDelivery() {
        var playerId = UUID.fromString("cc376fb4-4c73-4cba-889c-4f6da8eaa0c9");
        var player = mock(ServerPlayer.class);
        var level = mock(ServerLevel.class);
        when(player.getUUID()).thenReturn(playerId);
        when(player.getGameProfile()).thenReturn(new GameProfile(playerId, "Alex"));
        when(player.level()).thenReturn(level);
        when(player.blockPosition()).thenReturn(new BlockPos(10, 64, -3));
        when(level.dimension()).thenReturn(Level.OVERWORLD);
        var handler = mock(PlayerHandler.class);
        var events = new MonitorEvents(handler, mock(WorldHandler.class));

        events.onChat(new ServerChatEvent(player, "hello", net.minecraft.network.chat.Component.literal("rendered")));
        var canceled = new ServerChatEvent(player, "secret", net.minecraft.network.chat.Component.literal("secret"));
        canceled.setCanceled(true);
        NeoForge.EVENT_BUS.register(events);
        try {
            NeoForge.EVENT_BUS.post(canceled);
        } finally {
            NeoForge.EVENT_BUS.unregister(events);
        }

        verify(handler).onChat(playerId, UUID.fromString("462cf3b4-c73a-31fa-8e13-89e00f110626"),
            new BlockPosition(10, 64, -3), Component.text("hello"));
        verify(handler).onChat(playerId, UUID.fromString("462cf3b4-c73a-31fa-8e13-89e00f110626"),
            new BlockPosition(10, 64, -3), Component.text("secret"));
        org.mockito.Mockito.verifyNoMoreInteractions(handler);
    }

    @Test
    void nativeWorldLoadUsesStableDimensionIdentity() {
        var level = mock(ServerLevel.class);
        when(level.dimension()).thenReturn(Level.NETHER);
        var worlds = mock(WorldHandler.class);

        new MonitorEvents(mock(PlayerHandler.class), worlds).onLevelLoad(new LevelEvent.Load(level));

        verify(worlds).onLoad(UUID.fromString("e4d86236-4a26-34c7-8585-2c59482d56a0"), "minecraft:the_nether");
    }

    @Test
    void nativePlayerCommandAuditsSlashAndLocationEvenWhenAnotherModCancelsExecution() {
        var playerId = UUID.fromString("cc376fb4-4c73-4cba-889c-4f6da8eaa0c9");
        var player = mock(ServerPlayer.class);
        var level = mock(ServerLevel.class);
        when(player.getUUID()).thenReturn(playerId);
        when(player.level()).thenReturn(level);
        when(player.blockPosition()).thenReturn(new BlockPos(-7, 70, 9));
        when(level.dimension()).thenReturn(Level.OVERWORLD);
        var source = mock(CommandSourceStack.class);
        when(source.getPlayer()).thenReturn(player);
        @SuppressWarnings("unchecked")
        var results = (ParseResults<CommandSourceStack>) mock(ParseResults.class);
        @SuppressWarnings("unchecked")
        var context = (CommandContextBuilder<CommandSourceStack>) mock(CommandContextBuilder.class);
        when(results.getContext()).thenReturn(context);
        when(context.getSource()).thenReturn(source);
        when(results.getReader()).thenReturn(new StringReader("say hello"));
        var handler = mock(PlayerHandler.class);
        var events = new MonitorEvents(handler, mock(WorldHandler.class));

        events.onCommand(new CommandEvent(results));
        when(results.getReader()).thenReturn(new StringReader("stop"));
        var canceled = new CommandEvent(results);
        canceled.setCanceled(true);
        NeoForge.EVENT_BUS.register(events);
        try {
            NeoForge.EVENT_BUS.post(canceled);
        } finally {
            NeoForge.EVENT_BUS.unregister(events);
        }

        verify(handler).onWorldCommand(playerId, UUID.fromString("462cf3b4-c73a-31fa-8e13-89e00f110626"),
            new BlockPosition(-7, 70, 9), "/say hello");
        verify(handler).onWorldCommand(playerId, UUID.fromString("462cf3b4-c73a-31fa-8e13-89e00f110626"),
            new BlockPosition(-7, 70, 9), "/stop");
        org.mockito.Mockito.verifyNoMoreInteractions(handler);
    }

    @Test
    void nativeAnvilPickupRecordsChangedCustomNameAndItem() {
        var playerId = UUID.fromString("cc376fb4-4c73-4cba-889c-4f6da8eaa0c9");
        var player = mock(ServerPlayer.class);
        var level = mock(ServerLevel.class);
        when(player.getUUID()).thenReturn(playerId);
        when(player.level()).thenReturn(level);
        when(player.blockPosition()).thenReturn(new BlockPos(2, 65, 8));
        when(level.dimension()).thenReturn(Level.OVERWORLD);
        when(level.registryAccess()).thenReturn(RegistryAccess.EMPTY);
        var input = mock(ItemStack.class);
        var renamed = mock(ItemStack.class);
        when(input.copy()).thenReturn(input);
        when(renamed.copy()).thenReturn(renamed);
        when(renamed.getCustomName()).thenReturn(net.minecraft.network.chat.Component.literal("Star"));
        when(renamed.getItem()).thenReturn(Items.DIAMOND_SWORD);
        when(renamed.getCount()).thenReturn(1);
        var menu = mock(AnvilMenu.class);
        when(((ItemCombinerMenuAccess) menu).monitor$access())
            .thenReturn(ContainerLevelAccess.create(level, new BlockPos(3, 64, 9)));
        var handler = mock(PlayerHandler.class);
        var events = new MonitorEvents(handler, mock(WorldHandler.class));

        events.onAnvil(new AnvilCraftEvent.Post(menu, player, input, mock(ItemStack.class), renamed));

        verify(handler).onRenameItem(playerId, UUID.fromString("462cf3b4-c73a-31fa-8e13-89e00f110626"),
            new BlockPosition(3, 64, 9), net.kyori.adventure.key.Key.key("minecraft:diamond_sword"), Component.text("Star"), 1);

        when(input.getCustomName()).thenReturn(net.minecraft.network.chat.Component.literal("Old"));
        when(renamed.getCustomName()).thenReturn(null);
        events.onAnvil(new AnvilCraftEvent.Post(menu, player, input, mock(ItemStack.class), renamed));
        verify(handler).onRenameItem(playerId, UUID.fromString("462cf3b4-c73a-31fa-8e13-89e00f110626"),
            new BlockPosition(3, 64, 9), net.kyori.adventure.key.Key.key("minecraft:diamond_sword"), Component.empty(), 1);
    }

    @Test
    void logoutPreservesKickReasonAndSeparatesTimeoutAndNetworkError() {
        var playerId = UUID.fromString("cc376fb4-4c73-4cba-889c-4f6da8eaa0c9");
        var player = mock(ServerPlayer.class);
        player.connection = mock(ServerGamePacketListenerImpl.class);
        var connection = mock(Connection.class);
        var level = mock(ServerLevel.class);
        when(player.getUUID()).thenReturn(playerId);
        when(player.connection.getConnection()).thenReturn(connection);
        when(player.connection.getRemoteAddress()).thenReturn(new InetSocketAddress("127.0.0.1", 25565));
        when(player.level()).thenReturn(level);
        when(level.registryAccess()).thenReturn(RegistryAccess.EMPTY);
        var handler = mock(PlayerHandler.class);
        var events = new MonitorEvents(handler, mock(WorldHandler.class));

        when(connection.getDisconnectionDetails())
            .thenReturn(new DisconnectionDetails(net.minecraft.network.chat.Component.literal("Banned")));
        events.onLogout(new PlayerEvent.PlayerLoggedOutEvent(player));
        when(connection.getDisconnectionDetails())
            .thenReturn(new DisconnectionDetails(net.minecraft.network.chat.Component.translatable("disconnect.timeout")));
        events.onLogout(new PlayerEvent.PlayerLoggedOutEvent(player));
        when(connection.getDisconnectionDetails())
            .thenReturn(new DisconnectionDetails(net.minecraft.network.chat.Component.translatable("multiplayer.disconnect.idling")));
        events.onLogout(new PlayerEvent.PlayerLoggedOutEvent(player));
        when(connection.getDisconnectionDetails())
            .thenReturn(new DisconnectionDetails(net.minecraft.network.chat.Component.translatable("disconnect.genericReason")));
        events.onLogout(new PlayerEvent.PlayerLoggedOutEvent(player));

        var address = new InetSocketAddress("127.0.0.1", 25565);
        verify(handler).onLeave(playerId, PlayerConnectLog.Action.KICKED, address, Component.text("Banned"));
        verify(handler).onLeave(playerId, PlayerConnectLog.Action.TIMED_OUT, address, Component.empty());
        verify(handler).onLeave(playerId, PlayerConnectLog.Action.KICKED, address, Component.translatable("multiplayer.disconnect.idling"));
        verify(handler).onLeave(playerId, PlayerConnectLog.Action.ERRONEOUS_STATE, address, Component.empty());
    }
}
