package com.fathzer.chess;

import static org.junit.jupiter.api.Assertions.*;

import org.junit.jupiter.api.Test;

import com.fathzer.chess.Perft.Result;

class PerftTest {

	@Test
	void test() {
		final Board board = new Board("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR");
		final Perft perft = new Perft();
		
		assertThrows(IllegalArgumentException.class, () -> perft.perft(null, 1, false));
		assertThrows(IllegalArgumentException.class, () -> perft.perft(board, -1, false));
		assertThrows(IllegalArgumentException.class, () -> perft.perft(board, 0, false));
		
		final Result result1 = perft.perft(board, 1, true);
		assertEquals(20, result1.leafNodesCount());
		result1.divide().forEach((move, count) -> assertEquals(1, count));	
		
		final Result result2 = perft.perft(board, 2, true);
		assertEquals(400, result2.leafNodesCount());
		
		final Board board2 = new Board("rnbqkbnr/pp1ppppp/2p5/8/6P1/2P5/PP1PPP1P/RNBQKBNR");
		assertEquals(21, perft.perft(board2, 1, false).leafNodesCount());
		assertEquals(463, perft.perft(board2, 2, false).leafNodesCount());
	}
}
