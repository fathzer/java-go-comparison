package com.fathzer.chess.optimized;

import static com.fathzer.chess.optimized.Piece.*;

import java.util.Arrays;

final class MoveGenerators {
    private enum Direction {
        NORTH(ROW_WIDTH),
        SOUTH(-ROW_WIDTH),
        EAST(+1),
        WEST(-1),
        NORTH_EAST(ROW_WIDTH+1),
        NORTH_WEST(ROW_WIDTH-1),
        SOUTH_EAST(-ROW_WIDTH+1),
        SOUTH_WEST(-ROW_WIDTH-1);

        private final int delta;

        private Direction(int delta) {
            this.delta = delta;
        }
    
        public int getDelta() {
            return delta;
        }
    }

    private static final int ROW_WIDTH = 10;

    private static final MoveBuilder WHITE_KING_MOVE_BUILDER = new KingMoveBuilder(true);
    private static final MoveBuilder WHITE_QUEEN_MOVE_BUILDER = new SliderMoveBuilder(new Direction[] { Direction.NORTH, Direction.SOUTH, Direction.EAST, Direction.WEST, Direction.NORTH_EAST, Direction.NORTH_WEST, Direction.SOUTH_EAST, Direction.SOUTH_WEST }, true);
    private static final MoveBuilder WHITE_ROOK_MOVE_BUILDER = new SliderMoveBuilder(new Direction[] { Direction.NORTH, Direction.SOUTH, Direction.EAST, Direction.WEST }, true);
    private static final MoveBuilder WHITE_BISHOP_MOVE_BUILDER = new SliderMoveBuilder(new Direction[] { Direction.NORTH_EAST, Direction.NORTH_WEST, Direction.SOUTH_EAST, Direction.SOUTH_WEST }, true);
    private static final MoveBuilder WHITE_KNIGHT_MOVE_BUILDER = new KnightMoveBuilder(true);
    private static final MoveBuilder WHITE_PAWN_MOVE_BUILDER = new PawnMoveBuilder(true);

    private static final MoveBuilder BLACK_KING_MOVE_BUILDER = new KingMoveBuilder(false);
    private static final MoveBuilder BLACK_QUEEN_MOVE_BUILDER = new SliderMoveBuilder(new Direction[] { Direction.NORTH, Direction.SOUTH, Direction.EAST, Direction.WEST, Direction.NORTH_EAST, Direction.NORTH_WEST, Direction.SOUTH_EAST, Direction.SOUTH_WEST }, false);
    private static final MoveBuilder BLACK_ROOK_MOVE_BUILDER = new SliderMoveBuilder(new Direction[] { Direction.NORTH, Direction.SOUTH, Direction.EAST, Direction.WEST }, false);
    private static final MoveBuilder BLACK_BISHOP_MOVE_BUILDER = new SliderMoveBuilder(new Direction[] { Direction.NORTH_EAST, Direction.NORTH_WEST, Direction.SOUTH_EAST, Direction.SOUTH_WEST }, false);
    private static final MoveBuilder BLACK_KNIGHT_MOVE_BUILDER = new KnightMoveBuilder(false);
    private static final MoveBuilder BLACK_PAWN_MOVE_BUILDER = new PawnMoveBuilder(false);

    private static final MoveBuilder[] MOVE_BUILDERS;

    static {
        MOVE_BUILDERS = new MoveBuilder[BLACK_KING+1];
        MOVE_BUILDERS[WHITE_KING] = WHITE_KING_MOVE_BUILDER;
        MOVE_BUILDERS[WHITE_QUEEN] = WHITE_QUEEN_MOVE_BUILDER;
        MOVE_BUILDERS[WHITE_ROOK] = WHITE_ROOK_MOVE_BUILDER;
        MOVE_BUILDERS[WHITE_BISHOP] = WHITE_BISHOP_MOVE_BUILDER;
        MOVE_BUILDERS[WHITE_KNIGHT] = WHITE_KNIGHT_MOVE_BUILDER;
        MOVE_BUILDERS[WHITE_PAWN] = WHITE_PAWN_MOVE_BUILDER;
        MOVE_BUILDERS[BLACK_KING] = BLACK_KING_MOVE_BUILDER;
        MOVE_BUILDERS[BLACK_QUEEN] = BLACK_QUEEN_MOVE_BUILDER;
        MOVE_BUILDERS[BLACK_ROOK] = BLACK_ROOK_MOVE_BUILDER;
        MOVE_BUILDERS[BLACK_BISHOP] = BLACK_BISHOP_MOVE_BUILDER;
        MOVE_BUILDERS[BLACK_KNIGHT] = BLACK_KNIGHT_MOVE_BUILDER;
        MOVE_BUILDERS[BLACK_PAWN] = BLACK_PAWN_MOVE_BUILDER;
    }

    private MoveGenerators() {
    }

    public static MoveBuilder get(int piece) {
        return MOVE_BUILDERS[piece];
    }
    
    private abstract static class BasicMoveBuilder implements MoveBuilder {
        private final int[] deltas;
        private final boolean isWhite;

        protected BasicMoveBuilder(int[] deltas, boolean isWhite) {
            this.deltas = deltas;
            this.isWhite = isWhite;
        }

        @Override
        public void build(IntList moves, Board board, int from) {
            for (int delta : deltas) {
                int to = from + delta;
                int piece = board.getPiece(to);
                if (piece == NONE || canBeCapturedBy(piece, isWhite)) {
                    moves.add(IntMove.move(from, to, piece));
                }
            }
        }
    }

    private static class KingMoveBuilder extends BasicMoveBuilder {
        public KingMoveBuilder(boolean isWhite) {
            super(new int[] { Direction.NORTH.getDelta(), Direction.SOUTH.getDelta(), Direction.EAST.getDelta(), Direction.WEST.getDelta(), Direction.NORTH_EAST.getDelta(), Direction.NORTH_WEST.getDelta(), Direction.SOUTH_EAST.getDelta(), Direction.SOUTH_WEST.getDelta() }, isWhite);
        }
    }

    private static class KnightMoveBuilder extends BasicMoveBuilder {
        public KnightMoveBuilder(boolean isWhite) {
            super(getDeltas(), isWhite);
        }
    
        private static int[] getDeltas() {
            final int[] deltas = new int[8];
            deltas[0] = 2*Direction.NORTH.getDelta()+Direction.EAST.getDelta();
            deltas[1] = 2*Direction.NORTH.getDelta()+Direction.WEST.getDelta();
            deltas[2] = 2*Direction.SOUTH.getDelta()+Direction.EAST.getDelta();
            deltas[3] = 2*Direction.SOUTH.getDelta()+Direction.WEST.getDelta();
            deltas[4] = Direction.NORTH.getDelta()+2*Direction.EAST.getDelta();
            deltas[5] = Direction.NORTH.getDelta()+2*Direction.WEST.getDelta();
            deltas[6] = Direction.SOUTH.getDelta()+2*Direction.EAST.getDelta();
            deltas[7] = Direction.SOUTH.getDelta()+2*Direction.WEST.getDelta();
            return deltas;
        }
    }

    /**
     * Generates moves for pawns.
     * <br>WARNING: This is a very basic implementation. It does not manage en passant, promotion.
     */
    private static class PawnMoveBuilder implements MoveBuilder {
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
        public void build(IntList moves, Board board, int from) {
            int to = from + advanceDelta;
            if (board.getPiece(to) == NONE) {
                moves.add(IntMove.move(from, to, NONE));
                to += advanceDelta;
                if (twoAdvanceRank == Board.getRank(from) && board.getPiece(to)==NONE) {
                    moves.add(IntMove.move(from, to, NONE));
                }
            }
            to = from + captureDeltaWest;
            int captured = board.getPiece(to);
            if (captured != NONE && canBeCapturedBy(captured, isWhite)) {
                moves.add(IntMove.move(from, to, captured));
            }
            to = from + captureDeltaEast;
            captured = board.getPiece(to);
            if (captured != NONE && canBeCapturedBy(captured, isWhite)) {
                moves.add(IntMove.move(from, to, captured));
            }
        }
    }

    private static class SliderMoveBuilder implements MoveBuilder {
        private final int[] deltas;
        private final boolean isWhite;

        public SliderMoveBuilder(Direction[] deltas, boolean isWhite) {
            this.deltas = Arrays.stream(deltas).mapToInt(Direction::getDelta).toArray();
            this.isWhite = isWhite;
        }

        public void scanDirection(IntList moves, Board board, int from, int delta) {
            int to = from + delta;
            while (true) {
                int piece = board.getPiece(to);
                if (piece == NONE) {
                    moves.add(IntMove.move(from, to, NONE));
                } else {
                    if (canBeCapturedBy(piece, isWhite)) {
                        moves.add(IntMove.move(from, to, piece));
                    }
                    break;
                }
                to += delta;
            }
        }

        @Override
        public void build(IntList moves, Board board, int from) {
            for (int delta : deltas) {
                scanDirection(moves, board, from, delta);
            }
        }
    }
}
