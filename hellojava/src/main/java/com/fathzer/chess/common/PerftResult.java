package com.fathzer.chess.common;

import java.util.Objects;

/**
 * The results of a Perft (Performance Test) calculation.
 */
public class PerftResult {
    private long searchedNodesCount;
    private long leafNodesCount;

    public PerftResult() {
    	super();
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

    public void incrementSearchedNodesCount() {
        searchedNodesCount++;
    }

    public void setLeafNodesCount(long leafNodesCount) {
        this.leafNodesCount = leafNodesCount;
    }

    @Override
    public boolean equals(Object obj) {
        if (this == obj) return true;
        if (obj == null || getClass() != obj.getClass()) return false;
        PerftResult that = (PerftResult) obj;
        return searchedNodesCount == that.searchedNodesCount && leafNodesCount == that.leafNodesCount;
    }

    @Override
    public int hashCode() {
        return Objects.hash(searchedNodesCount, leafNodesCount);
    }
}
