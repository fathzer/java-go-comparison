package com.fathzer.chess.movegenerator;

import java.util.List;

@FunctionalInterface
public interface MoveBuilder {
    /**
     * Scans the board for legal moves from a given square.
     * @param moves the list of moves to add to
     * @param board the board to explore
     * @param from the square to scan from
     */
    <T> void build(List<T> moves, Explorable board, int from, MoveConstructor<T> moveBuilder);
}
