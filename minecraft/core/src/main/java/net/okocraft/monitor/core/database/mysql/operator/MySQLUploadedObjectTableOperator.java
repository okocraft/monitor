package net.okocraft.monitor.core.database.mysql.operator;

import net.okocraft.monitor.core.database.operator.UploadedObjectTableOperator;
import net.okocraft.monitor.core.models.data.UploadedObject;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.SQLException;

public class MySQLUploadedObjectTableOperator implements UploadedObjectTableOperator {
    @Override
    public void insertUploadedObject(Connection connection, UploadedObject object) throws SQLException {
        try (PreparedStatement statement = connection.prepareStatement("""
            INSERT INTO minecraft_uploaded_objects (object_id, type, version, created_by_uuid, created_by_name, expires_at, created_at)
            VALUES (?, ?, ?, ?, ?, ?, ?)
            """)) {
            statement.setBytes(1, MySQLUUID.uuidToBytes(object.id()));
            statement.setInt(2, object.type());
            statement.setInt(3, object.version());
            statement.setBytes(4, MySQLUUID.uuidToBytes(object.createdByUUID()));
            statement.setString(5, object.createdByName());
            statement.setTimestamp(6, MySQLDateTime.from(object.expiresAt()));
            statement.setTimestamp(7, MySQLDateTime.from(object.createdTime()));
            statement.execute();
        }
    }
}
