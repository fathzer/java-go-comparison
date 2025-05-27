package com.fathzer.chess.optimized;

import static com.fathzer.chess.optimized.IntMove.*;
import static com.fathzer.chess.optimized.Piece.*;

import java.util.Arrays;

/** A tiny chess board
 * WARNING: This is a very basic implementation. It is not a chess move generator.
 * It lacks important features like king safety, en passant, castling, promotion.
 */
public class Board {
    private final int[] pieces;
	public Board(Board copy) {
		this.pieces = copy.pieces.clone();
	}

    public Board(String fen) {
        this.pieces = new int[120];
        fillBlockers();
		int rank = 7;
		int file = 0;
		for (char c : fen.toCharArray()) {
			if (Character.isDigit(c)) {
				final int count = Character.getNumericValue(c);
				if (count>8-file) {
					throw new IllegalArgumentException("Invalid FEN: too many pieces on rank " + rank);
				}
				file += count;
			} else if (c == '/') {
				if (file!=8) {
					throw new IllegalArgumentException("Invalid FEN: missing files on rank " + rank);
				}
				if (rank == 0) {
					throw new IllegalArgumentException("Invalid FEN: too many ranks");
				}
				file = 0;
				rank--;
			} else {
				final int piece = Piece.fromCode(c);
				if (piece == Piece.NONE) {
					throw new IllegalArgumentException("Invalid FEN: unknown piece " + c);
				}
				this.pieces[21+rank*10+file] = piece;
				file++;
			}
		}
		if (file!=8) {
			throw new IllegalArgumentException("Invalid FEN: missing files on rank " + rank);
		}
		if (rank!=0) {
			throw new IllegalArgumentException("Invalid FEN: missing ranks");
		}
    }
    
    private void fillBlockers() {
    	Arrays.fill(pieces, 0, 20, BLOCKER);
    	Arrays.fill(pieces, 100, 120, BLOCKER);
    	for (int i=2; i<10; i++) {
			final int startRank = i*10;
    		pieces[startRank] = BLOCKER;
    		pieces[startRank+9] = BLOCKER;
    	}
    }
	
	static int getSquare(String uciSquare) {
		if (uciSquare.length() != 2) {
			throw new IllegalArgumentException("Invalid UCI square: " + uciSquare);
		}
		final char file = uciSquare.charAt(0);
		final char rank = uciSquare.charAt(1);
		if (rank < '1' || rank > '8' || file < 'a' || file > 'h') {
			throw new IllegalArgumentException("Invalid UCI square: " + uciSquare);
		}
		return 21+ 10*(Integer.parseInt(uciSquare.substring(1)) - 1) + (file - 'a');
	}
	
	static int getRank(int square) {
		return (square-21)/10;
	}
	
	static String getUCI(int square) {
		square -= 21;
		return (char)(square%10 + 'a') + "" +(square/10 + 1);
	}

	public int getPiece(String uciSquare) {
		return pieces[getSquare(uciSquare)];
	}

	int getPiece(int square) {
		return pieces[square];
	}

	public IntList getMoves(IntList moves, boolean white) {
		moves.clear();
		for (int square = 20; square < 100; square++) {
			int piece = getPiece(square);
			if (piece != NONE && piece!=BLOCKER && isWhite(piece)==white) {
				MoveGenerators.get(piece).build(moves, this, square);
			}
		}
		return moves;
	}

	public void makeMove(int move) {
	    int from = from(move);
	    int to = to(move);
	    
	    // Move the piece
	    pieces[to] = pieces[from];
	    pieces[from] = NONE; // Empty the source square
	}

	public void unmakeMove(int move) {
	    int from = from(move);
	    int to = to(move);
	    
	    // Restore the moved piece
	    pieces[from] = pieces[to];
	    
	    // Restore the captured piece (if any)
	    pieces[to] = capture(move);
	}

	@Override
	public String toString() {
		final StringBuilder b = new StringBuilder();
		for (char rank = '8'; rank >= '1'; rank--) {
			b.append(rank);
			b.append(' ');
			for (char file = 'a'; file <= 'h'; file++) {
				int piece = getPiece(file+""+rank);
				b.append(getCode(piece));
			}
			b.append('\n');
		}
		b.append("  ");
		for (char file = 'a'; file < 'h' ; file++) {
			b.append(file);
		}
		return b.toString();
	}
}
