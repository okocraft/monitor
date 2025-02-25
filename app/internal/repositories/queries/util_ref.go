package queries

func tableColumnRef(tableName string, columnName string) string {
	return tableName + "." + columnName
}
