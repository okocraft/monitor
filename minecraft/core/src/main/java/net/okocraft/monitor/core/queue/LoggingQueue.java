package net.okocraft.monitor.core.queue;

import ca.spottedleaf.concurrentutil.collection.MultiThreadedQueue;
import org.jetbrains.annotations.NotNullByDefault;
import org.jetbrains.annotations.Nullable;

import java.util.Objects;
import java.util.Queue;

@NotNullByDefault
public final class LoggingQueue<T> {

    private final Queue<T> queue;

    public LoggingQueue() {
        this.queue = new MultiThreadedQueue<>();
    }

    public void push(T item) {
        Objects.requireNonNull(item, "item cannot be null");
        this.queue.add(item);
    }

    public @Nullable T poll() {
        return this.queue.poll();
    }
}
