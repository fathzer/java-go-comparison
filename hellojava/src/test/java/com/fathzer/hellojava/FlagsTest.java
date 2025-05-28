package com.fathzer.hellojava;

import static org.junit.jupiter.api.Assertions.*;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.stream.Stream;

class FlagsTest {
    @ParameterizedTest(name = "{0}")
    @MethodSource("parseIntegerArgumentCases")
    void testParseIntegerArgument(String description, String[] args, int expected, boolean shouldThrow) {
        if (shouldThrow) {
            assertThrows(IllegalArgumentException.class, 
                () -> Flags.parseIntegerArgument(args, "pl", "piLoops", 2000));
        } else {
            int result = Flags.parseIntegerArgument(args, "pl", "piLoops", 2000);
            assertEquals(expected, result, "Unexpected result for test case: " + description);
        }
    }

    private static Stream<Arguments> parseIntegerArgumentCases() {
        return Stream.of(
            Arguments.of("default", new String[]{"cmd"}, 2000, false),
            Arguments.of("short flag", new String[]{"cmd", "-pl=42"}, 42, false),
            Arguments.of("long flag", new String[]{"cmd", "--piLoops=99"}, 99, false),
            Arguments.of("reject -piLoops", new String[]{"cmd", "-piLoops=5"}, 0, true),
            Arguments.of("reject --pl", new String[]{"cmd", "--pl=5"}, 0, true)
        );
    }

    @ParameterizedTest(name = "{0}")
    @MethodSource("parseStringArgumentCases")
    void testParseStringArgument(String description, String[] args, String expected, boolean shouldThrow) {
        if (shouldThrow) {
            assertThrows(IllegalArgumentException.class,
                () -> Flags.parseStringArgument(args, "f", "fen", "default-value"));
        } else {
            String result = Flags.parseStringArgument(args, "f", "fen", "default-value");
            assertEquals(expected, result, "Unexpected result for test case: " + description);
        }
    }

    private static Stream<Arguments> parseStringArgumentCases() {
        return Stream.of(
            Arguments.of("default", new String[]{"cmd"}, "default-value", false),
            Arguments.of("short flag", new String[]{"cmd", "-f=test"}, "test", false),
            Arguments.of("long flag", new String[]{"cmd", "--fen=rnbqkbnr/8/8/8/8/8/8/RNBQKBNR"}, "rnbqkbnr/8/8/8/8/8/8/RNBQKBNR", false),
            Arguments.of("reject -fen", new String[]{"cmd", "-fen=test"}, null, true),
            Arguments.of("reject --f", new String[]{"cmd", "--f=test"}, null, true)
        );
    }

    @ParameterizedTest(name = "{0}")
    @MethodSource("parseBooleanArgumentCases")
    void testParseBooleanArgument(String description, String[] args, boolean expected, boolean shouldThrow) {
        if (shouldThrow) {
            assertThrows(IllegalArgumentException.class,
                () -> Flags.parseBooleanArgument(args, "b", "blackPlaying"));
        } else {
            boolean result = Flags.parseBooleanArgument(args, "b", "blackPlaying");
            assertEquals(expected, result, "Unexpected result for test case: " + description);
        }
    }

    private static Stream<Arguments> parseBooleanArgumentCases() {
        return Stream.of(
            Arguments.of("default", new String[]{"cmd"}, false, false),
            Arguments.of("short flag", new String[]{"cmd", "-b"}, true, false),
            Arguments.of("long flag", new String[]{"cmd", "--blackPlaying"}, true, false),
            Arguments.of("reject -blackPlaying", new String[]{"cmd", "-blackPlaying"}, false, true),
            Arguments.of("reject --b", new String[]{"cmd", "--b"}, false, true)
        );
    }
}
