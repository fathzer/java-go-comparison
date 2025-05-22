package com.fathzer.chess;

import java.util.Arrays;
import java.util.List;

final class MoveGenerators {
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
        MOVE_BUILDERS = new MoveBuilder[Piece.values().length];
        MOVE_BUILDERS[Piece.WHITE_KING.ordinal()] = WHITE_KING_MOVE_BUILDER;
        MOVE_BUILDERS[Piece.WHITE_QUEEN.ordinal()] = WHITE_QUEEN_MOVE_BUILDER;
        MOVE_BUILDERS[Piece.WHITE_ROOK.ordinal()] = WHITE_ROOK_MOVE_BUILDER;
        MOVE_BUILDERS[Piece.WHITE_BISHOP.ordinal()] = WHITE_BISHOP_MOVE_BUILDER;
        MOVE_BUILDERS[Piece.WHITE_KNIGHT.ordinal()] = WHITE_KNIGHT_MOVE_BUILDER;
        MOVE_BUILDERS[Piece.WHITE_PAWN.ordinal()] = WHITE_PAWN_MOVE_BUILDER;
        MOVE_BUILDERS[Piece.BLACK_KING.ordinal()] = BLACK_KING_MOVE_BUILDER;
        MOVE_BUILDERS[Piece.BLACK_QUEEN.ordinal()] = BLACK_QUEEN_MOVE_BUILDER;
        MOVE_BUILDERS[Piece.BLACK_ROOK.ordinal()] = BLACK_ROOK_MOVE_BUILDER;
        MOVE_BUILDERS[Piece.BLACK_BISHOP.ordinal()] = BLACK_BISHOP_MOVE_BUILDER;
        MOVE_BUILDERS[Piece.BLACK_KNIGHT.ordinal()] = BLACK_KNIGHT_MOVE_BUILDER;
        MOVE_BUILDERS[Piece.BLACK_PAWN.ordinal()] = BLACK_PAWN_MOVE_BUILDER;
    }

    private MoveGenerators() {
    }

    public static MoveBuilder get(Piece piece) {
        return MOVE_BUILDERS[piece.ordinal()];
    }
    
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
    
    private abstract static class BasicMoveBuilder implements MoveBuilder {
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

    private static class SliderMoveBuilder implements MoveBuilder {
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
}
