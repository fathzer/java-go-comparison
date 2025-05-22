package chess

import (
	"testing"
)

func TestFromCode(t *testing.T) {
	// Test all valid piece codes
	tests := []struct {
		code rune
		want Piece
	}{
		{'P', WHITE_PAWN},
		{'N', WHITE_KNIGHT},
		{'B', WHITE_BISHOP},
		{'R', WHITE_ROOK},
		{'Q', WHITE_QUEEN},
		{'K', WHITE_KING},
		{'p', BLACK_PAWN},
		{'n', BLACK_KNIGHT},
		{'b', BLACK_BISHOP},
		{'r', BLACK_ROOK},
		{'q', BLACK_QUEEN},
		{'k', BLACK_KING},
	}
	for _, tt := range tests {
		got := FromCode(tt.code)
		if got == nil || *got != tt.want {
			t.Errorf("FromCode(%q) = %v, want %v", tt.code, got, tt.want)
		}
	}
	// Test invalid code
	if got := FromCode('X'); got != nil {
		t.Errorf("FromCode('X') = %v, want nil", got)
	}
	if got := FromCode(' '); got != nil {
		t.Errorf("FromCode(' ') = %v, want nil", got)
	}
}

func TestGetCode(t *testing.T) {
	if WHITE_PAWN.GetCode() != 'P' {
		t.Errorf("WHITE_PAWN.GetCode() = %c, want 'P'", WHITE_PAWN.GetCode())
	}
	if WHITE_KNIGHT.GetCode() != 'N' {
		t.Errorf("WHITE_KNIGHT.GetCode() = %c, want 'N'", WHITE_KNIGHT.GetCode())
	}
	if BLACK_PAWN.GetCode() != 'p' {
		t.Errorf("BLACK_PAWN.GetCode() = %c, want 'p'", BLACK_PAWN.GetCode())
	}
	if BLACK_KING.GetCode() != 'k' {
		t.Errorf("BLACK_KING.GetCode() = %c, want 'k'", BLACK_KING.GetCode())
	}
}

func TestIsWhitePiece(t *testing.T) {
	// Test white pieces
	whitePieces := []*Piece{&WHITE_PAWN, &WHITE_KNIGHT, &WHITE_BISHOP, &WHITE_ROOK, &WHITE_QUEEN, &WHITE_KING}
	for _, p := range whitePieces {
		if !p.IsWhitePiece() {
			t.Errorf("%v.IsWhitePiece() = false, want true", p)
		}
	}
	// Test black pieces
	blackPieces := []*Piece{&BLACK_PAWN, &BLACK_KNIGHT, &BLACK_BISHOP, &BLACK_ROOK, &BLACK_QUEEN, &BLACK_KING}
	for _, p := range blackPieces {
		if p.IsWhitePiece() {
			t.Errorf("%v.IsWhitePiece() = true, want false", p)
		}
	}
}

func TestCanBeCapturedBy(t *testing.T) {
	// Test white pieces can be captured by black
	whitePieces := []*Piece{&WHITE_PAWN, &WHITE_KNIGHT, &WHITE_BISHOP, &WHITE_ROOK, &WHITE_QUEEN, &WHITE_KING}
	for _, p := range whitePieces {
		if !p.CanBeCapturedBy(false) {
			t.Errorf("%v.CanBeCapturedBy(false) = false, want true", p)
		}
		if p.CanBeCapturedBy(true) {
			t.Errorf("%v.CanBeCapturedBy(true) = true, want false", p)
		}
	}
	// Test black pieces can be captured by white
	blackPieces := []*Piece{&BLACK_PAWN, &BLACK_KNIGHT, &BLACK_BISHOP, &BLACK_ROOK, &BLACK_QUEEN, &BLACK_KING}
	for _, p := range blackPieces {
		if !p.CanBeCapturedBy(true) {
			t.Errorf("%v.CanBeCapturedBy(true) = false, want true", p)
		}
		if p.CanBeCapturedBy(false) {
			t.Errorf("%v.CanBeCapturedBy(false) = true, want false", p)
		}
	}
}

func TestBlocker(t *testing.T) {
	if (&BLOCKER).CanBeCapturedBy(true) {
		t.Errorf("BLOCKER.CanBeCapturedBy(true) = true, want false")
	}
	if (&BLOCKER).CanBeCapturedBy(false) {
		t.Errorf("BLOCKER.CanBeCapturedBy(false) = true, want false")
	}
}
