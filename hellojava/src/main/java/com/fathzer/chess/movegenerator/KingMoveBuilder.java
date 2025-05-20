package com.fathzer.chess.movegenerator;

import com.fathzer.chess.Direction;

public class KingMoveBuilder extends BasicMoveBuilder {
    public KingMoveBuilder(boolean isWhite) {
        super(new int[] { Direction.NORTH.getDelta(), Direction.SOUTH.getDelta(), Direction.EAST.getDelta(), Direction.WEST.getDelta(), Direction.NORTH_EAST.getDelta(), Direction.NORTH_WEST.getDelta(), Direction.SOUTH_EAST.getDelta(), Direction.SOUTH_WEST.getDelta() }, isWhite);
    }
}
