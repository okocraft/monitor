package net.okocraft.monitor.platform.neoforge;

import com.mojang.serialization.JsonOps;
import net.kyori.adventure.text.serializer.gson.GsonComponentSerializer;
import net.minecraft.core.HolderLookup;
import net.minecraft.network.chat.ComponentSerialization;
import net.minecraft.resources.RegistryOps;

final class NativeText {
    static net.kyori.adventure.text.Component toAdventure(net.minecraft.network.chat.Component text,
                                                          HolderLookup.Provider registries) {
        var encoded = ComponentSerialization.CODEC
            .encodeStart(RegistryOps.create(JsonOps.INSTANCE, registries), text)
            .getOrThrow();
        return GsonComponentSerializer.gson().deserializeFromTree(encoded);
    }

    static net.minecraft.network.chat.Component toMinecraft(net.kyori.adventure.text.Component text,
                                                            HolderLookup.Provider registries) {
        var encoded = GsonComponentSerializer.gson().serializeToTree(text);
        return ComponentSerialization.CODEC
            .parse(RegistryOps.create(JsonOps.INSTANCE, registries), encoded)
            .getOrThrow();
    }

    private NativeText() {
    }
}
