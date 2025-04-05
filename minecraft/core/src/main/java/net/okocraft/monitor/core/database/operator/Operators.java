package net.okocraft.monitor.core.database.operator;

import org.jetbrains.annotations.NotNullByDefault;

@NotNullByDefault
public record Operators(
    ServersTableOperator servers,
    PlayersTableOperator players,
    WorldsTableOperator worlds,
    PlayerNameHistoryTableOperator playerNameHistory,
    PlayerConnectLogsTableOperator playerConnectLogs
) {
}
