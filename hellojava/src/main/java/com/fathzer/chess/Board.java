package com.fathzer.chess;

import java.util.LinkedList;
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
		this.playedMoves = new LinkedList<>(copy.playedMoves);
		this.playedMoves.addAll(copy.playedMoves);
		this.captures = new LinkedList<>(copy.captures);
		this.captures.addAll(copy.captures);
	}

    public Board(String fen) {
		this.playedMoves = new LinkedList<>();
		this.captures = new LinkedList<>();
        this.pieces = new Piece[64];
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
				if (piece == null) {
					throw new IllegalArgumentException("Invalid FEN: unknown piece " + c);
				}
				this.pieces[rank*8+file] = piece;
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

	public int getSquare(int rank, int file) {
		return rank*8+file;
	}

	public Piece getPiece(int square) {
		return pieces[square];
	}

	public List<Move> getLegalMoves() {
		final List<Move> moves = new LinkedList<>();
		for (int square = 0; square < 64; square++) {
			Piece piece = getPiece(square);
			if (piece != null) {
				moves.addAll(piece.getScanner().build(moves, this, square));
			}
		}
		return new LinkedList<>();
	}

	public void makeMove(Move move) {
	    int from = move.from();
	    int to = move.to();
	    
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
		for (int rank = 7; rank >= 0; rank--) {
			for (int file = 0; file < 8; file++) {
				Piece piece = getPiece(getSquare(rank, file));
				b.append(piece==null ? ' ' : piece.getCode());
			}
			b.append('\n');
		}
		return b.toString();
	}
}
