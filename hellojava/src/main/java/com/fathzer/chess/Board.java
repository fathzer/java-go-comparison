package com.fathzer.chess;

import static com.fathzer.chess.Piece.BLOCKER;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/** A tiny chess board
 * WARNING: This is a very basic implementation. It is not a chess move generator.
 * It lacks important features like king safety, en passant, castling, promotion.
 */
public class Board {
    private final Piece[] pieces;
	private final List<Move> playedMoves;
	private final List<Piece> captures;

	public Board(Board copy) {
		this.pieces = copy.pieces.clone();
		this.playedMoves = new ArrayList<>(copy.playedMoves);
		this.captures = new ArrayList<>(copy.captures);
	}

    public Board(String fen) {
		this.playedMoves = new ArrayList<>();
		this.captures = new ArrayList<>();
        this.pieces = new Piece[120];
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
				final Piece piece = Piece.fromCode(c);
				if (piece == null || piece==BLOCKER) {
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

	public Piece getPiece(String uciSquare) {
		return pieces[getSquare(uciSquare)];
	}

	Piece getPiece(int square) {
		return pieces[square];
	}

	public List<Move> getMoves(boolean white) {
		final List<Move> moves = new ArrayList<>(40);
		for (int square = 20; square < 100; square++) {
			Piece piece = getPiece(square);
			if (piece != null && piece!=BLOCKER && piece.isWhite()==white) {
				MoveGenerators.get(piece).build(moves, this, square);
			}
		}
		return moves;
	}

	public void makeMove(Move move) {
		if (move==null) {
			throw new IllegalArgumentException("Move cannot be null");
		}	    
	    int from = move.from();
	    int to = move.to();
		if (from<20 || from>119 || to<20 || to>119) {
			throw new IllegalArgumentException("Illegal move");
		}
		Piece movingPiece = pieces[from];
		Piece capturedPiece = pieces[to];
		if (movingPiece==null || movingPiece==BLOCKER || (capturedPiece!=null && !capturedPiece.canBeCapturedBy(((Piece)movingPiece).isWhite()))) {
			throw new IllegalArgumentException("Illegal move");
		}
	    // Store the captured piece (if any)
	    captures.add(pieces[to]);
	    
	    // Move the piece
	    pieces[to] = pieces[from];
	    pieces[from] = null; // Empty the source square
	    
	    // Record the move
	    playedMoves.add(move);
	}

	public void unmakeMove() {
	    if (playedMoves.isEmpty()) {
	        throw new IllegalStateException("No moves to unmake");
	    }
	    
	    // Get the last move and remove it from history
	    Move lastMove = playedMoves.remove(playedMoves.size() - 1);
	    int from = lastMove.from();
	    int to = lastMove.to();
	    
	    // Restore the moved piece
	    pieces[from] = pieces[to];
	    
	    // Restore the captured piece (if any)
	    pieces[to] = captures.remove(captures.size() - 1);
	}

	@Override
	public String toString() {
		final StringBuilder b = new StringBuilder();
		for (char rank = '8'; rank >= '1'; rank--) {
			b.append(rank);
			b.append(' ');
			for (char file = 'a'; file <= 'h'; file++) {
				Piece piece = getPiece(file+""+rank);
				b.append(piece==null ? ' ' : piece.getCode());
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
