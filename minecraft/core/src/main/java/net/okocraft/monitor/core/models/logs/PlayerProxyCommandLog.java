package net.okocraft.monitor.core.models.logs;

import java.time.LocalDateTime;

public record PlayerProxyCommandLog(int playerId, int serverId, String command, LocalDateTime time) {
}
