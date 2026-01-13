package net.okocraft.monitor.core.config.notification;

import dev.siroshun.codec4j.api.codec.Codec;
import dev.siroshun.codec4j.api.decoder.Decoder;
import dev.siroshun.codec4j.api.decoder.collection.MapDecoder;
import dev.siroshun.codec4j.api.decoder.collection.SetDecoder;
import dev.siroshun.codec4j.api.decoder.object.FieldDecoder;
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
        FieldDecoder.optional("thread-id", Codec.LONG, 0L),
        FieldDecoder.required("enabled-ores", SetDecoder.create(Codec.STRING)),
        FieldDecoder.required("max-search-count", Codec.INT),
        FieldDecoder.required("format", Codec.STRING),
        FieldDecoder.required("display-name-map", MapDecoder.create(Codec.STRING, Codec.STRING))
    );

    public static final OreNotification EMPTY = new OreNotification(
        0L,
        Set.of(), 100,
        "%player_name% found a vein of **%block_type%**x%vein_count% at %block_location% in %block_world%", Map.of()
    );
}
