package com.fathzer.hellojava;

import static com.fathzer.pi.PiCalculator.computePi;

import java.math.BigDecimal;
import java.math.RoundingMode;
import java.util.concurrent.atomic.AtomicReference;
import java.util.stream.IntStream;

import com.fathzer.chess.Board;
import com.fathzer.chess.Perft;
import com.fathzer.chess.Perft.Result;
import com.fathzer.oop.ACoolThing;
import com.fathzer.oop.AbstractThing;
import com.fathzer.oop.CoolInterface;

public class Main {
    public static void main(String[] args) {
        try {
            int piLoops = Flags.parsePiLoops(args);
            int perftDepth = Flags.parsePerftDepth(args);

            System.out.println("Starting OOP concepts test...");
            oopConceptsTest();
            System.out.println("Starting concurrency test with " + piLoops + " loops...");
            concurrencyTest(piLoops);
            System.out.println("Starting Perft test...");
            perftTest(perftDepth);
        } catch (IllegalArgumentException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }

    private static void oopConceptsTest() {
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

    private static void perftTest(int depth) {
        for (int i = 1; i <= 5; i++) {
            doPerft(depth, new Board("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"), true);
            doPerft(depth, new Board("rnbqkbnr/pp1ppppp/2p5/8/6P1/2P5/PP1PPP1P/RNBQKBNR"), false);
        }
    }

    private static void doPerft(int depth, Board board, boolean whitePlaying) {
        final String DURATION_FORMAT = "duration (%d depth): %d ms";
        final String FOUND_FORMAT = "Found: %d leaf nodes. Generated: %d";
        final Perft perft = new Perft();
        long start = System.currentTimeMillis();
        final Result result = perft.perft(board, depth, whitePlaying);
        System.out.println(String.format(DURATION_FORMAT, depth, System.currentTimeMillis() - start));
        System.out.println(String.format(FOUND_FORMAT, result.leafNodesCount(), result.searchedNodesCount()));
    }
}
