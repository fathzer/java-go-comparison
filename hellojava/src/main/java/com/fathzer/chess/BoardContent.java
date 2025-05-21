package com.fathzer.chess;

public interface BoardContent {
    boolean canBeCapturedBy(boolean white);
    default boolean canMove() { return true; }
}
