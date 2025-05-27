package net.okocraft.monitor.core.config;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.codec.object.ObjectCodec;

public record CommandConfig(boolean enabled, String customLabel) {

    public static final Codec<CommandConfig> CODEC = ObjectCodec.create(
        CommandConfig::new,
        Codec.BOOLEAN.toFieldCodec("enabled").defaultValue(false).required(CommandConfig::enabled),
        Codec.STRING.toFieldCodec("custom-label").defaultValue("").required(CommandConfig::customLabel)
    );

    public static final CommandConfig EMPTY = new CommandConfig(false, "");

}
