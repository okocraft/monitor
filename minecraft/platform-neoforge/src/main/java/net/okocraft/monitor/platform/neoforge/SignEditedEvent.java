package net.okocraft.monitor.platform.neoforge;

import net.minecraft.world.entity.player.Player;
import net.minecraft.world.level.block.entity.SignBlockEntity;
import net.neoforged.bus.api.Event;

public final class SignEditedEvent extends Event {
    private final Player player;
    private final SignBlockEntity sign;
    private final boolean front;

    public SignEditedEvent(Player player, SignBlockEntity sign, boolean front) {
        this.player = player;
        this.sign = sign;
        this.front = front;
    }

    public Player player() {
        return player;
    }

    public SignBlockEntity sign() {
        return sign;
    }

    public boolean front() {
        return front;
    }
}
