package com.fathzer.chess.movegenerator;

import java.util.List;

import com.fathzer.chess.Board;
import com.fathzer.chess.BoardContent;
import com.fathzer.chess.Move;

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
            BoardContent piece = board.getContent(to);
            if (piece == null || piece.canBeCapturedBy(isWhite)) {
                moves.add(new Move(from, to));
            }
        }
        return moves;
    }
}
