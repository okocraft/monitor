package net.okocraft.monitor.platform.neoforge;

import net.minecraft.network.DisconnectionDetails;
import net.minecraft.network.chat.contents.TranslatableContents;
import net.okocraft.monitor.core.models.logs.PlayerConnectLog;

final class DisconnectAction {
    static PlayerConnectLog.Action fromDetails(DisconnectionDetails details) {
        if (details == null) {
            return PlayerConnectLog.Action.DISCONNECT;
        }
        if (details.reason().getContents() instanceof TranslatableContents translated) {
            return switch (translated.getKey()) {
                case "disconnect.timeout" -> PlayerConnectLog.Action.TIMED_OUT;
                case "disconnect.genericReason" -> PlayerConnectLog.Action.ERRONEOUS_STATE;
                case "disconnect.endOfStream", "multiplayer.disconnect.quitting",
                     "multiplayer.disconnect.server_shutdown", "multiplayer.disconnect.generic" -> PlayerConnectLog.Action.DISCONNECT;
                default -> PlayerConnectLog.Action.KICKED;
            };
        }
        return PlayerConnectLog.Action.KICKED;
    }

    private DisconnectAction() {
    }
}
