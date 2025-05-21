package com.fathzer.chess;

public enum Direction {
    NORTH,
    SOUTH,
    EAST,
    WEST,
    NORTH_EAST,
    NORTH_WEST,
    SOUTH_EAST,
    SOUTH_WEST;

    public int getDelta() {
        switch (this) {
            case NORTH: return +10;
            case SOUTH: return -10;
            case EAST: return +1;
            case WEST: return -1;
            case NORTH_EAST: return +11;
            case NORTH_WEST: return +9;
            case SOUTH_EAST: return -9;
            case SOUTH_WEST: return -11;
            default: return 0;
        }
    }
}
