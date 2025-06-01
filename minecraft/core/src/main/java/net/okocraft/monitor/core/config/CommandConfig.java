package net.okocraft.monitor.core.config;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.decoder.Decoder;
import dev.siroshun.codec4j.api.decoder.object.ObjectDecoder;

public record CommandConfig(boolean enabled, String customLabel) {

    public static final Decoder<CommandConfig> CODEC = ObjectDecoder.create(
        CommandConfig::new,
        Codec.BOOLEAN.toOptionalFieldDecoder("enabled", false),
        Codec.STRING.toOptionalFieldDecoder("custom-label", "")
    );

    public static final CommandConfig EMPTY = new CommandConfig(false, "");

}
