package net.okocraft.monitor.platform.paper.listener;

import net.okocraft.monitor.core.config.notification.OreNotification;
import net.okocraft.monitor.core.webhook.discord.DiscordWebhook;
import net.okocraft.monitor.platform.paper.util.VeinFinder;
import org.bukkit.GameMode;
import org.bukkit.Location;
import org.bukkit.Material;
import org.bukkit.World;
import org.bukkit.entity.Player;
import org.bukkit.event.EventHandler;
import org.bukkit.event.EventPriority;
import org.bukkit.event.Listener;
import org.bukkit.event.block.BlockBreakEvent;
import org.bukkit.event.block.BlockPlaceEvent;
import org.jetbrains.annotations.NotNull;

import java.util.EnumSet;
import java.util.Map;
import java.util.UUID;
import java.util.concurrent.ConcurrentHashMap;

public class OreBlockListener implements Listener {

    private static final EnumSet<GameMode> IGNORING_GAME_MODES = EnumSet.of(GameMode.CREATIVE, GameMode.SPECTATOR);

    private final Map<UUID, VeinFinder> finderByWorld = new ConcurrentHashMap<>();
    private final DiscordWebhook webhook;
    private final OreNotification setting;

    public OreBlockListener(@NotNull DiscordWebhook webhook, @NotNull OreNotification setting) {
        this.webhook = webhook;
        this.setting = setting;
    }

    @EventHandler(priority = EventPriority.MONITOR, ignoreCancelled = true)
    public void onPlace(@NotNull BlockPlaceEvent event) {
        if (this.setting.enabledOres().contains(event.getBlockPlaced().getType().key().asString())) {
            this.getVeinFinder(event.getBlockPlaced().getWorld()).recordPlaced(event.getBlockPlaced().getLocation());
        }
    }

    @EventHandler(priority = EventPriority.MONITOR, ignoreCancelled = true)
    public void onBreak(@NotNull BlockBreakEvent event) {
        if (IGNORING_GAME_MODES.contains(event.getPlayer().getGameMode())) {
            return;
        }

        var block = event.getBlock();

        if (!this.setting.enabledOres().contains(event.getBlock().getType().key().asString())) {
            return;
        }

        var finder = this.getVeinFinder(block.getWorld());
        var center = block.getLocation();

        if (finder.isPlacedLocation(center)) {
            return;
        }

        var player = event.getPlayer();
        var type = block.getType();
        var veinCount = finder.findSameBlock(block).count();

        if (0 < veinCount) {
            this.onVeinFound(player, type, center, veinCount);
        }
    }

    private @NotNull VeinFinder getVeinFinder(@NotNull World world) {
        return this.finderByWorld.computeIfAbsent(world.getUID(), ignored -> new VeinFinder(this.setting.maxSearchCount()));
    }

    public void onVeinFound(@NotNull Player player, @NotNull Material type, @NotNull Location center, int count) {
        this.webhook.send(
            this.setting.format()
                .replace("%player_name%", player.getName())
                .replace("%block_type%", this.setting.displayNameMap().getOrDefault(type.key().asString(), type.key().asMinimalString()))
                .replace("%block_world%", center.getWorld().getName())
                .replace("%block_location%", center.getBlockX() + " " + center.getBlockY() + " " + center.getBlockZ())
                .replace("%vein_count%", Integer.toString(count))
        );
    }
}
