package net.okocraft.monitor.core.models.logs;

import org.jetbrains.annotations.NotNullByDefault;

import java.time.LocalDateTime;

@NotNullByDefault
public record PlayerConnectLog(int playerId, int serverId, Action action, String address, String reason, LocalDateTime time) {
    public enum Action {
        UNKNOWN,
        CONNECT,
        DISCONNECT,
        TIMED_OUT,
        ERRONEOUS_STATE,
        KICKED,;

        public static Action byId(int id) {
            Action[] values = Action.values();
            if (id < 0 || id >= values.length) {
                return UNKNOWN;
            }
            return values[id];
        }

        public int id() {
            return this.ordinal();
        }
    }
}
