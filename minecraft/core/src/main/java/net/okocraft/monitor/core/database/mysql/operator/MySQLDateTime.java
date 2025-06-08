package net.okocraft.monitor.core.database.mysql.operator;

import java.sql.Timestamp;
import java.time.Instant;
import java.time.LocalDateTime;

final class MySQLDateTime {

    static Timestamp now() {
        return Timestamp.valueOf(LocalDateTime.now());
    }

    static Timestamp from(LocalDateTime localDateTime) {
        return Timestamp.valueOf(localDateTime);
    }

    static Timestamp from(Instant instant) {
        return Timestamp.from(instant);
    }

    private MySQLDateTime() {
        throw new UnsupportedOperationException();
    }
}
