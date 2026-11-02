package net.okocraft.monitor.platform.neoforge;

import net.kyori.adventure.text.Component;
import net.kyori.adventure.text.format.NamedTextColor;
import net.minecraft.ChatFormatting;
import net.minecraft.core.RegistryAccess;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class NativeTextTest {
    @Test
    void itemAndSignTextRetainColorAcrossMinecraftComponentBoundary() {
        var nativeText = net.minecraft.network.chat.Component.literal("Star").withStyle(ChatFormatting.GOLD);
        assertEquals(Component.text("Star", NamedTextColor.GOLD),
            NativeText.toAdventure(nativeText, RegistryAccess.EMPTY));
        var back = NativeText.toMinecraft(Component.text("Star", NamedTextColor.GOLD), RegistryAccess.EMPTY);
        assertEquals("Star", back.getString());
        assertEquals(net.minecraft.network.chat.TextColor.fromLegacyFormat(ChatFormatting.GOLD), back.getStyle().getColor());
    }
}
