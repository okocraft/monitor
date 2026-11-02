package net.okocraft.monitor.platform.neoforge;

import com.mojang.brigadier.arguments.StringArgumentType;
import net.minecraft.commands.Commands;
import net.minecraft.server.MinecraftServer;
import net.neoforged.fml.common.Mod;
import net.neoforged.fml.ModList;
import net.neoforged.fml.loading.FMLPaths;
import net.neoforged.neoforge.common.NeoForge;
import net.neoforged.neoforge.event.RegisterCommandsEvent;
import net.neoforged.neoforge.event.server.ServerStartedEvent;
import net.neoforged.neoforge.event.server.ServerStoppingEvent;
import net.okocraft.monitor.core.Monitor;
import net.okocraft.monitor.core.bootstrap.MonitorBootstrap;
import net.okocraft.monitor.core.command.Command;
import net.okocraft.monitor.core.config.notification.OreNotification;
import net.okocraft.monitor.core.config.notification.ServerStatusNotification;
import net.okocraft.monitor.core.handler.Handlers;
import net.okocraft.monitor.core.platform.CancellableTask;
import net.okocraft.monitor.core.platform.PlatformAdapter;
import net.okocraft.monitor.core.webhook.discord.DiscordWebhook;
import org.slf4j.LoggerFactory;

import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.TimeUnit;

@Mod("monitor")
public final class MonitorNeoForge implements PlatformAdapter {
    private final ScheduledExecutorService scheduler = Executors.newSingleThreadScheduledExecutor(r -> {
        Thread thread = new Thread(r, "Monitor-logging");
        thread.setDaemon(true);
        return thread;
    });
    private MinecraftServer server;
    private Monitor monitor;
    private MonitorEvents events;
    private OreEvents oreEvents;
    private String commandLabel;
    private Command command;

    public MonitorNeoForge() {
        NeoForge.EVENT_BUS.addListener(this::start);
        NeoForge.EVENT_BUS.addListener(this::stop);
        NeoForge.EVENT_BUS.addListener(this::registerCommands);
    }

    private void start(ServerStartedEvent event) {
        server = event.getServer();
        var loaded = MonitorBootstrap.load(LoggerFactory.getLogger("monitor"), FMLPaths.CONFIGDIR.get().resolve("monitor"));
        if (loaded.isEmpty()) {
            return;
        }
        monitor = loaded.get();
        monitor.start(this);
    }

    private void stop(ServerStoppingEvent event) {
        if (monitor != null) {
            monitor.shutdown(this);
            monitor = null;
        }
        scheduler.shutdown();
    }

    @Override
    public String pluginVersion() {
        return ModList.get().getModContainerById("monitor")
            .map(container -> container.getModInfo().getVersion().toString())
            .orElse("unknown");
    }

    @Override
    public void registerEventListeners(Handlers handlers) {
        events = new MonitorEvents(handlers.player(), handlers.world());
        for (var level : server.getAllLevels()) {
            events.load(level);
        }
        NeoForge.EVENT_BUS.register(events);
    }

    @Override
    public void registerVeinFindListener(OreNotification setting, DiscordWebhook webhook) {
        if (webhook != null && !setting.enabledOres().isEmpty() && !setting.format().isEmpty()) {
            oreEvents = new OreEvents(setting, webhook);
            NeoForge.EVENT_BUS.register(oreEvents);
        }
    }

    @Override
    public void unregisterEventListeners() {
        if (events != null) {
            NeoForge.EVENT_BUS.unregister(events);
            events = null;
        }
        if (oreEvents != null) {
            NeoForge.EVENT_BUS.unregister(oreEvents);
            oreEvents = null;
        }
    }

    @Override
    public void registerCommand(String label, Command command) {
        this.commandLabel = label;
        this.command = command;
        installCommand(server.getCommands().getDispatcher());
    }

    private void registerCommands(RegisterCommandsEvent event) {
        if (command != null) {
            installCommand(event.getDispatcher());
        }
    }

    private void installCommand(com.mojang.brigadier.CommandDispatcher<net.minecraft.commands.CommandSourceStack> dispatcher) {
        dispatcher.register(Commands.literal(commandLabel)
            .executes(context -> runCommand(context.getSource(), new String[0]))
            .then(Commands.argument("arguments", StringArgumentType.greedyString())
                .executes(context -> runCommand(context.getSource(), StringArgumentType.getString(context, "arguments").split("\\s+")))
                .suggests((context, builder) -> {
                    String arguments = context.getInput().substring(commandLabel.length()).trim();
                    String[] args = arguments.isEmpty() ? new String[]{""} : arguments.split("\\s+", -1);
                    var currentToken = builder.createOffset(builder.getInput().lastIndexOf(' ') + 1);
                    return command.tabComplete(new MonitorCommandSource(context.getSource()), args)
                        .thenApply(suggestions -> {
                            suggestions.forEach(currentToken::suggest);
                            return currentToken.build();
                        });
                })));
    }

    private int runCommand(net.minecraft.commands.CommandSourceStack source, String[] args) {
        command.execute(new MonitorCommandSource(source), args);
        return 1;
    }

    @Override
    public CancellableTask startServerStatusChecker(ServerStatusNotification notification, DiscordWebhook webhook) {
        return () -> { };
    }

    @Override
    public CancellableTask scheduleTask(Runnable task, long delay, long interval, TimeUnit unit) {
        var future = scheduler.scheduleAtFixedRate(task, delay, interval, unit);
        return () -> future.cancel(false);
    }
}
