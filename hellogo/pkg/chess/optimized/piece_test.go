package optimized

import (
	"testing"
)

func TestFromCode(t *testing.T) {
	// Test all valid piece codes
	tests := []struct {
		code  rune
		piece int
	}{
		{'P', WhitePawn},
		{'N', WhiteKnight},
		{'B', WhiteBishop},
		{'R', WhiteRook},
		{'Q', WhiteQueen},
		{'K', WhiteKing},
		{'p', BlackPawn},
		{'n', BlackKnight},
		{'b', BlackBishop},
		{'r', BlackRook},
		{'q', BlackQueen},
		{'k', BlackKing},
		{'X', None}, // Invalid code returns None
		{' ', None}, // Invalid code returns None
	}

	for _, tt := range tests {
		piece := FromCode(tt.code)
		if piece != tt.piece {
			t.Errorf("FromCode('%c') = %d, want %d", tt.code, piece, tt.piece)
		}
	}
}

func TestGetCode(t *testing.T) {
	tests := []struct {
		piece int
		code  rune
	}{
		{None, ' '},
		{WhitePawn, 'P'},
		{WhiteKnight, 'N'},
		{WhiteBishop, 'B'},
		{WhiteRook, 'R'},
		{WhiteQueen, 'Q'},
		{WhiteKing, 'K'},
		{BlackPawn, 'p'},
		{BlackKnight, 'n'},
		{BlackBishop, 'b'},
		{BlackRook, 'r'},
		{BlackQueen, 'q'},
		{BlackKing, 'k'},
	}

	for _, tt := range tests {
		got := GetCode(tt.piece)
		if got != tt.code {
			t.Errorf("GetCode(%d) = '%c' (0x%X), want '%c' (0x%X)", tt.piece, got, got, tt.code, tt.code)
		}
	}
}

func TestIsWhite(t *testing.T) {
	// Test white pieces
	whitePieces := []int{WhitePawn, WhiteKnight, WhiteBishop, WhiteRook, WhiteQueen, WhiteKing}
	for _, piece := range whitePieces {
		if !IsWhite(piece) {
			t.Errorf("IsWhite(%d) = false, want true", piece)
		}
	}

	// Test black pieces
	blackPieces := []int{BlackPawn, BlackKnight, BlackBishop, BlackRook, BlackQueen, BlackKing}
	for _, piece := range blackPieces {
		if IsWhite(piece) {
			t.Errorf("IsWhite(%d) = true, want false", piece)
		}
	}
}

func TestCanBeCapturedBy(t *testing.T) {
	// Test white pieces can be captured by black and not by white
	whitePieces := []int{WhitePawn, WhiteKnight, WhiteBishop, WhiteRook, WhiteQueen, WhiteKing}
	for _, piece := range whitePieces {
		if !CanBeCapturedBy(piece, false) {
			t.Errorf("CanBeCapturedBy(%d, false) = false, want true", piece)
		}
		if CanBeCapturedBy(piece, true) {
			t.Errorf("CanBeCapturedBy(%d, true) = true, want false", piece)
		}
	}

	// Test black pieces can be captured by white and not by black
	blackPieces := []int{BlackPawn, BlackKnight, BlackBishop, BlackRook, BlackQueen, BlackKing}
	for _, piece := range blackPieces {
		if !CanBeCapturedBy(piece, true) {
			t.Errorf("CanBeCapturedBy(%d, true) = false, want true", piece)
		}
		if CanBeCapturedBy(piece, false) {
			t.Errorf("CanBeCapturedBy(%d, false) = true, want false", piece)
		}
	}

	// Test special cases
	if CanBeCapturedBy(Blocker, true) || CanBeCapturedBy(Blocker, false) {
		t.Error("Blocker should not be capturable")
	}
}
