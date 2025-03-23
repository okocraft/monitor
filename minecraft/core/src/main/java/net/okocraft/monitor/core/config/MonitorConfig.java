package net.okocraft.monitor.core.config;

import dev.siroshun.configapi.core.node.MapNode;
import dev.siroshun.configapi.core.serialization.record.RecordSerialization;
import dev.siroshun.configapi.format.yaml.YamlFormat;
import org.jetbrains.annotations.NotNullByDefault;

import java.nio.file.Path;
import java.util.concurrent.atomic.AtomicReference;

@NotNullByDefault
public record MonitorConfig(DatabaseConfig database) {

    private static final RecordSerialization<MonitorConfig> SERIALIZATION = RecordSerialization.create(MonitorConfig.class);

    public static Holder load(Path filepath) throws Exception {
        var loaded = YamlFormat.COMMENT_PROCESSING.load(filepath);
        var defaultConfig = SERIALIZATION.serializer().serializeDefault(MonitorConfig.class);
        applyDefaults(loaded, defaultConfig);
        YamlFormat.COMMENT_PROCESSING.save(loaded, filepath);
        return new Holder(filepath, SERIALIZATION.deserializer().deserialize(loaded));
    }

    private static void applyDefaults(MapNode target, MapNode defaults) {
        for (var entry : defaults.value().entrySet()) {
            var key = entry.getKey();
            var value = entry.getValue();
            if (!target.containsKey(key)) {
                target.set(key, value);
                continue;
            }

            if (value instanceof MapNode defaultChild &&
                target.get(key) instanceof MapNode targetChild) {
                applyDefaults(targetChild, defaultChild);
            }
        }
    }

    public static final class Holder {

        private final Path filepath;
        private final AtomicReference<MonitorConfig> ref;

        public Holder(Path filepath, MonitorConfig initial) {
            this.filepath = filepath;
            this.ref = new AtomicReference<>(initial);
        }

        public MonitorConfig get() {
            return this.ref.get();
        }

        public void reload() throws Exception {
            var loaded = YamlFormat.DEFAULT.load(this.filepath);
            this.ref.set(SERIALIZATION.deserializer().deserialize(loaded));
        }
    }
}
