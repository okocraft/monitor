package net.okocraft.monitor.platform.neoforge;

import net.minecraft.commands.CommandSourceStack;
import net.minecraft.core.RegistryAccess;
import net.minecraft.server.MinecraftServer;
import net.minecraft.server.permissions.PermissionSet;
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
}
