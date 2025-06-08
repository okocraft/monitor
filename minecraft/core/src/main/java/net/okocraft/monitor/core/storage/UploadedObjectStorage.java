package net.okocraft.monitor.core.storage;

import net.okocraft.monitor.core.database.Database;
import net.okocraft.monitor.core.database.operator.Operators;
import net.okocraft.monitor.core.models.data.UploadedObject;

import java.sql.Connection;
import java.sql.SQLException;

public class UploadedObjectStorage {

    private final Database database;
    private final Operators operators;

    public UploadedObjectStorage(Database database) {
        this.database = database;
        this.operators = database.getOperators();
    }

    public void recordUploadedObject(UploadedObject object) throws SQLException {
        try (Connection connection = this.database.getConnection()) {
            this.operators.uploadedObjects().insertUploadedObject(connection, object);
        }
    }
}
