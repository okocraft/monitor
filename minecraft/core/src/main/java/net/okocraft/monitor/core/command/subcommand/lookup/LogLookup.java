package net.okocraft.monitor.core.command.subcommand.lookup;

import net.okocraft.monitor.core.models.lookup.CommonLookupParams;

import java.sql.SQLException;
import java.util.function.Consumer;

public record LogLookup<P extends CommonLookupParams, T>(String type, ParamParser<P> parser, CountByParam<P> countByParam, LookupByParam<P, T> lookupByParam) {

    public interface CountByParam<P extends CommonLookupParams> {
        long count(P params) throws SQLException;
    }

    public interface LookupByParam<P extends CommonLookupParams, T> {
        void lookup(P params, Consumer<T> consumer) throws SQLException;
    }

}
