package net.okocraft.monitor.core.command.subcommand.lookup;

import org.jetbrains.annotations.NotNullByDefault;
import org.jetbrains.annotations.Nullable;

import java.time.LocalDateTime;
import java.util.HashMap;
import java.util.Map;

@NotNullByDefault
public class ParamContext {

    public static final TypedKey<LocalDateTime> KEY_START = new TypedKey<>("start", LocalDateTime.class);
    public static final TypedKey<LocalDateTime> KEY_END = new TypedKey<>("end", LocalDateTime.class);
    public static final TypedKey<Long> KEY_PAGE = new TypedKey<>("page", Long.class);

    public static ParamContext create() {
        return new ParamContext();
    }

    private final Map<TypedKey<?>, Object> params = new HashMap<>();

    private ParamContext() {
    }

    public boolean has(TypedKey<?> key) {
        return this.params.containsKey(key);
    }

    public <T> @Nullable T get(TypedKey<T> key) {
        Object value = this.params.get(key);
        return value != null ? key.type().cast(value) : null;
    }

    public <T> T getOrDefault(TypedKey<T> key, T defaultValue) {
        Object value = this.params.get(key);
        if (value == null) {
            return defaultValue;
        }
        return key.type().cast(value);
    }

    public <T> void put(TypedKey<T> key, T value) {
        this.params.put(key, value);
    }

    public record TypedKey<T>(String key, Class<T> type) {
    }
}
