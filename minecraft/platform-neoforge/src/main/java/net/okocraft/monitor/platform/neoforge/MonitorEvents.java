package net.okocraft.monitor.platform.neoforge;

import net.kyori.adventure.text.Component;
import net.minecraft.server.level.ServerLevel;
import net.minecraft.server.level.ServerPlayer;
import net.neoforged.bus.api.SubscribeEvent;
import net.neoforged.bus.api.EventPriority;
import net.neoforged.neoforge.event.CommandEvent;
import net.neoforged.neoforge.event.ServerChatEvent;
import net.neoforged.neoforge.event.entity.player.AnvilCraftEvent;
import net.neoforged.neoforge.event.entity.player.PlayerEvent;
import net.neoforged.neoforge.event.level.LevelEvent;
import net.okocraft.monitor.core.handler.PlayerHandler;
import net.okocraft.monitor.core.handler.WorldHandler;
import net.okocraft.monitor.core.models.logs.PlayerConnectLog;
import net.okocraft.monitor.core.models.logs.PlayerEditSignLog;
import net.okocraft.monitor.platform.neoforge.mixin.ItemCombinerMenuAccess;

import java.util.Arrays;
import java.util.List;

final class MonitorEvents {
    private final PlayerHandler players;
    private final WorldHandler worlds;

    MonitorEvents(PlayerHandler players, WorldHandler worlds) {
        this.players = players;
        this.worlds = worlds;
    }

    void load(ServerLevel level) {
        worlds.onLoad(DimensionIds.uuid(level), DimensionIds.name(level));
    }

    @SubscribeEvent
    public void onLevelLoad(LevelEvent.Load event) {
        if (event.getLevel() instanceof ServerLevel level) {
            load(level);
        }
    }

    @SubscribeEvent
    public void onLogin(PlayerEvent.PlayerLoggedInEvent event) {
        if (event.getEntity() instanceof ServerPlayer player) {
            players.onJoin(player.getUUID(), player.getGameProfile().name(), player.connection.getRemoteAddress());
        }
    }

    @SubscribeEvent
    public void onLogout(PlayerEvent.PlayerLoggedOutEvent event) {
        if (event.getEntity() instanceof ServerPlayer player) {
            var details = player.connection.getConnection().getDisconnectionDetails();
            var action = DisconnectAction.fromDetails(details);
            Component reason = action == PlayerConnectLog.Action.KICKED && details != null
                ? NativeText.toAdventure(details.reason(), player.level().registryAccess()) : Component.empty();
            players.onLeave(player.getUUID(), action, player.connection.getRemoteAddress(), reason);
        }
    }

    @SubscribeEvent(priority = EventPriority.LOWEST, receiveCanceled = true)
    public void onChat(ServerChatEvent event) {
        ServerPlayer player = event.getPlayer();
        players.onChat(player.getUUID(), DimensionIds.uuid(player.level()), DimensionIds.position(player.blockPosition()), Component.text(event.getRawText()));
    }

    @SubscribeEvent(priority = EventPriority.LOWEST, receiveCanceled = true)
    public void onCommand(CommandEvent event) {
        ServerPlayer player = event.getParseResults().getContext().getSource().getPlayer();
        if (player != null) {
            players.onWorldCommand(player.getUUID(), DimensionIds.uuid(player.level()), DimensionIds.position(player.blockPosition()), "/" + event.getParseResults().getReader().getString());
        }
    }

    @SubscribeEvent
    public void onAnvil(AnvilCraftEvent.Post event) {
        if (!(event.getEntity() instanceof ServerPlayer player)) {
            return;
        }
        var input = event.getLeft();
        var output = event.getOutput();
        if (output.isEmpty()) {
            return;
        }
        var originalName = input.getCustomName();
        var resultName = output.getCustomName();
        if (java.util.Objects.equals(resultName, originalName)) {
            return;
        }
        var itemKey = net.minecraft.core.registries.BuiltInRegistries.ITEM.getKey(output.getItem());
        var location = ((ItemCombinerMenuAccess) event.getMenu()).monitor$access()
            .evaluate((level, pos) -> pos, player.blockPosition());
        players.onRenameItem(player.getUUID(), DimensionIds.uuid(player.level()), DimensionIds.position(location),
            net.kyori.adventure.key.Key.key(itemKey.toString()),
            resultName == null ? Component.empty() : NativeText.toAdventure(resultName, player.level().registryAccess()), output.getCount());
    }

    void signEdit(net.minecraft.world.entity.player.Player player,
                  net.minecraft.world.level.block.entity.SignBlockEntity sign,
                  boolean frontText) {
        if (!(player instanceof ServerPlayer serverPlayer) || !(sign.getLevel() instanceof ServerLevel level)) {
            return;
        }
        var key = net.minecraft.core.registries.BuiltInRegistries.BLOCK.getKey(sign.getBlockState().getBlock());
        List<Component> lines = Arrays.stream(sign.getText(frontText).getMessages(player.isTextFilteringEnabled()))
            .map(line -> NativeText.toAdventure(line, level.registryAccess())).toList();
        players.onEditSign(serverPlayer.getUUID(), DimensionIds.uuid(level), DimensionIds.position(sign.getBlockPos()),
            net.kyori.adventure.key.Key.key(key.toString()), frontText ? PlayerEditSignLog.Side.FRONT : PlayerEditSignLog.Side.BACK, lines);
    }

    @SubscribeEvent
    public void onSignEdited(SignEditedEvent event) {
        signEdit(event.player(), event.sign(), event.front());
    }
}
