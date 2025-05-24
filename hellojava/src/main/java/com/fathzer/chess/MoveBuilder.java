package com.fathzer.chess;

import java.util.List;

@FunctionalInterface
interface MoveBuilder {
    /**
     * Scans the board for legal moves from a given square.
     * @param moves the list of moves to add to
     * @param board the board to explore
     * @param from the square to scan from
     */
    void build(List<Move> moves, Board board, int from);
}
