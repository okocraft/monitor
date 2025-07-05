package net.okocraft.monitor.core.command.subcommand.lookup;

import dev.siroshun.jfun.result.Result;

import java.time.Duration;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;
import java.time.format.DateTimeFormatter;
import java.time.format.DateTimeParseException;
import java.time.temporal.ChronoUnit;
import java.util.Map;
import java.util.regex.Matcher;
import java.util.regex.Pattern;
import java.util.stream.Collectors;

final class ParseUtils {

    static Result<LocalDateTime, ParamParseError> parseAsDateTime(String arg, boolean includeEndOfDay) {
        try {
            return Result.success(DateTimeFormatter.ISO_LOCAL_DATE_TIME.parse(arg, LocalDateTime::from));
        } catch (DateTimeParseException ignored) {
        }

        try {
            LocalDate date = DateTimeFormatter.ISO_LOCAL_DATE.parse(arg, LocalDate::from);
            return Result.success(includeEndOfDay ? date.atTime(LocalTime.MAX) : date.atStartOfDay());
        } catch (DateTimeParseException ignored) {
        }

        try {
            return Result.success(DateTimeFormatter.ISO_LOCAL_TIME.parse(arg, LocalTime::from).atDate(LocalDate.now()));
        } catch (DateTimeParseException ignored) {
        }

        return Result.failure(new ParamParseError("unable to parse datetime: " + arg));
    }

    private static final Map<ChronoUnit, String> UNITS_PATTERNS = Map.of(
        ChronoUnit.YEARS, "y(?:ear)?s?",
        ChronoUnit.MONTHS, "mo(?:nth)?s?",
        ChronoUnit.WEEKS, "w(?:eek)?s?",
        ChronoUnit.DAYS, "d(?:ay)?s?",
        ChronoUnit.HOURS, "h(?:our|r)?s?",
        ChronoUnit.MINUTES, "m(?:inute|in)?s?",
        ChronoUnit.SECONDS, "s(?:econd|ec)?s?"
    );

    private static final ChronoUnit[] UNITS = UNITS_PATTERNS.keySet().toArray(new ChronoUnit[0]);

    private static final Pattern PATTERN = Pattern.compile(
        UNITS_PATTERNS.values().stream()
            .map(pattern -> "(?:(\\d+)\\s*" + pattern + "[,\\s]*)?")
            .collect(Collectors.joining("", "^\\s*", "$")),
        Pattern.CASE_INSENSITIVE
    );

    static Result<Duration, ParamParseError> parseAsDuration(String arg) {
        Matcher matcher = PATTERN.matcher(arg);
        if (!matcher.matches()) {
            return Result.failure(new ParamParseError("unable to parse duration: " + arg));
        }

        Duration duration = Duration.ZERO;

        for (int i = 0; i < UNITS.length; i++) {
            ChronoUnit unit = UNITS[i];
            int g = i + 1;

            if (matcher.group(g) != null && !matcher.group(g).isEmpty()) {
                int n = Integer.parseInt(matcher.group(g));
                if (n > 0) {
                    duration = duration.plus(unit.getDuration().multipliedBy(n));
                }
            }
        }

        return Result.success(duration);
    }
    
    private ParseUtils() {
        throw new UnsupportedOperationException();
    }
}
