package com.fathzer.hellojava;

import java.util.Arrays;
import java.util.List;

public class Flags {
    private Flags() {}

    public static int parseIntegerArgument(String[] args, String shortPrefix, String longPrefix, int defaultValue) {
        checkIllegalUsage(args, shortPrefix, longPrefix);
        int arg = defaultValue;
        arg = parseIntegerArgument(args, "--"+longPrefix+"=", arg);
        arg = parseIntegerArgument(args, "-"+shortPrefix+"=", arg);
        return arg;
    }

    private static int parseIntegerArgument(String[] args, String prefix, int defaultValue) {
        for (String arg : args) {
            if (arg.startsWith(prefix)) {
                try {
                    return Integer.parseInt(arg.substring(prefix.length()));
                } catch (NumberFormatException e) {
                    throw new IllegalArgumentException("Invalid value for " + prefix + ": " + arg, e);
                }
            }
        }
        return defaultValue;
    }

    private static void checkIllegalUsage(String[] args, String shortPrefix, String longPrefix) {
        for (String arg : args) {
            if (arg.startsWith("--"+shortPrefix+"=") || arg.startsWith("-"+longPrefix+"=")) {
                throw new IllegalArgumentException(String.format("%s is not supported. Use -%s or --%s instead", arg, shortPrefix, longPrefix));
            }
        }
    }

    public static String parseStringArgument(String[] args, String shortPrefix, String longPrefix, String defaultValue) {
        checkIllegalUsage(args, shortPrefix, longPrefix);
        String arg = defaultValue;
        arg = parseStringArgument(args, "--"+longPrefix+"=", arg);
        arg = parseStringArgument(args, "-"+shortPrefix+"=", arg);
        return arg;
    }

    private static String parseStringArgument(String[] args, String prefix, String defaultValue) {
        for (String arg : args) {
            if (arg.startsWith(prefix)) {
                return arg.substring(prefix.length());
            }
        }
        return defaultValue;
    }

    public static boolean parseBooleanArgument(String[] args, String shortPrefix, String longPrefix) {
        checkIllegalBooleanUsage(args, shortPrefix, longPrefix);
        final List<String> asList = Arrays.asList(args);
        return asList.contains("--"+longPrefix) || asList.contains("-"+shortPrefix);
    }

    private static void checkIllegalBooleanUsage(String[] args, String shortPrefix, String longPrefix) {
        for (String arg : args) {
            if (arg.equals("--"+shortPrefix) || arg.equals("-"+longPrefix)) {
                throw new IllegalArgumentException(String.format("%s is not supported. Use -%s or --%s instead", arg, shortPrefix, longPrefix));
            }
        }
    }
}
