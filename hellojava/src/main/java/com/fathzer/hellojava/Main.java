package com.fathzer.hellojava;

import static com.fathzer.pi.PiCalculator.computePi;

import java.math.BigDecimal;
import java.math.RoundingMode;
import java.util.concurrent.atomic.AtomicReference;
import java.util.stream.IntStream;

import com.fathzer.chess.common.PerftResult;
import com.fathzer.oop.ACoolThing;
import com.fathzer.oop.AbstractThing;
import com.fathzer.oop.CoolInterface;

public class Main {
    public static void main(String[] args) {
        try {
            int piLoops = Flags.parsePiLoops(args);
            int perftDepth = Flags.parsePerftDepth(args);

            oopConceptsTest();
            concurrencyTest(piLoops);
            perftTest(perftDepth);
        } catch (IllegalArgumentException e) {
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

    private static void perftTest(int depth) {
        System.out.println("Starting Perft test...");
        perftTest(depth, false);
        System.out.println("Starting optimized perft tests");
        perftTest(depth, true);
    }

    private static final String PERFT_FORMAT = "Found: %d leaf nodes with %d move generation at depth %d in %d ms for %s";
    
    private static void perftTest(int depth, boolean fast) {
        for (int i = 1; i <= 3; i++) {
            final String start = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR";
            final String other = "rnbqkbnr/pp1ppppp/2p5/8/6P1/2P5/PP1PPP1P/RNBQKBNR";
            if (fast) {
				doFastPerft(start, true, depth);
				doFastPerft(other, false, depth);
            } else {
				doPerft(start, true, depth);
				doPerft(other, false, depth);
            }
        }
    }
    
    private static void doPerft(String fen, boolean whitePlaying, int depth) {
        long start = System.currentTimeMillis();
        final com.fathzer.chess.Perft perft = new com.fathzer.chess.Perft();
        com.fathzer.chess.Board board = new com.fathzer.chess.Board(fen);
        final PerftResult result = perft.perft(board, depth, whitePlaying);
        System.out.println(String.format(PERFT_FORMAT, result.leafNodesCount(), result.searchedNodesCount(), depth, (System.currentTimeMillis() - start), fen));
    }
    
    
    private static void doFastPerft(String fen, boolean whitePlaying, int depth) {
        long start = System.currentTimeMillis();
        final com.fathzer.chess.optimized.Perft perft = new com.fathzer.chess.optimized.Perft();
        com.fathzer.chess.optimized.Board board = new com.fathzer.chess.optimized.Board(fen);
        final PerftResult result = perft.perft(board, depth, whitePlaying);
        System.out.println(String.format(PERFT_FORMAT, result.leafNodesCount(), result.searchedNodesCount(), depth, (System.currentTimeMillis() - start), fen));
    }
}
