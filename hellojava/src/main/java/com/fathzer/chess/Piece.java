package com.fathzer.chess;

import java.util.HashMap;
import java.util.Map;

import com.fathzer.chess.movegenerator.MoveBuilder;
import com.fathzer.chess.movegenerator.PawnMoveBuilder;
import com.fathzer.chess.movegenerator.SliderMoveBuilder;
import com.fathzer.chess.movegenerator.KingMoveBuilder;
import com.fathzer.chess.movegenerator.KnightMoveBuilder;

public enum Piece {
	BLOCKER('X', true, (l, b, f) -> {}),
	WHITE_PAWN('P', true, new PawnMoveBuilder(true)),
	WHITE_KNIGHT('N', true, new KnightMoveBuilder(true)),
	WHITE_BISHOP('B', true, new SliderMoveBuilder(new Direction[] { Direction.NORTH_EAST, Direction.NORTH_WEST, Direction.SOUTH_EAST, Direction.SOUTH_WEST }, true)),
	WHITE_ROOK('R', true, new SliderMoveBuilder(new Direction[] { Direction.NORTH, Direction.SOUTH, Direction.EAST, Direction.WEST }, true)),
	WHITE_QUEEN('Q', true, new SliderMoveBuilder(new Direction[] { Direction.NORTH, Direction.SOUTH, Direction.EAST, Direction.WEST, Direction.NORTH_EAST, Direction.NORTH_WEST, Direction.SOUTH_EAST, Direction.SOUTH_WEST }, true)),
	WHITE_KING('K', true, new KingMoveBuilder(true)),
	BLACK_PAWN('p', false, new PawnMoveBuilder(false)),
	BLACK_KNIGHT('n', false, new KnightMoveBuilder(false)),
	BLACK_BISHOP('b', false, new SliderMoveBuilder(new Direction[] { Direction.NORTH_EAST, Direction.NORTH_WEST, Direction.SOUTH_EAST, Direction.SOUTH_WEST }, false)),
	BLACK_ROOK('r', false, new SliderMoveBuilder(new Direction[] { Direction.NORTH, Direction.SOUTH, Direction.EAST, Direction.WEST }, false)),
	BLACK_QUEEN('q', false, new SliderMoveBuilder(new Direction[] { Direction.NORTH, Direction.SOUTH, Direction.EAST, Direction.WEST, Direction.NORTH_EAST, Direction.NORTH_WEST, Direction.SOUTH_EAST, Direction.SOUTH_WEST }, false)),
	BLACK_KING('k', false, new KingMoveBuilder(false));

	private static final Map<Character, Piece> CODE_TO_PIECE = new HashMap<>();
	static {
		for (Piece piece : values()) {
			CODE_TO_PIECE.put(piece.code, piece);
		}
	}

	private final char code;
	private final boolean isWhite;
	private final MoveBuilder moveBuilder;

	private Piece(char code, boolean isWhite, MoveBuilder scanner) {
		this.code = code;
		this.isWhite = isWhite;
		this.moveBuilder = scanner;
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

	public MoveBuilder getMoveBuilder() {
		return moveBuilder;
	}

    public boolean canBeCapturedBy(boolean white) {
        return this!=Piece.BLOCKER && white != isWhite;
    }
}