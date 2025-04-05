package net.okocraft.monitor.core.models.logs;

import org.jetbrains.annotations.NotNullByDefault;

@NotNullByDefault
public record PlayerConnectLog(int playerId, int serverId, Action action, String address, String reason) {
    public enum Action {
        CONNECT,
        DISCONNECT,
        TIMED_OUT,
        ERRONEOUS_STATE,
        KICKED,;

        public int id() {
            return this.ordinal() + 1;
        }
    }
}
