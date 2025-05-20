package com.fathzer.chess;

public class Notation {
    private Notation() {}
    
    public static String toUCI(int rank, int file) {
        final char rankChar = (char)('1' + rank);
        final char fileChar = (char)('a' + file);
        return "" + fileChar + rankChar;
    }
}
