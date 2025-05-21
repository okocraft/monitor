package net.okocraft.monitor.platform.paper.listener;

import io.papermc.paper.datacomponent.DataComponentTypes;
import io.papermc.paper.event.player.AsyncChatEvent;
import net.kyori.adventure.text.Component;
import net.okocraft.monitor.core.handler.PlayerHandler;
import net.okocraft.monitor.core.models.logs.PlayerConnectLog;
import net.okocraft.monitor.core.models.logs.PlayerEditSignLog;
import net.okocraft.monitor.platform.paper.adapter.PositionAdapter;
import org.bukkit.Location;
import org.bukkit.block.Block;
import org.bukkit.block.Sign;
import org.bukkit.entity.Player;
import org.bukkit.event.EventHandler;
import org.bukkit.event.EventPriority;
import org.bukkit.event.Listener;
import org.bukkit.event.block.SignChangeEvent;
import org.bukkit.event.inventory.InventoryClickEvent;
import org.bukkit.event.player.PlayerCommandPreprocessEvent;
import org.bukkit.event.player.PlayerJoinEvent;
import org.bukkit.event.player.PlayerKickEvent;
import org.bukkit.event.player.PlayerQuitEvent;
import org.bukkit.inventory.AnvilInventory;
import org.bukkit.inventory.ItemStack;

import java.util.Objects;

public class PlayerListener implements Listener {

    private static final int ANVIL_INPUT_SLOT = 0;
    private static final int ANVIL_OUTPUT_SLOT = 2;

    private final PlayerHandler handler;

    public PlayerListener(PlayerHandler handler) {
        this.handler = handler;
    }

    @EventHandler
    public void onJoin(PlayerJoinEvent event) {
        Player player = event.getPlayer();
        this.handler.onJoin(player.getUniqueId(), player.getName(), player.getAddress());
    }

    @EventHandler
    public void onKick(PlayerKickEvent event) {
        Player player = event.getPlayer();
        this.handler.onLeave(player.getUniqueId(), PlayerConnectLog.Action.KICKED, player.getAddress(), event.reason());
    }

    @EventHandler
    public void onQuit(PlayerQuitEvent event) {
        if (event.getReason() == PlayerQuitEvent.QuitReason.KICKED) {
            return;
        }

        Player player = event.getPlayer();
        PlayerConnectLog.Action action = switch (event.getReason()) {
            case DISCONNECTED -> PlayerConnectLog.Action.DISCONNECT;
            case KICKED -> throw new AssertionError();
            case TIMED_OUT -> PlayerConnectLog.Action.TIMED_OUT;
            case ERRONEOUS_STATE -> PlayerConnectLog.Action.ERRONEOUS_STATE;
        };
        this.handler.onLeave(player.getUniqueId(), action, player.getAddress(), Component.empty());
    }

    @EventHandler(priority = EventPriority.MONITOR)
    public void onChat(AsyncChatEvent event) {
        Player player = event.getPlayer();
        this.handler.onChat(player.getUniqueId(), player.getWorld().getUID(), PositionAdapter.fromLocation(player.getLocation()), event.originalMessage());
    }

    @EventHandler(priority = EventPriority.MONITOR)
    public void onCommand(PlayerCommandPreprocessEvent event) {
        Player player = event.getPlayer();
        this.handler.onWorldCommand(player.getUniqueId(), player.getWorld().getUID(), PositionAdapter.fromLocation(player.getLocation()), event.getMessage());
    }

    @EventHandler(priority = EventPriority.MONITOR, ignoreCancelled = true)
    public void onRenameItem(InventoryClickEvent event) {
        if (!(event.getWhoClicked() instanceof Player player)) {
            return;
        }

        if (!(event.getClickedInventory() instanceof AnvilInventory anvilInventory)) {
            return;
        }

        int clickedSlotIndex = event.getRawSlot();
        if (clickedSlotIndex != ANVIL_OUTPUT_SLOT) {
            return;
        }

        ItemStack inputItem = anvilInventory.getItem(ANVIL_INPUT_SLOT);

        ItemStack outputItem = event.getCurrentItem();
        if (outputItem == null) {
            return;
        }

        Component originalName = inputItem == null ? null : inputItem.getData(DataComponentTypes.CUSTOM_NAME);
        Component resultName = outputItem.getData(DataComponentTypes.CUSTOM_NAME);

        if ((originalName == null && resultName == null) ||
            (originalName != null && originalName.equals(resultName))) {
            return;
        }

        Location location = Objects.requireNonNullElseGet(anvilInventory.getLocation(), player::getLocation);
        this.handler.onRenameItem(player.getUniqueId(), player.getWorld().getUID(), PositionAdapter.fromLocation(location), outputItem.getType().key(), resultName, outputItem.getAmount());
    }

    @EventHandler(priority = EventPriority.MONITOR, ignoreCancelled = true)
    public void onEditSign(SignChangeEvent event) {
        PlayerEditSignLog.Side side = switch (event.getSide()) {
            case BACK -> PlayerEditSignLog.Side.BACK;
            case FRONT -> PlayerEditSignLog.Side.FRONT;
        };

        Block block = event.getBlock();
        if (!(block.getState() instanceof Sign sign)) {
            return;
        }

        if (sign.getSide(event.getSide()).lines().equals(event.lines())) {
            return;
        }

        this.handler.onEditSign(
            event.getPlayer().getUniqueId(), block.getWorld().getUID(), PositionAdapter.fromLocation(block.getLocation()),
            block.getType().key(), side, event.lines()
        );
    }
}
