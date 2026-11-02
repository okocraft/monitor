package net.okocraft.monitor.platform.neoforge;

import net.minecraft.commands.CommandSourceStack;
import net.minecraft.core.RegistryAccess;
import net.minecraft.server.MinecraftServer;
import net.minecraft.server.permissions.PermissionSet;
import net.minecraft.server.level.ServerPlayer;
import net.minecraft.server.level.ServerLevel;
import net.minecraft.server.players.PlayerList;
import net.neoforged.fml.ModList;
import net.luckperms.api.LuckPerms;
import net.luckperms.api.LuckPermsProvider;
import net.luckperms.api.model.user.User;
import net.luckperms.api.cacheddata.CachedPermissionData;
import net.luckperms.api.query.QueryOptions;
import net.luckperms.api.util.Tristate;
import net.kyori.adventure.text.Component;
import net.kyori.adventure.text.format.NamedTextColor;
import net.minecraft.ChatFormatting;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;
import static org.mockito.Mockito.mockStatic;
import static org.mockito.Mockito.RETURNS_DEEP_STUBS;
import static org.mockito.ArgumentMatchers.any;

class MonitorCommandSourceTest {
    @Test
    void nonPlayerSourceNeedsAdministrativePermissionForMonitorData() {
        var commandSource = mock(CommandSourceStack.class);
        when(commandSource.permissions()).thenReturn(PermissionSet.NO_PERMISSIONS);
        assertFalse(new MonitorCommandSource(commandSource).hasPermission("monitor.command.lookup"));

        when(commandSource.permissions()).thenReturn(PermissionSet.ALL_PERMISSIONS);
        assertTrue(new MonitorCommandSource(commandSource).hasPermission("monitor.command.lookup"));
    }

    @Test
    void adventureMessageKeepsColorWhenSentToMinecraft() {
        var commandSource = mock(CommandSourceStack.class);
        var server = mock(MinecraftServer.class);
        when(commandSource.getServer()).thenReturn(server);
        when(server.registryAccess()).thenReturn(RegistryAccess.EMPTY);

        new MonitorCommandSource(commandSource).sendMessage(Component.text("Report", NamedTextColor.GOLD));

        var message = org.mockito.ArgumentCaptor.forClass(net.minecraft.network.chat.Component.class);
        verify(commandSource).sendSystemMessage(message.capture());
        assertEquals("Report", message.getValue().getString());
        assertEquals(net.minecraft.network.chat.TextColor.fromLegacyFormat(ChatFormatting.GOLD),
            message.getValue().getStyle().getColor());
    }

    @Test
    void luckPermsDecisionOverridesOperatorDefaultAndUndefinedFallsBack() {
        var source = mock(CommandSourceStack.class);
        var player = mock(ServerPlayer.class);
        var level = mock(ServerLevel.class);
        var server = mock(MinecraftServer.class);
        var players = mock(PlayerList.class);
        var profile = new com.mojang.authlib.GameProfile(java.util.UUID.randomUUID(), "MonitorProbe");
        when(source.getPlayer()).thenReturn(player);
        when(player.getGameProfile()).thenReturn(profile);
        when(player.getUUID()).thenReturn(profile.id());
        when(player.level()).thenReturn(level);
        when(level.getServer()).thenReturn(server);
        when(server.getPlayerList()).thenReturn(players);
        when(players.isOp(any())).thenReturn(true);
        var api = mock(LuckPerms.class, RETURNS_DEEP_STUBS);
        var user = mock(User.class, RETURNS_DEEP_STUBS);
        var options = mock(QueryOptions.class);
        var permissions = mock(CachedPermissionData.class);
        var modList = mock(ModList.class);
        when(modList.isLoaded("luckperms")).thenReturn(true);
        when(api.getUserManager().getUser(profile.id())).thenReturn(user);
        when(api.getContextManager().getQueryOptions(user)).thenReturn(java.util.Optional.empty());
        when(api.getContextManager().getStaticQueryOptions()).thenReturn(options);
        when(user.getCachedData().getPermissionData(options)).thenReturn(permissions);
        String permission = "monitor.command.lookup";
        try (var mods = mockStatic(ModList.class); var provider = mockStatic(LuckPermsProvider.class)) {
            mods.when(ModList::get).thenReturn(modList);
            provider.when(LuckPermsProvider::get).thenReturn(api);
            when(permissions.checkPermission(permission)).thenReturn(Tristate.UNDEFINED);
            assertTrue(new MonitorCommandSource(source).hasPermission(permission));
            when(permissions.checkPermission(permission)).thenReturn(Tristate.FALSE);
            assertFalse(new MonitorCommandSource(source).hasPermission(permission));
            when(players.isOp(any())).thenReturn(false);
            when(permissions.checkPermission(permission)).thenReturn(Tristate.TRUE);
            assertTrue(new MonitorCommandSource(source).hasPermission(permission));
        }
    }

}
