package com.fathzer.hellojava;

public class Flags {
    private Flags() {}

    public static int parsePiLoops(String[] args) {
        int piLoops = 2000; // default value
        piLoops = parseIntegerArgument(args, "--piLoops=", piLoops);
        piLoops = parseIntegerArgument(args, "-pl=", piLoops);
        return piLoops;
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

    public static void printUsage() {
        System.out.println("Usage: java Main [--piLoops=<number>]");
        System.out.println("  --piLoops=<number>  Number of Pi calculations to perform (default: 2000)");
    }
}
