package com.fathzer.chess;

import java.util.List;

@FunctionalInterface
interface MoveBuilder {
    @FunctionalInterface
    interface MoveConstructor<T> {
        T create(int from, int to);
    }

    @FunctionalInterface
    interface Capturable {
        boolean canBeCapturedBy(boolean white);
    }

    interface Explorable {
        Capturable getCapturable(int index);
        int getRank(int index);
    }

    /**
     * Scans the board for legal moves from a given square.
     * @param moves the list of moves to add to
     * @param board the board to explore
     * @param from the square to scan from
     */
    <T> void build(List<T> moves, Explorable board, int from, MoveConstructor<T> moveBuilder);
}
