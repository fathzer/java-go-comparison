package com.fathzer.chess;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

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

    /**
     * The results of a Perft (Performance Test) calculation.
     */
    public static class Result {
        private long searchedNodesCount;
        private long leafNodesCount;
        private final Map<Move, Long> nodesPerMove;

        private Result() {
             this.nodesPerMove = new HashMap<>();
        }
        /** Gets the number of leaf nodes
         * @return a long
         */
        public long leafNodesCount() {
            return leafNodesCount;
        }

        /** Gets the number of nodes for which the move generation has been searched
         * @return a long
         */
        public long searchedNodesCount() {
            return searchedNodesCount;
        }

        /** Gets the number of nodes per move at first depth
         * @return a map of moves to the number of nodes
         */
        public Map<Move, Long> divide() {
            return nodesPerMove;
        }
    }
    
    /** Performs a non bulk Perft (Performance Test) calculation.
     * @param board The board to run the performance test on.
     * @param depth The depth to run the performance test to
     * @param whitePlaying true if white is playing, false otherwise
     * @return a non null result
     */
    public Result perft(Board board, int depth, boolean whitePlaying) {
        return perft(board, depth, Type.NON_BULK, whitePlaying);
    }

    /**  Performs a Perft (Performance Test) calculation.
     * @param board The board to run the performance test on.
     * @param depth The depth to run the performance test to
     * @param type The type of Perft to run
     * @param whitePlaying true if white is playing, false otherwise
     * @return a non null result
     */
    public Result perft(Board board, int depth, Type type, boolean whitePlaying) {
        if (board==null) {
            throw new IllegalArgumentException("Board cannot be null");
        }
    	if (depth<=0) {
    		throw new IllegalArgumentException("Depth must be greater than 0");
    	}
        final Result result = new Result();
        result.leafNodesCount = perft(board, result, depth, depth, type, whitePlaying);
        return result;
    }

    private long perft(Board board, Result result, int depth, int originalDepth, Type type, boolean whitePlaying) {
        result.searchedNodesCount++;
        final List<Move> moves = board.getMoves(whitePlaying);
        if (depth == 1 && type == Type.NON_BULK) {
            return moves.size();
        } else if (depth == 0) {
            return 1;
        }
        long leafNodesCount = 0;
        for (Move move : moves) {
            board.makeMove(move);
            long moveCount = perft(board, result, depth - 1, originalDepth, type, !whitePlaying);
            if (depth == originalDepth) {
                result.nodesPerMove.put(move, moveCount);
            }
            leafNodesCount += moveCount;
            board.unmakeMove();
        }
        return leafNodesCount;
    }
}
