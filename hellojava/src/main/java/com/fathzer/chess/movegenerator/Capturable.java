package com.fathzer.chess.movegenerator;

@FunctionalInterface
public interface Capturable {
    boolean canBeCapturedBy(boolean white);
}
