package chess

import (
	"testing"
)

func TestValidMove(t *testing.T) {
	move := NewMove(0, 16)
	if move.From() != 0 {
		t.Errorf("expected from=0, got %d", move.From())
	}
	if move.To() != 16 {
		t.Errorf("expected to=16, got %d", move.To())
	}
}

func TestToString(t *testing.T) {
	move, err := NewMoveFromStrings("e2", "e4")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if move.String() != "e2e4" {
		t.Errorf("expected e2e4, got %s", move.String())
	}
}

func TestEquality(t *testing.T) {
	move1 := NewMove(0, 16)
	move2 := NewMove(0, 16)
	move3 := NewMove(16, 0)

	if !move1.Equals(move2) {
		t.Errorf("expected move1 == move2")
	}
	if move1.Equals(move3) {
		t.Errorf("expected move1 != move3")
	}
	// HashCode is not idiomatic in Go, so not tested.
}
