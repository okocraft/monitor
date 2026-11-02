package net.okocraft.monitor.platform.neoforge.mixin;

import net.minecraft.server.network.FilteredText;
import net.minecraft.world.entity.player.Player;
import net.minecraft.world.level.block.entity.SignBlockEntity;
import net.minecraft.world.level.block.entity.SignText;
import net.neoforged.neoforge.common.NeoForge;
import net.okocraft.monitor.platform.neoforge.SignEditedEvent;
import org.spongepowered.asm.mixin.Mixin;
import org.spongepowered.asm.mixin.injection.At;
import org.spongepowered.asm.mixin.injection.Redirect;

import java.util.Arrays;
import java.util.List;
import java.util.function.UnaryOperator;

@Mixin(SignBlockEntity.class)
public abstract class SignEditMixin {
    @Redirect(method = "updateSignText", at = @At(value = "INVOKE", target = "Lnet/minecraft/world/level/block/entity/SignBlockEntity;updateText(Ljava/util/function/UnaryOperator;Z)Z"))
    private boolean monitor$recordEdit(SignBlockEntity sign, UnaryOperator<SignText> operation, boolean front,
                                       Player player, boolean frontText, List<FilteredText> lines) {
        var before = sign.getText(front).getMessages(player.isTextFilteringEnabled());
        boolean changed = sign.updateText(operation, front);
        if (changed && !Arrays.equals(before, sign.getText(front).getMessages(player.isTextFilteringEnabled()))) {
            NeoForge.EVENT_BUS.post(new SignEditedEvent(player, sign, front));
        }
        return changed;
    }
}
