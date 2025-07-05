package net.okocraft.monitor.core.util;

import org.jetbrains.annotations.NotNull;

import java.util.Locale;
import java.util.concurrent.ThreadFactory;
import java.util.concurrent.atomic.AtomicInteger;

public final class NamedThreadFactory implements ThreadFactory {

    public static final NamedThreadFactory DEFAULT = createDefault();

    private static @NotNull NamedThreadFactory createDefault() {
        return new NamedThreadFactory("Monitor-Thread-%d");
    }

    private final String nameFormat;
    private final AtomicInteger counter = new AtomicInteger();

    private NamedThreadFactory(@NotNull String nameFormat) {
        this.nameFormat = nameFormat;
        this.nextName(); // check if the format is valid
    }

    @Override
    public @NotNull Thread newThread(@NotNull Runnable r) {
        return new Thread(r, this.nextName());
    }

    private @NotNull String nextName() {
        return format(this.nameFormat, this.counter.getAndIncrement());
    }

    private static @NotNull String format(String format, Object... args) {
        return String.format(Locale.ROOT, format, args);
    }
}
