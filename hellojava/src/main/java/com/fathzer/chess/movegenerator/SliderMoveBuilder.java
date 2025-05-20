package com.fathzer.chess.movegenerator;

import java.util.Arrays;
import java.util.List;

import com.fathzer.chess.Board;
import com.fathzer.chess.Direction;
import com.fathzer.chess.Move;
import com.fathzer.chess.Piece;

public final class SliderMoveBuilder implements MoveBuilder {
    private final int[] deltas;
    private final boolean isWhite;

    public SliderMoveBuilder(Direction[] deltas, boolean isWhite) {
        this.deltas = Arrays.stream(deltas).mapToInt(Direction::getDelta).toArray();
        this.isWhite = isWhite;
    }

    public void scanDirection(List<Move> moves, Board board, int from, int delta) {
        int to = from + delta;
        while (to >= 0 && to < 64) {
            Piece piece = board.getPiece(to);
            if (piece == null) {
                moves.add(new Move(from, to));
            } else {
                if (piece.isWhite() != isWhite) {
                    moves.add(new Move(from, to));
                }
                break;
            }
            to += delta;
        }
    }

    @Override
    public List<Move> build(List<Move> moves, Board board, int from) {
        for (int delta : deltas) {
            scanDirection(moves, board, from, delta);
        }
        return moves;
    }
}
