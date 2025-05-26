package chesscommon

// PerftType represents the type of Perft calculation.
// Please note that since the move generator generates only legal moves, both types should yield the same result.
type PerftType int

const (
	// NonBulk is a non-bulk Perft (Performance Test) calculation; moves at last depth are not played.
	NonBulk PerftType = iota
	// Bulk is a bulk Perft (Performance Test) calculation.
	Bulk
)
