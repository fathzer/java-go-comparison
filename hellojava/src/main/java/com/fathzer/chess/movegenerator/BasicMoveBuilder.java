package com.fathzer.chess.movegenerator;

import java.util.List;

public abstract class BasicMoveBuilder implements MoveBuilder {
    private final int[] deltas;
    private final boolean isWhite;

    protected BasicMoveBuilder(int[] deltas, boolean isWhite) {
        this.deltas = deltas;
        this.isWhite = isWhite;
    }

    @Override
    public <T> void build(List<T> moves, Explorable board, int from, MoveConstructor<T> moveBuilder) {
        for (int delta : deltas) {
            int to = from + delta;
            Capturable piece = board.getCapturable(to);
            if (piece == null || piece.canBeCapturedBy(isWhite)) {
                moves.add(moveBuilder.create(from, to));
            }
        }
    }
}
