package com.fathzer.chess;

public record Move(int from, int to) {
    public static Move fromUCI(String uci) {
        if (uci.length() != 4) {
            throw new IllegalArgumentException("Invalid UCI move: " + uci);
        }
        return new Move(Board.getSquare(uci.substring(0, 2)), Board.getSquare(uci.substring(2)));
    }
    
    public Move(String from, String to) {
        this(Board.getSquare(from), Board.getSquare(to));
    }

    @Override
    public String toString() {
        return Board.getUCI(from) + Board.getUCI(to);
    }
}
