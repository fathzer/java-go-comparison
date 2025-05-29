package com.fathzer.chess.common;

/** The type of Perft calculation.
 * <br>Please note that as Calvin move generator generates only legal moves both type should yield the same result.
 */
public enum PerftType {
    /** A non bulk Perft (Performance Test) calculation; moves at last depth are not played  */
    NON_BULK,
    /** A bulk Perft (Performance Test) calculation. */
    BULK
}