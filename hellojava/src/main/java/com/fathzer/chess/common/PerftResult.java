package com.fathzer.chess.common;

import java.util.HashMap;
import java.util.Map;

/**
 * The results of a Perft (Performance Test) calculation.
 */
public class PerftResult<T> {
    private long searchedNodesCount;
    private long leafNodesCount;
    private final Map<T, Long> nodesPerMove;

    public PerftResult() {
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
    public Map<T, Long> divide() {
        return nodesPerMove;
    }

    public void incrementSearchedNodesCount() {
        searchedNodesCount++;
    }

    public void setLeafNodesCount(long leafNodesCount) {
        this.leafNodesCount = leafNodesCount;
    }

    public void setNodesPerMove(T move, long nodes) {
        nodesPerMove.put(move, nodes);
    }
}
