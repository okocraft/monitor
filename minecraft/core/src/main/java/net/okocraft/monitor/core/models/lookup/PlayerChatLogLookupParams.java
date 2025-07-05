package net.okocraft.monitor.core.models.lookup;

import java.time.LocalDateTime;

public record PlayerChatLogLookupParams(LocalDateTime start, LocalDateTime end, long limit,
                                        long offset) implements CommonLookupParams {
}
