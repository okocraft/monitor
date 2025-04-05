package net.okocraft.monitor.core.models;

import org.jetbrains.annotations.NotNullByDefault;

import java.util.UUID;

@NotNullByDefault
public record MonitorWorld(int worldId, int serverId, UUID uid, String name) {
}
