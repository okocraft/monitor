package net.okocraft.monitor.core.database.operator;

import org.jetbrains.annotations.NotNullByDefault;

@NotNullByDefault
public record Operators(
    ServersTableOperator servers,
    PlayersTableOperator players,
    PlayerNameHistoryTableOperator playerNameHistory,
    PlayerConnectLogsTableOperator playerConnectLogs
) {
}
