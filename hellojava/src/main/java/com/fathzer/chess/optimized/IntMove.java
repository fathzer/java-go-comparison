package com.fathzer.chess.optimized;

/**
 * Utility class for encoding and decoding move information into a single int.
 * The encoding uses the following bit layout:
 * - from (7 bits): bits 0-6 (0-127)
 * - to (7 bits): bits 7-13 (0-127)
 * - capture (5 bits): bits 14-18 (0-31)
 */
public final class IntMove {
    private static final int FROM_MASK = 0x7F;        // 7 bits for from (0-127)
    private static final int TO_MASK = 0x3F80;         // 7 bits for to (shifted left by 7)
    private static final int CAPTURE_MASK = 0x7C000;    // 5 bits for capture (shifted left by 14)
    
    private static final int TO_SHIFT = 7;
    private static final int CAPTURE_SHIFT = 14;

    private IntMove() {
        // Private constructor to prevent instantiation
    }

    public static int fromUCI(Board board, String uci) {
        if (uci.length() != 4) {
            throw new IllegalArgumentException("Invalid UCI move: " + uci);
        }
        final int toSquare = Board.getSquare(uci.substring(2));
        final int capturedPiece = board.getPiece(toSquare);
        return move(Board.getSquare(uci.substring(0, 2)), toSquare, capturedPiece);
    }

    
    /**
     * Creates a packed move integer from individual components.
     * @param from The source square (0-127)
     * @param to The target square (0-127)
     * @param capture The capture value (0-31)
     * @return A packed integer containing all three values
     */
    public static int move(int from, int to, int capture) {
        return (from) | 
               (to << TO_SHIFT) | 
               (capture << CAPTURE_SHIFT);
    }
    
    /**
     * Extracts the source square from a packed move.
     * 
     * @param move The packed move
     * @return The source square (0-127)
     */
    public static int from(int move) {
        return move & FROM_MASK;
    }
    
    /**
     * Extracts the target square from a packed move.
     * 
     * @param move The packed move
     * @return The target square (0-127)
     */
    public static int to(int move) {
        return (move & TO_MASK) >>> TO_SHIFT;
    }
    
    /**
     * Extracts the capture value from a packed move.
     * 
     * @param move The packed move
     * @return The capture value (0-31)
     */
    public static int capture(int move) {
        return (move & CAPTURE_MASK) >>> CAPTURE_SHIFT;
    }
    
    public static String toString(int move) {
    	final int capture = capture(move);
        if (capture==0) {
            return Board.getUCI(from(move)) + Board.getUCI(to(move));
        } else {
            return Board.getUCI(from(move)) + "x" + Board.getUCI(to(move)) + "(" + Piece.getCode(capture) + ")";
        }
    }
}
