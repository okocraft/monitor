package net.okocraft.monitor.core.config.notification;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.decoder.Decoder;
import dev.siroshun.codec4j.api.decoder.object.ObjectDecoder;

import java.util.Map;
import java.util.Set;

public record OreNotification(
    long threadId,
    Set<String> enabledOres, int maxSearchCount,
    String format, Map<String, String> displayNameMap
) {

    public static final Decoder<OreNotification> DECODER = ObjectDecoder.create(
        OreNotification::new,
        Codec.LONG.toOptionalFieldDecoder("thread-id", 0L),
        Codec.STRING.toSetDecoder().toRequiredFieldDecoder("enabled-ores"),
        Codec.INT.toRequiredFieldDecoder("max-search-count"),
        Codec.STRING.toRequiredFieldDecoder("format"),
        Codec.STRING.toMapCodecAsKey(Codec.STRING).toRequiredFieldDecoder("display-name-map")
    );

    public static final OreNotification EMPTY = new OreNotification(
        0L,
        Set.of(), 100,
        "%player_name% found a vein of **%block_type%**x%vein_count% at %block_location% in %block_world%", Map.of()
    );
}
