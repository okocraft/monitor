package net.okocraft.monitor.platform.velocity;

import com.google.inject.Inject;
import com.velocitypowered.api.command.SimpleCommand;
import com.velocitypowered.api.event.Subscribe;
import com.velocitypowered.api.event.proxy.ProxyInitializeEvent;
import com.velocitypowered.api.event.proxy.ProxyShutdownEvent;
import com.velocitypowered.api.plugin.annotation.DataDirectory;
import com.velocitypowered.api.proxy.ProxyServer;
import com.velocitypowered.api.scheduler.ScheduledTask;
import net.okocraft.monitor.core.Monitor;
import net.okocraft.monitor.core.bootstrap.MonitorBootstrap;
import net.okocraft.monitor.core.command.Command;
import net.okocraft.monitor.core.handler.Handlers;
import net.okocraft.monitor.core.logger.MonitorLogger;
import net.okocraft.monitor.core.platform.CancellableTask;
import net.okocraft.monitor.core.platform.PlatformAdapter;
import net.okocraft.monitor.platform.velocity.adapter.CommandSenderAdapter;
import net.okocraft.monitor.platform.velocity.listener.PlayerListener;
import org.slf4j.Logger;

import java.nio.file.Path;
import java.util.List;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.TimeUnit;
import java.util.stream.Stream;

public class MonitorVelocity implements PlatformAdapter {

    private final ProxyServer server;
    private final Logger logger;
    private final Path dataDirectory;

    private Monitor monitor;

    @Inject
    public MonitorVelocity(ProxyServer server, Logger logger,
                           @DataDirectory Path dataDirectory) {
        this.server = server;
        this.logger = logger;
        this.dataDirectory = dataDirectory;
    }


    @Subscribe
    public void onEnable(ProxyInitializeEvent ignored) {
        var monitor = MonitorBootstrap.load(this.logger, this.dataDirectory);
        if (monitor.isEmpty()) {
            MonitorLogger.logger().error("Failed to load Monitor");
            return;
        }
        this.monitor = monitor.get();
        this.monitor.start(this);
    }

    @Subscribe
    public void onDisable(ProxyShutdownEvent ignored) {
        if (this.monitor == null) {
            return;
        }

        this.monitor.shutdown(this);
    }

    @Override
    public String pluginVersion() {
        return this.server.getPluginManager()
            .fromInstance(this)
            .flatMap(plugin -> plugin.getDescription().getVersion())
            .orElse("unknown");
    }

    @Override
    public void registerEventListeners(Handlers handlers) {
        Stream.of(
            new PlayerListener(this.server, handlers.player())
        ).forEach(listener -> this.server.getEventManager().register(this, listener));
    }

    @Override
    public void unregisterEventListeners() {
        this.server.getEventManager().unregisterListeners(this);
    }

    @Override
    public void registerCommand(String label, Command command) {
        this.server.getCommandManager().register(
            this.server.getCommandManager().metaBuilder(label).plugin(this).build(),
            new SimpleCommand() {
                @Override
                public void execute(Invocation invocation) {
                    command.execute(CommandSenderAdapter.wrap(invocation.source()), invocation.arguments());
                }

                @Override
                public CompletableFuture<List<String>> suggestAsync(Invocation invocation) {
                    return command.tabComplete(CommandSenderAdapter.wrap(invocation.source()), invocation.arguments());
                }
            }
        );
    }

    @Override
    public CancellableTask scheduleTask(Runnable task, long delay, long interval, TimeUnit unit) {
        ScheduledTask scheduled = this.server.getScheduler().buildTask(this, ignored -> task.run()).delay(delay, unit).repeat(interval, unit).schedule();
        return scheduled::cancel;
    }
}
