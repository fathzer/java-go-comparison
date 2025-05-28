package com.fathzer.hellojava;

import com.fathzer.chess.common.PerftResult;

@FunctionalInterface
interface PerftRun {
    record PerfTResultWithDuration(PerftResult result, long duration) {}
    static final com.fathzer.chess.Perft perft = new com.fathzer.chess.Perft();
    static final com.fathzer.chess.optimized.Perft optimizedPerft = new com.fathzer.chess.optimized.Perft();
    PerftResult run(String fen, boolean whitePlaying, int depth);
    private static PerftRun of(boolean optimized) {
        if (optimized) {
            return (fen, whitePlaying, depth) -> optimizedPerft.perft(new com.fathzer.chess.optimized.Board(fen), depth, whitePlaying);
        } else {
            return (fen, whitePlaying, depth) -> perft.perft(new com.fathzer.chess.Board(fen), depth, whitePlaying);
        }
    }
    static PerfTResultWithDuration run(String fen, boolean whitePlaying, int depth, boolean fast) {
        final PerftRun run = of(fast);
        final long start = System.currentTimeMillis();
        final PerftResult result = run.run(fen, whitePlaying, depth);
        return new PerfTResultWithDuration(result, System.currentTimeMillis() - start);
    }
}