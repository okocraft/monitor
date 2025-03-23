package net.okocraft.monitor.core.platform;

import net.okocraft.monitor.core.handler.Handlers;

public interface PlatformAdapter {

    void registerEventListeners(Handlers handlers);

    void unregisterEventListeners();

}
