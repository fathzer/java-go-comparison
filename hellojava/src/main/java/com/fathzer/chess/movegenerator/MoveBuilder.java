package com.fathzer.chess.movegenerator;

import java.util.List;

import com.fathzer.chess.Board;
import com.fathzer.chess.Move;

public interface MoveBuilder {
    /**
     * Scans the board for legal moves from a given square.
     * @param moves the list of moves to add to
     * @param board the board to scan
     * @param from the square to scan from
     */
    List<Move> build(List<Move> moves, Board board, int from);
}
