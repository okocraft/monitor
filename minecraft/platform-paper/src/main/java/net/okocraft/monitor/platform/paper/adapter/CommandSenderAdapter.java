package net.okocraft.monitor.platform.paper.adapter;

import io.papermc.paper.command.brigadier.CommandSourceStack;
import net.kyori.adventure.text.Component;
import net.okocraft.monitor.core.command.CommandSender;
import org.bukkit.entity.Player;
import org.jetbrains.annotations.NotNullByDefault;

import java.util.UUID;

@NotNullByDefault
public final class CommandSenderAdapter {

    public static CommandSender wrap(CommandSourceStack source) {
        return new Wrapper(source);
    }

    private record Wrapper(CommandSourceStack source) implements CommandSender {

        @Override
        public UUID uuid() {
            return this.source.getSender() instanceof Player player ? player.getUniqueId() : new UUID(0, 0);
        }

        @Override
        public String name() {
            return this.source.getSender().getName();
        }

        @Override
        public boolean hasPermission(String permission) {
            return this.source.getSender().hasPermission(permission);
        }

        @Override
        public void sendPlainMessage(String message) {
            this.source.getSender().sendPlainMessage(message);
        }

        @Override
        public void sendMessage(Component message) {
            this.source.getSender().sendMessage(message);
        }
    }

    private CommandSenderAdapter() {
        throw new UnsupportedOperationException();
    }

}
