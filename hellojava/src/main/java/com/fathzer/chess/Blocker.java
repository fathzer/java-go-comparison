package com.fathzer.chess;

@SuppressWarnings("s6548")
public class Blocker implements BoardContent {
    public static final Blocker INSTANCE = new Blocker();

    private Blocker() {}

    @Override
    public boolean canBeCapturedBy(boolean white) {
        return false;
    }

    @Override
    public boolean canMove() {
        return false;
    }

	@Override
	public String toString() {
		return "BLOCK";
	}
}
