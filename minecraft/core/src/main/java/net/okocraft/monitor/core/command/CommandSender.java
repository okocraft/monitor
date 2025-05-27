package net.okocraft.monitor.core.command;

import net.kyori.adventure.text.Component;
import org.jetbrains.annotations.NotNullByDefault;

import java.util.UUID;

@NotNullByDefault
public interface CommandSender {

    UUID uuid();

    String name();

    boolean hasPermission(String permission);

    void sendPlainMessage(String message);

    void sendMessage(Component message);

}
