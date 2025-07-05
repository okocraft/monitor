package net.okocraft.monitor.platform.velocity.adapter;

import com.velocitypowered.api.command.CommandSource;
import com.velocitypowered.api.proxy.Player;
import net.kyori.adventure.text.Component;
import net.okocraft.monitor.core.command.CommandSender;
import org.jetbrains.annotations.NotNullByDefault;

import java.util.UUID;

@NotNullByDefault
public final class CommandSenderAdapter {

    public static CommandSender wrap(CommandSource source) {
        return new Wrapper(source);
    }

    private record Wrapper(CommandSource source) implements CommandSender {

        @Override
        public UUID uuid() {
            return this.source instanceof Player player ? player.getUniqueId() : new UUID(0, 0);
        }

        @Override
        public String name() {
            return this.source instanceof Player player ? player.getUsername() : "Console";
        }

        @Override
        public boolean hasPermission(String permission) {
            return this.source.hasPermission(permission);
        }

        @Override
        public void sendPlainMessage(String message) {
            this.source.sendPlainMessage(message);
        }

        @Override
        public void sendMessage(Component message) {
            this.source.sendMessage(message);
        }
    }

    private CommandSenderAdapter() {
        throw new UnsupportedOperationException();
    }
}
