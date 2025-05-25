package com.fathzer.chess.optimized;

import java.util.HashMap;
import java.util.Map;
import java.util.stream.IntStream;

public final class Piece {
	public static final int BLOCKER = Integer.MAX_VALUE;
	public static final int NONE = 0;
	public static final int WHITE_PAWN = 1;
	public static final int WHITE_KNIGHT = 2;
	public static final int WHITE_BISHOP = 3;
	public static final int WHITE_ROOK = 4;
	public static final int WHITE_QUEEN = 5;
	public static final int WHITE_KING = 6;
	public static final int BLACK_PAWN = 7;
	public static final int BLACK_KNIGHT = 8;
	public static final int BLACK_BISHOP = 9;
	public static final int BLACK_ROOK = 10;
	public static final int BLACK_QUEEN = 11;
	public static final int BLACK_KING = 12;
	
	static final char[] CODES = new char[] {' ', 'P', 'N', 'B', 'R', 'Q', 'K', 'p', 'n', 'b', 'r', 'q', 'k' };

	private static final Map<Character, Integer> CODE_TO_PIECE = new HashMap<>();
	static {
		IntStream.rangeClosed(WHITE_PAWN, BLACK_KING).forEach(piece -> CODE_TO_PIECE.put(CODES[piece], piece));
	}

	public static int fromCode(char code) {
		Integer value = CODE_TO_PIECE.get(code);
		return value==null?NONE:value;
	}

	public static char getCode(int piece) {
		return CODES[piece];
	}

	public static boolean isWhite(int piece) {
		return piece <= WHITE_KING;
	}

    public static boolean canBeCapturedBy(int piece, boolean white) {
        return piece!=Piece.BLOCKER && white != isWhite(piece);
    }

	private Piece() {
	}
}