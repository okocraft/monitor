package net.okocraft.monitor.core.models.lookup;

import java.time.LocalDateTime;

public interface CommonLookupParams {

    LocalDateTime start();

    LocalDateTime end();

    long offset();

    long limit();

    record Record(LocalDateTime start, LocalDateTime end, long limit, long offset) implements CommonLookupParams {
    }

}
