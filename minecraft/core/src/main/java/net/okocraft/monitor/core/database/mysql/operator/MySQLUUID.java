package net.okocraft.monitor.core.database.mysql.operator;

import org.jetbrains.annotations.NotNullByDefault;

import java.nio.ByteBuffer;
import java.util.UUID;

@NotNullByDefault
final class MySQLUUID {

    static byte[] uuidToBytes(UUID uuid) {
        ByteBuffer bb = ByteBuffer.wrap(new byte[16]);
        bb.putLong(uuid.getMostSignificantBits());
        bb.putLong(uuid.getLeastSignificantBits());
        return bb.array();
    }

    static UUID bytesToUUID(byte[] bytes) {
        ByteBuffer bb = ByteBuffer.wrap(bytes);
        return new UUID(bb.getLong(), bb.getLong());
    }

    private MySQLUUID() {
        throw new UnsupportedOperationException();
    }
}
