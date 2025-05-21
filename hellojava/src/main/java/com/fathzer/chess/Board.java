package com.fathzer.chess;

import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;

/** A tiny chess board
 * WARNING: This is a very basic implementation. It is not a chess move generator.
 * It lacks important features like king safety, en passant, castling, promotion.
 */
public class Board {
    private final BoardContent[] content;
	private final List<Move> playedMoves;
	private final List<Piece> captures;

	public Board(Board copy) {
		this.content = copy.content.clone();
		this.playedMoves = new LinkedList<>(copy.playedMoves);
		this.playedMoves.addAll(copy.playedMoves);
		this.captures = new LinkedList<>(copy.captures);
		this.captures.addAll(copy.captures);
	}

    public Board(String fen) {
		this.playedMoves = new LinkedList<>();
		this.captures = new LinkedList<>();
        this.content = new BoardContent[120];
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
				if (piece == null) {
					throw new IllegalArgumentException("Invalid FEN: unknown piece " + c);
				}
				this.content[21+rank*10+file] = piece;
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
    	Arrays.fill(content, 0, 20, Blocker.INSTANCE);
    	Arrays.fill(content, 100, 120, Blocker.INSTANCE);
    	for (int i=2; i<10; i++) {
			final int startRank = i*10;
    		content[startRank] = Blocker.INSTANCE;
    		content[startRank+9] = Blocker.INSTANCE;
    	}
    }
	
	public static int getSquare(String uciSquare) {
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
	
	public static int getRank(int square) {
		return (square-21)/10;
	}
	
	public static String getUCI(int square) {
		square -= 21;
		return (char)(square%10 + 'a') + "" +(square/10 + 1);
	}

	public Piece getPiece(String uciSquare) {
		return (Piece)content[getSquare(uciSquare)];
	}

	public BoardContent getContent(int square) {
		return content[square];
	}

	public List<Move> getMoves(boolean white) {
		final List<Move> moves = new LinkedList<>();
		for (int square = 20; square < 100; square++) {
			BoardContent squareContent = getContent(square);
			if (squareContent != null && squareContent.canMove() && ((Piece)squareContent).isWhite()==white) {
				((Piece)squareContent).getMoveBuilder().build(moves, this, square);
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
		BoardContent movingPiece = content[from];
		BoardContent capturedPiece = content[to];
		if (movingPiece==null || !movingPiece.canMove() || (capturedPiece!=null && !capturedPiece.canBeCapturedBy(((Piece)movingPiece).isWhite()))) {
			throw new IllegalArgumentException("Illegal move");
		}
	    // Store the captured piece (if any)
	    captures.add((Piece)content[to]);
	    
	    // Move the piece
	    content[to] = content[from];
	    content[from] = null; // Empty the source square
	    
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
	    content[from] = content[to];
	    
	    // Restore the captured piece (if any)
	    content[to] = captures.remove(captures.size() - 1);
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
