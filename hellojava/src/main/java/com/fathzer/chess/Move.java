package com.fathzer.chess;

public record Move(int from, int to) {
    public Move(int fromRank, int fromFile, int toRank, int toFile) {
        this(fromRank*8+fromFile, toRank*8+toFile);
    }
    
    public Move(int from, int to) {
        if (from < 0 || from > 63 || to < 0 || to > 63 || from == to) {
            throw new IllegalArgumentException("Invalid move");
        }
        this.from = from;
        this.to = to;
    }

    @Override
    public String toString() {
        return Notation.toUCI(from/8, from%8) + Notation.toUCI(to/8, to%8);
    }
}
