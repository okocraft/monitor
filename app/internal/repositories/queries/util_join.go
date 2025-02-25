package queries

func joinCondition(tableName string, columnName string, targetTableName string, targetColumnName string) string {
	return tableName + "." + columnName + " = " + targetTableName + "." + targetColumnName
}
