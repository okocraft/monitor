package net.okocraft.monitor.core.command;

import org.jetbrains.annotations.NotNullByDefault;
import org.jetbrains.annotations.Nullable;

import java.util.Collections;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Locale;
import java.util.Map;
import java.util.function.Predicate;

@NotNullByDefault
public class SubCommandMap {

    public static Builder builder() {
        return new Builder();
    }

    private final Map<String, RegisteredCommand> commandMap;

    private SubCommandMap(Map<String, RegisteredCommand> commandMap) {
        this.commandMap = Collections.unmodifiableMap(commandMap);
    }

    public @Nullable Command find(String subcommand) {
        RegisteredCommand command = this.commandMap.get(subcommand.toLowerCase(Locale.ENGLISH));
        return command == null ? null : command.command();
    }

    public @Nullable Command findWithPermission(String subCommand, CommandSender sender) {
        RegisteredCommand command = this.commandMap.get(subCommand.toLowerCase(Locale.ENGLISH));
        return command == null || !sender.hasPermission(command.permission()) ? null : command.command();
    }

    public List<String> findLabelsWithPermission(String prefix, CommandSender sender) {
        String filter = prefix.toLowerCase(Locale.ENGLISH);
        return this.commandMap.values().stream()
            .filter(Predicate.not(RegisteredCommand::isAlias))
            .filter(command -> command.label().startsWith(filter))
            .filter(command -> sender.hasPermission(command.permission()))
            .map(RegisteredCommand::label)
            .toList();
    }

    public static class Builder {
        private final Map<String, RegisteredCommand> commandMap = new LinkedHashMap<>();

        public Builder add(String label, String permission, Command command) {
            String lowerCaseLabel = label.toLowerCase(Locale.ENGLISH);
            this.commandMap.put(lowerCaseLabel, new RegisteredCommand(lowerCaseLabel, permission, command, false));
            return this;
        }

        public Builder add(String label, String permission, Command command, String... aliases) {
            this.add(label, permission, command);
            for (String alias : aliases) {
                String lowerCaseAlias = alias.toLowerCase(Locale.ENGLISH);
                this.commandMap.put(lowerCaseAlias, new RegisteredCommand(lowerCaseAlias, permission, command, true));
            }
            return this;
        }

        public SubCommandMap build() {
            return new SubCommandMap(this.commandMap);
        }
    }

    private record RegisteredCommand(String label, String permission, Command command, boolean isAlias) {
    }
}
