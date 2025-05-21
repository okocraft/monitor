package net.okocraft.monitor.core.models.logs;

import net.kyori.adventure.key.Key;
import net.kyori.adventure.text.Component;
import net.okocraft.monitor.core.models.BlockPosition;

import java.time.LocalDateTime;
import java.util.List;

public record PlayerEditSignLog(int playerId, int worldId, BlockPosition position, Key blockType, Side side,
                                List<Component> lines, LocalDateTime time) {
    public enum Side {
        FRONT,
        BACK;

        public int id() {
            return this.ordinal() + 1;
        }
    }
}
