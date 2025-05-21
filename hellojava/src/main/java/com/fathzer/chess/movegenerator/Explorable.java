package com.fathzer.chess.movegenerator;

public interface Explorable {
    Capturable getCapturable(int index);
    int getRank(int index);
}
