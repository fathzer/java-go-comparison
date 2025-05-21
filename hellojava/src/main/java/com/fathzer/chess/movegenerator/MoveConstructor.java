package com.fathzer.chess.movegenerator;

@FunctionalInterface
public interface MoveConstructor<T> {
	T create(int from, int to);
}
