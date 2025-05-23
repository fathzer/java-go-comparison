package com.fathzer.hellojava;

public class Flags {
    private Flags() {}

    public static int parsePiLoops(String[] args) {
        int piLoops = 2000; // default value
        piLoops = parseIntegerArgument(args, "--piLoops=", piLoops);
        piLoops = parseIntegerArgument(args, "-pl=", piLoops);
        return piLoops;
    }

    public static int parsePerftDepth(String[] args) {
        int perftDepth = 5; // default value
        perftDepth = parseIntegerArgument(args, "--perftDepth=", perftDepth);
        perftDepth = parseIntegerArgument(args, "-pd=", perftDepth);
        return perftDepth;
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
        System.out.println("Usage: java Main [--piLoops=<number>] [--perftDepth=<number>]");
        System.out.println("  --piLoops=<number>  Number of Pi calculations to perform (default: 2000)");
        System.out.println("  --perftDepth=<number>  Depth of Perft (Performance Test) to perform (default: 5)");
    }
}
