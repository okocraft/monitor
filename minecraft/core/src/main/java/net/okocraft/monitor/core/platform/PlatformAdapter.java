package net.okocraft.monitor.core.platform;

import net.okocraft.monitor.core.command.Command;
import net.okocraft.monitor.core.command.MonitorCommand;
import net.okocraft.monitor.core.handler.Handlers;
import org.jetbrains.annotations.NotNullByDefault;

import java.util.concurrent.TimeUnit;

@NotNullByDefault
public interface PlatformAdapter {

    String pluginVersion();

    void registerEventListeners(Handlers handlers);

    void unregisterEventListeners();

    void registerCommand(String label, Command command);

    CancellableTask scheduleTask(Runnable task, long delay, long interval, TimeUnit unit);

}
