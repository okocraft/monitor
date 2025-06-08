package net.okocraft.monitor.core.database.operator;

import net.okocraft.monitor.core.database.mysql.operator.MySQLUploadedObjectTableOperator;
import org.jetbrains.annotations.NotNullByDefault;

@NotNullByDefault
public record Operators(
    ServersTableOperator servers,
    PlayersTableOperator players,
    WorldsTableOperator worlds,
    PlayerNameHistoryTableOperator playerNameHistory,
    LogsTableOperator logs,
    MySQLUploadedObjectTableOperator uploadedObjects
) {
}
