package com.fathzer.chess;

import java.util.Objects;

public class Move {
    private final int from;
    private final int to;

    Move(int from, int to) {
        this.from = from;
        this.to = to;
    }
    
    public static Move fromUCI(String uci) {
        if (uci.length() != 4) {
            throw new IllegalArgumentException("Invalid UCI move: " + uci);
        }
        return new Move(Board.getSquare(uci.substring(0, 2)), Board.getSquare(uci.substring(2)));
    }
    
    public Move(String from, String to) {
        this(Board.getSquare(from), Board.getSquare(to));
    }

    public int from() {
        return from;
    }

    public int to() {
        return to;
    }

    
    @Override
	public int hashCode() {
		return Objects.hash(from, to);
	}

	@Override
	public boolean equals(Object obj) {
		if (this == obj) return true;
		if (obj == null) return false;
		if (getClass() != obj.getClass()) return false;
		Move other = (Move) obj;
		return from == other.from && to == other.to;
	}

	@Override
    public String toString() {
        return Board.getUCI(from) + Board.getUCI(to);
    }
}
