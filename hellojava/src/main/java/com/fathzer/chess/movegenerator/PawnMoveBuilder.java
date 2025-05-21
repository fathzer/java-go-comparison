package com.fathzer.chess.movegenerator;

import java.util.List;

import com.fathzer.chess.Move;
import com.fathzer.chess.Board;
import com.fathzer.chess.BoardContent;
import com.fathzer.chess.Direction;

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
    public List<Move> build(List<Move> moves, Board board, int from) {
        int to = from + advanceDelta;
        if (board.getContent(to) == null) {
            moves.add(new Move(from, to));
            to += advanceDelta;
            if (twoAdvanceRank == Board.getRank(from) && board.getContent(to)==null) {
                moves.add(new Move(from, to));
            }
        }
        to = from + captureDeltaWest;
        BoardContent captured = board.getContent(to);
        if (captured != null && captured.canBeCapturedBy(isWhite)) {
            moves.add(new Move(from, to));
        }
        to = from + captureDeltaEast;
        captured = board.getContent(to);
        if (captured != null && captured.canBeCapturedBy(isWhite)) {
            moves.add(new Move(from, to));
        }
        return moves;
    }
}
