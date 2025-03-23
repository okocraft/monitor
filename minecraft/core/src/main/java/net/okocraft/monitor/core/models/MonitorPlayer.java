package net.okocraft.monitor.core.models;

import org.jetbrains.annotations.NotNullByDefault;

import java.util.UUID;

@NotNullByDefault
public record MonitorPlayer(int playerId, UUID uuid, String name) {
}
