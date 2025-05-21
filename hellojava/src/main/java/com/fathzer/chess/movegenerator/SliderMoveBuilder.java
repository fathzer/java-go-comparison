package com.fathzer.chess.movegenerator;

import java.util.Arrays;
import java.util.List;

public final class SliderMoveBuilder implements MoveBuilder {
    private final int[] deltas;
    private final boolean isWhite;

    public SliderMoveBuilder(Direction[] deltas, boolean isWhite) {
        this.deltas = Arrays.stream(deltas).mapToInt(Direction::getDelta).toArray();
        this.isWhite = isWhite;
    }

    public <T> void scanDirection(List<T> moves, Explorable board, int from, int delta, MoveConstructor<T> moveBuilder) {
        int to = from + delta;
        while (true) {
            Capturable piece = board.getCapturable(to);
            if (piece == null) {
                moves.add(moveBuilder.create(from, to));
            } else {
                if (piece.canBeCapturedBy(isWhite)) {
                    moves.add(moveBuilder.create(from, to));
                }
                break;
            }
            to += delta;
        }
    }

    @Override
    public <T> void build(List<T> moves, Explorable board, int from, MoveConstructor<T> moveBuilder) {
        for (int delta : deltas) {
            scanDirection(moves, board, from, delta, moveBuilder);
        }
    }
}
