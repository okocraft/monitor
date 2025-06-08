package net.okocraft.monitor.core.database.operator;

import net.okocraft.monitor.core.models.data.UploadedObject;

import java.sql.Connection;
import java.sql.SQLException;

public interface UploadedObjectTableOperator {

    void insertUploadedObject(Connection connection, UploadedObject object) throws SQLException;

}
