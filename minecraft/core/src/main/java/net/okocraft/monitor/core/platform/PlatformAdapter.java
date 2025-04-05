package net.okocraft.monitor.core.platform;

import net.okocraft.monitor.core.handler.Handlers;
import org.jetbrains.annotations.NotNullByDefault;

import java.util.concurrent.TimeUnit;

@NotNullByDefault
public interface PlatformAdapter {

    void registerEventListeners(Handlers handlers);

    void unregisterEventListeners();

    CancellableTask scheduleTask(Runnable task, long delay, long interval, TimeUnit unit);

}
