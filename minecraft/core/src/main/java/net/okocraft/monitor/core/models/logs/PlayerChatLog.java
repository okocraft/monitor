package net.okocraft.monitor.core.models.logs;

import net.okocraft.monitor.core.models.BlockPosition;

import java.time.LocalDateTime;

public record PlayerChatLog(int playerId, int worldId, BlockPosition position, String message, LocalDateTime time) {
}
