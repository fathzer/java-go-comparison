package com.fathzer.chess.movegenerator;

import com.fathzer.chess.Direction;

public class KnightMoveBuilder extends BasicMoveBuilder {
    public KnightMoveBuilder(boolean isWhite) {
        super(getDeltas(), isWhite);
    }

    private static int[] getDeltas() {
        final int[] deltas = new int[8];
        deltas[0] = 2*Direction.NORTH.getDelta()+Direction.EAST.getDelta();
        deltas[1] = 2*Direction.NORTH.getDelta()+Direction.WEST.getDelta();
        deltas[2] = 2*Direction.SOUTH.getDelta()+Direction.EAST.getDelta();
        deltas[3] = 2*Direction.SOUTH.getDelta()+Direction.WEST.getDelta();
        deltas[4] = Direction.NORTH.getDelta()+2*Direction.EAST.getDelta();
        deltas[5] = Direction.NORTH.getDelta()+2*Direction.WEST.getDelta();
        deltas[6] = Direction.SOUTH.getDelta()+2*Direction.EAST.getDelta();
        deltas[7] = Direction.SOUTH.getDelta()+2*Direction.WEST.getDelta();
        return deltas;
    }
}
