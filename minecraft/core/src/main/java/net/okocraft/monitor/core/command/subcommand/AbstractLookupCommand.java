package net.okocraft.monitor.core.command.subcommand;

import net.okocraft.monitor.core.models.data.PlayerConnectLogData;
import net.okocraft.monitor.core.storage.PlayerLogStorage;
import org.jetbrains.annotations.NotNullByDefault;

import java.sql.SQLException;
import java.time.LocalDateTime;
import java.util.function.Consumer;

@NotNullByDefault
public abstract class AbstractLookupCommand {

    protected final PlayerLogStorage storage;

    protected AbstractLookupCommand(PlayerLogStorage storage) {
        this.storage = storage;
    }

    protected void lookupConnectLog(Consumer<PlayerConnectLogData> consumer) throws SQLException {
        this.storage.lookupConnectLogData(new PlayerConnectLogData.LookupParams(
            LocalDateTime.now().minusDays(10),
            LocalDateTime.now()
        ), consumer);
    }

}
