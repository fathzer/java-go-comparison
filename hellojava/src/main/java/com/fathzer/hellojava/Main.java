package com.fathzer.hellojava;

import static com.fathzer.hellojava.PerftRun.run;
import static com.fathzer.pi.PiCalculator.computePi;

import java.math.BigDecimal;
import java.math.RoundingMode;
import java.util.concurrent.atomic.AtomicReference;
import java.util.stream.IntStream;

import com.fathzer.hellojava.PerftRun.PerfTResultWithDuration;
import com.fathzer.oop.ACoolThing;
import com.fathzer.oop.AbstractThing;
import com.fathzer.oop.CoolInterface;

public class Main {
    public static void main(String[] args) {
        try {
            final int piLoops = Flags.parseIntegerArgument(args, "pl", "piLoops", 2000);
            final int perftDepth = Flags.parseIntegerArgument(args, "pd", "perftDepth", 6);
            final String fen = Flags.parseStringArgument(args, "f", "fen", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR");
            final boolean blackPlaying = Flags.parseBooleanArgument(args, "b", "blackPlaying");
            final int perftLoops = Flags.parseIntegerArgument(args, "ptl", "perftLoops", 5);
            final boolean optimizedPerft = Flags.parseBooleanArgument(args, "o", "optimizedPerft");
            final int coolingTime = Flags.parseIntegerArgument(args, "pct", "perftCoolingTime", 0);

            oopConceptsTest();
            concurrencyTest(piLoops);
            if (perftLoops > 0) {
                perftTest(fen, !blackPlaying, perftDepth, perftLoops, coolingTime, optimizedPerft);
            }
        } catch (IllegalArgumentException e) {
            printUsage();
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }

    private static void oopConceptsTest() {
        System.out.println("Starting OOP concepts test...");
        AbstractThing aThing = new ACoolThing("my abstract thing");
        ((CoolInterface)aThing).doSomethingCool();
        aThing.doSomething();

        CoolInterface iThing = new ACoolThing("my cool interface");
        ((AbstractThing)iThing).doSomething();
        iThing.doSomethingCool();

        ACoolThing acoolThing = new ACoolThing("my cool interface");
        acoolThing.doSomething();
        acoolThing.doSomethingCool();
    }

    private static void concurrencyTest(int nbLoops) {
        System.out.println("Starting concurrency test with " + nbLoops + " loops...");
        if (nbLoops < 1) {
            throw new IllegalArgumentException("Number of loops must be at least 1");
        }
        final AtomicReference<BigDecimal> result = new AtomicReference<>(new BigDecimal(0));
        long start = System.currentTimeMillis();
        IntStream.range(0, nbLoops).parallel().forEach(i -> {
            var pi = computePi(1000+i);
            synchronized(result) {
                result.set(result.get().add(pi));
            }
        });
        System.out.println("duration ("+nbLoops+" loops): " + (System.currentTimeMillis() - start)+" ms");
        System.out.println("result: " + result.get().divide(BigDecimal.valueOf(nbLoops), 10, RoundingMode.HALF_UP));
    }

    private static void perftTest(String fen, boolean whitePlaying, int depth, int nbLoops, long coolingTime, boolean optimized) {
        System.out.println(String.format("Starting %sPerft test with %dms cooling time before every perft computation ...", optimized?"optimized ":"", coolingTime));
        silentSleep(coolingTime);
        final PerfTResultWithDuration result = run(fen, whitePlaying, depth, optimized);
        System.out.println(String.format("First execution found: %d leaf nodes with %d move generation at depth %d in %d ms for %s", result.result().leafNodesCount(), result.result().searchedNodesCount(), depth, result.duration(), fen));
        if (nbLoops <= 1) return;
        long duration = 0;
        nbLoops--;
        for (int i = 1; i <= nbLoops; i++) {
            System.out.print(".");
            silentSleep(coolingTime);
            final PerfTResultWithDuration loopResult = run(fen, whitePlaying, depth, optimized);
            if (!loopResult.result().equals(result.result())) {
                System.out.println("\nResult mismatch at loop " + i);
                return;
            }
            duration += loopResult.duration();
        }
        System.out.println(String.format("\nduration (%d loops): %d ms (average: %d ms)", nbLoops, duration, duration/nbLoops));
   }

    private static void silentSleep(long pauseDuration) {
        try {
            Thread.sleep(pauseDuration);
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            return;
        }
    }

    private static void printUsage() {
        System.out.println("Usage: java Main [--piLoops=<number>] [--perftDepth=<number>] [--fen=<fen>] [--blackPlaying]");
        System.out.println("  --piLoops=<number>  Number of Pi calculations to perform (default: 2000)");
        System.out.println("  --optimizedPerft  Use optimized Perft (Performance Test) (default: false)");
        System.out.println("  --perftDepth=<number>  Depth of Perft (Performance Test) to perform (default: 5)");
        System.out.println("  --perftLoops=<number>  Number of Perft (Performance Test) to perform (default: 3)");
        System.out.println("  --fen=<fen>  FEN position to use for Perft (default: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR)");
        System.out.println("  --blackPlaying  If the black player is playing (default: false)");
        System.out.println("  --perftCoolingTime=<number>  Cooling time between Perft (Performance Test) computations (default: 1000)");
    }
}
