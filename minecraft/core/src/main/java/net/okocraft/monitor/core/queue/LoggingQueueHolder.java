package net.okocraft.monitor.core.queue;

import net.okocraft.monitor.core.logger.MonitorLogger;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.locks.ReentrantLock;
import java.util.function.Consumer;

public final class LoggingQueueHolder {

    public static final Consumer<Exception> DEFAULT_EXCEPTION_HANDLER = e -> MonitorLogger.logger().error("Failed to handle logs", e);

    private final List<RegisteredLoggingQueue<?>> queues = new ArrayList<>();
    private final ReentrantLock lock = new ReentrantLock();

    public <T> LoggingQueue<T> createQueue(LogHandler<T> handler, int handleLimit) {
        LoggingQueue<T> queue = new LoggingQueue<>();
        this.queues.add(new RegisteredLoggingQueue<>(queue, handler, handleLimit));
        return queue;
    }

    public void handleLimited() {
        this.lock.lock();
        try {
            for (RegisteredLoggingQueue<?> queue : this.queues) {
                try {
                    queue.handleLimited();
                } catch (Exception e) {
                    DEFAULT_EXCEPTION_HANDLER.accept(e);
                }
            }
        } finally {
            this.lock.unlock();
        }
    }

    public void handleAll() {
        this.lock.lock();
        try {
            for (RegisteredLoggingQueue<?> queue : this.queues) {
                try {
                    queue.handleAll();
                } catch (Exception e) {
                    DEFAULT_EXCEPTION_HANDLER.accept(e);
                }
            }
        } finally {
            this.lock.unlock();
        }
    }

    public interface LogHandler<T> {

        void handleLogs(List<T> list) throws Exception;

    }

    private record RegisteredLoggingQueue<T>(LoggingQueue<T> queue, LogHandler<T> handler, int handleLimit) {

        private void handleLimited() throws Exception {
            List<T> list = new ArrayList<>(this.handleLimit);
            for (int i = 0; i < this.handleLimit; i++) {
                T t = this.queue.poll();
                if (t == null) {
                    break;
                }
                list.add(t);
            }
            if (!list.isEmpty()) {
                this.handler.handleLogs(list);
            }
        }

        private void handleAll() throws Exception {
            boolean end = false;
            while (!end) {
                List<T> list = new ArrayList<>(this.handleLimit);
                for (int i = 0; i < this.handleLimit; i++) {
                    T t = this.queue.poll();
                    if (t == null) {
                        end = true;
                        break;
                    }
                    list.add(t);
                }
                if (!list.isEmpty()) {
                    this.handler.handleLogs(list);
                }
            }
        }
    }
}
