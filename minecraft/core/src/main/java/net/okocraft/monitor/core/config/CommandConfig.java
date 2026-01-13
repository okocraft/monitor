package net.okocraft.monitor.core.config;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.decoder.Decoder;
import dev.siroshun.codec4j.api.decoder.object.FieldDecoder;
import dev.siroshun.codec4j.api.decoder.object.ObjectDecoder;

public record CommandConfig(boolean enabled, String customLabel) {

    public static final Decoder<CommandConfig> CODEC = ObjectDecoder.create(
        CommandConfig::new,
        FieldDecoder.optional("enabled", Codec.BOOLEAN, false),
        FieldDecoder.optional("custom-label", Codec.STRING, "")
    );

    public static final CommandConfig EMPTY = new CommandConfig(false, "");

}
