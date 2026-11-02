package net.okocraft.monitor.platform.neoforge;

import net.kyori.adventure.text.Component;
import net.minecraft.commands.CommandSourceStack;
import net.minecraft.server.level.ServerPlayer;
import net.minecraft.server.players.NameAndId;
import net.minecraft.server.permissions.Permissions;
import net.neoforged.fml.ModList;
import net.okocraft.monitor.core.command.CommandSender;

import java.util.UUID;

record MonitorCommandSource(CommandSourceStack source) implements CommandSender {
    @Override
    public UUID uuid() {
        ServerPlayer player = source.getPlayer();
        return player == null ? new UUID(0, 0) : player.getUUID();
    }

    @Override
    public String name() {
        return source.getTextName();
    }

    @Override
    public boolean hasPermission(String permission) {
        ServerPlayer player = source.getPlayer();
        if (player == null) {
            return source.permissions().hasPermission(Permissions.COMMANDS_ADMIN);
        }
        if (ModList.get().isLoaded("luckperms")) {
            Boolean granted = LuckPermsPermission.has(player, permission);
            if (granted != null) {
                return granted;
            }
        }
        return player.level().getServer().getPlayerList().isOp(new NameAndId(player.getGameProfile()));
    }

    @Override
    public void sendPlainMessage(String message) {
        source.sendSystemMessage(net.minecraft.network.chat.Component.literal(message));
    }

    @Override
    public void sendMessage(Component message) {
        source.sendSystemMessage(NativeText.toMinecraft(message, source.getServer().registryAccess()));
    }

    private static final class LuckPermsPermission {
        static Boolean has(ServerPlayer player, String permission) {
            var api = net.luckperms.api.LuckPermsProvider.get();
            var user = api.getUserManager().getUser(player.getUUID());
            if (user == null) {
                return null;
            }
            var options = api.getContextManager().getQueryOptions(user)
                .orElse(api.getContextManager().getStaticQueryOptions());
            var result = user.getCachedData().getPermissionData(options).checkPermission(permission);
            return result == net.luckperms.api.util.Tristate.UNDEFINED ? null : result.asBoolean();
        }
    }
}
