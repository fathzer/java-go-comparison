package com.fathzer.chess.optimized;

import java.util.Arrays;

import com.fathzer.chess.common.PerftResult;

/**
 * <a href="https://www.chessprogramming.org/Perft">Perft, ('Performance Test')</a> is a Performance Test is a debugging function
 * that walks the move generation tree of strictly legal moves to count all the leaf nodes of a certain depth,
 * which can be compared to predetermined values and used to isolate bugs.
 */
public class Perft {
    /** The type of Perft calculation.
     * <br>Please note that as Calvin move generator generates only legal moves both type should yield the same result.
     */
    public enum Type {
        /** A non bulk Perft (Performance Test) calculation; moves at last depth are not played  */
        NON_BULK,
        /** A bulk Perft (Performance Test) calculation. */
        BULK
    }

    /** Performs a non bulk Perft (Performance Test) calculation.
     * @param board The board to run the performance test on.
     * @param depth The depth to run the performance test to
     * @param whitePlaying true if white is playing, false otherwise
     * @return a non null result
     */
    public PerftResult<Integer> perft(Board board, int depth, boolean whitePlaying) {
        return perft(board, depth, Type.NON_BULK, whitePlaying);
    }

    /**  Performs a Perft (Performance Test) calculation.
     * @param board The board to run the performance test on.
     * @param depth The depth to run the performance test to
     * @param type The type of Perft to run
     * @param whitePlaying true if white is playing, false otherwise
     * @return a non null result
     */
    public PerftResult<Integer> perft(Board board, int depth, Type type, boolean whitePlaying) {
        if (board==null) {
            throw new IllegalArgumentException("Board cannot be null");
        }
    	if (depth<=0) {
    		throw new IllegalArgumentException("Depth must be greater than 0");
    	}
        final PerftResult<Integer> result = new PerftResult<>();
        final IntList[] moveListCache = new IntList[depth+1];
        Arrays.setAll(moveListCache, i -> new IntList());
        result.setLeafNodesCount(perft(moveListCache, board, result, depth, depth, type, whitePlaying));
        return result;
    }

    private long perft(IntList[] moveListCache, Board board, PerftResult<Integer> result, int depth, int originalDepth, Type type, boolean whitePlaying) {
        result.incrementSearchedNodesCount();
        final IntList moves = moveListCache[depth];
        board.getMoves(moves, whitePlaying);
        if (depth == 1 && type == Type.NON_BULK) {
            return moves.size();
        } else if (depth == 0) {
            return 1;
        }
        long leafNodesCount = 0;
        for (int moveIndex = 0; moveIndex < moves.size(); moveIndex++) {
            int move = moves.get(moveIndex);
            board.makeMove(move);
            long moveCount = perft(moveListCache, board, result, depth - 1, originalDepth, type, !whitePlaying);
            if (depth == originalDepth) {
                result.setNodesPerMove(move, moveCount);
            }
            leafNodesCount += moveCount;
            board.unmakeMove();
        }
        return leafNodesCount;
    }
}
