package com.fathzer.chess;

import java.util.HashMap;
import java.util.Map;

public enum Piece {
	BLOCKER('X', true),
	WHITE_PAWN('P', true),
	WHITE_KNIGHT('N', true),
	WHITE_BISHOP('B', true),
	WHITE_ROOK('R', true),
	WHITE_QUEEN('Q', true),
	WHITE_KING('K', true),
	BLACK_PAWN('p', false),
	BLACK_KNIGHT('n', false),
	BLACK_BISHOP('b', false),
	BLACK_ROOK('r', false),
	BLACK_QUEEN('q', false),
	BLACK_KING('k', false);

	private static final Map<Character, Piece> CODE_TO_PIECE = new HashMap<>();
	static {
		for (Piece piece : values()) {
			if (piece != BLOCKER) {
				CODE_TO_PIECE.put(piece.code, piece);
			}
		}
	}

	private final char code;
	private final boolean isWhite;

	private Piece(char code, boolean isWhite) {
		this.code = code;
		this.isWhite = isWhite;
	}

	public static Piece fromCode(char code) {
		return CODE_TO_PIECE.get(code);
	}

	public char getCode() {
		return code;
	}

	public boolean isWhite() {
		return isWhite;
	}

    public boolean canBeCapturedBy(boolean white) {
        return this!=Piece.BLOCKER && white != isWhite;
    }
}