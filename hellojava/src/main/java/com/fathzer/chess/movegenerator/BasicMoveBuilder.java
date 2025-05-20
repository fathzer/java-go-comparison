package com.fathzer.chess.movegenerator;

import java.util.List;

import com.fathzer.chess.Board;
import com.fathzer.chess.Move;
import com.fathzer.chess.Piece;

public abstract class BasicMoveBuilder implements MoveBuilder {
    private final int[] deltas;
    private final boolean isWhite;

    protected BasicMoveBuilder(int[] deltas, boolean isWhite) {
        this.deltas = deltas;
        this.isWhite = isWhite;
    }

    @Override
    public List<Move> build(List<Move> moves, Board board, int from) {
        for (int delta : deltas) {
            int to = from + delta;
            if (to<0 || to>63) {
                continue;
            }
            Piece piece = board.getPiece(to);
            if (piece == null || piece.isWhite() != isWhite) {
                moves.add(new Move(from, to));
            }
        }
        return moves;
    }
}
