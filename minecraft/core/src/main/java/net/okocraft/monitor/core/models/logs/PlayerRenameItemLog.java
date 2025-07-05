package net.okocraft.monitor.core.models.logs;

import net.kyori.adventure.key.Key;
import net.kyori.adventure.text.Component;
import net.okocraft.monitor.core.models.BlockPosition;

import java.time.LocalDateTime;

public record PlayerRenameItemLog(int playerId, int worldId, BlockPosition position, Key itemType, Component itemName,
                                  int amount, LocalDateTime time) {
}
