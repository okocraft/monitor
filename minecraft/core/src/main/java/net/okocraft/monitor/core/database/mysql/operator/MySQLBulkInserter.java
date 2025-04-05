package net.okocraft.monitor.core.database.mysql.operator;

import org.jetbrains.annotations.NotNullByDefault;

import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

@NotNullByDefault
public final class MySQLBulkInserter {

    private final String baseSql;
    private final String parameters;

    private MySQLBulkInserter(String baseSql, String parameters) {
        this.baseSql = baseSql;
        this.parameters = parameters;
    }

    public static MySQLBulkInserter create(String tableName, List<String> columns) {
        return new MySQLBulkInserter(
            "INSERT INTO " + tableName + " (" + String.join(", ", columns) + ") VALUES ",
            "(" + IntStream.range(0, columns.size()).mapToObj(ignored -> "?").collect(Collectors.joining(", ")) + ")"
        );
    }

    public String createQuery(int records) {
        if (records == 0) {
            throw new IllegalArgumentException("records cannot be zero");
        }
        StringBuilder builder = new StringBuilder(this.baseSql);
        for (int i = 0; i < records; i++) {
            if (i != 0) {
                builder.append(", ");
            }
            builder.append(this.parameters);
        }
        return builder.toString();
    }
}
