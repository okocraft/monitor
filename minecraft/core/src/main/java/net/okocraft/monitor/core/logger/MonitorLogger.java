package net.okocraft.monitor.core.logger;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.slf4j.helpers.SubstituteLogger;

public final class MonitorLogger {

    private static final SubstituteLogger LOGGER = new SubstituteLogger("monitor", null, true);

    static {
        try {
            Class.forName("org.junit.jupiter.api.Assertions");
            LOGGER.setDelegate(LoggerFactory.getLogger("OkoChat"));
        } catch (ClassNotFoundException ignored) {
        }
    }

    public static Logger logger() {
        return LOGGER;
    }

    private MonitorLogger() {
        throw new UnsupportedOperationException();
    }
}
