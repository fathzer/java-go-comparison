package com.fathzer.chess.movegenerator;

import java.util.List;

/**
 * Generates moves for pawns.
 * <br>WARNING: This is a very basic implementation. It does not manage en passant, promotion.
 */
public class PawnMoveBuilder implements MoveBuilder {
    private final boolean isWhite;
    private final int advanceDelta;
    private final int captureDeltaWest;
    private final int captureDeltaEast;
    private final int twoAdvanceRank;
    
    public PawnMoveBuilder(boolean isWhite) {
        this.isWhite = isWhite;
        this.advanceDelta = isWhite ? Direction.NORTH.getDelta() : Direction.SOUTH.getDelta();
        this.captureDeltaWest = isWhite ? Direction.NORTH_WEST.getDelta() : Direction.SOUTH_WEST.getDelta();
        this.captureDeltaEast = isWhite ? Direction.NORTH_EAST.getDelta() : Direction.SOUTH_EAST.getDelta();
        this.twoAdvanceRank = isWhite ? 1 : 6;
    }
    
    @Override
    public <T> void build(List<T> moves, Explorable board, int from, MoveConstructor<T> moveBuilder) {
        int to = from + advanceDelta;
        if (board.getCapturable(to) == null) {
            moves.add(moveBuilder.create(from, to));
            to += advanceDelta;
            if (twoAdvanceRank == board.getRank(from) && board.getCapturable(to)==null) {
                moves.add(moveBuilder.create(from, to));
            }
        }
        to = from + captureDeltaWest;
        Capturable captured = board.getCapturable(to);
        if (captured != null && captured.canBeCapturedBy(isWhite)) {
            moves.add(moveBuilder.create(from, to));
        }
        to = from + captureDeltaEast;
        captured = board.getCapturable(to);
        if (captured != null && captured.canBeCapturedBy(isWhite)) {
            moves.add(moveBuilder.create(from, to));
        }
    }
}
