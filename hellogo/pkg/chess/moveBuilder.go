package chess

// MoveBuilder scans the board for legal moves from a given square.
type MoveBuilder interface {
	Build(moves *[]Move, board *Board, from int)
}
