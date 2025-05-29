package com.fathzer.chess.optimized;

@FunctionalInterface
interface MoveBuilder {
    /**
     * Scans the board for legal moves from a given square.
     * @param moves the list of moves to add to
     * @param board the board to explore
     * @param from the square to scan from
     */
    void build(IntList moves, Board board, int from);
}
