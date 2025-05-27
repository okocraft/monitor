package net.okocraft.monitor.platform.paper;

import io.papermc.paper.command.brigadier.BasicCommand;
import io.papermc.paper.command.brigadier.CommandSourceStack;
import io.papermc.paper.plugin.lifecycle.event.types.LifecycleEvents;
import io.papermc.paper.threadedregions.scheduler.ScheduledTask;
import net.okocraft.monitor.core.Monitor;
import net.okocraft.monitor.core.bootstrap.MonitorBootstrap;
import net.okocraft.monitor.core.command.Command;
import net.okocraft.monitor.core.command.MonitorCommand;
import net.okocraft.monitor.core.handler.Handlers;
import net.okocraft.monitor.core.logger.MonitorLogger;
import net.okocraft.monitor.core.platform.CancellableTask;
import net.okocraft.monitor.core.platform.PlatformAdapter;
import net.okocraft.monitor.platform.paper.adapter.CommandSenderAdapter;
import net.okocraft.monitor.platform.paper.listener.PlayerListener;
import net.okocraft.monitor.platform.paper.listener.WorldListener;
import org.bukkit.event.HandlerList;
import org.bukkit.plugin.java.JavaPlugin;
import org.jetbrains.annotations.NotNullByDefault;
import org.jetbrains.annotations.Nullable;

import java.util.Collection;
import java.util.concurrent.TimeUnit;
import java.util.stream.Stream;

@NotNullByDefault
public class MonitorPaperPlugin extends JavaPlugin implements PlatformAdapter {

    private @Nullable Monitor monitor;

    @Override
    public void onLoad() {
        var monitor = MonitorBootstrap.load(this.getSLF4JLogger(), this.getDataPath());
        if (monitor.isEmpty()) {
            MonitorLogger.logger().error("Failed to load Monitor");
            return;
        }
        this.monitor = monitor.get();
    }

    @Override
    public void onEnable() {
        if (this.monitor == null) {
            return;
        }

        this.monitor.start(this);
    }

    @Override
    public void onDisable() {
        if (this.monitor == null) {
            return;
        }

        this.monitor.shutdown(this);
    }

    @Override
    public String pluginVersion() {
        return this.getPluginMeta().getVersion();
    }

    @Override
    public void registerEventListeners(Handlers handlers) {
        Stream.of(
            new PlayerListener(handlers.player()),
            new WorldListener(handlers.world())
        ).forEach(listener -> this.getServer().getPluginManager().registerEvents(listener, this));
    }

    @Override
    public void unregisterEventListeners() {
        HandlerList.unregisterAll(this);
    }

    @Override
    public void registerCommand(String label, Command command) {
        this.getLifecycleManager().registerEventHandler(LifecycleEvents.COMMANDS.newHandler(event -> {
            event.registrar().register(label, new BasicCommand() {
                @Override
                public void execute(CommandSourceStack source, String[] args) {
                    command.execute(CommandSenderAdapter.wrap(source), args);
                }

                @Override
                public Collection<String> suggest(CommandSourceStack source, String[] args) {
                    return command.tabComplete(CommandSenderAdapter.wrap(source), args).join();
                }
            });
        }));
    }

    @Override
    public CancellableTask scheduleTask(Runnable task, long delay, long interval, TimeUnit unit) {
        ScheduledTask scheduled = this.getServer().getAsyncScheduler().runAtFixedRate(this, ignored -> task.run(), delay, interval, unit);
        return scheduled::cancel;
    }
}
