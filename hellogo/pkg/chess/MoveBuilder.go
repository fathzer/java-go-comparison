package chess

// MoveConstructor creates a move from two indices.
type MoveConstructor[T any] interface {
	Create(from, to int) T
}

// Capturable represents a piece that can be captured.
type Capturable interface {
	CanBeCapturedBy(white bool) bool
}

// Explorable represents a board that can be explored for moves.
type Explorable interface {
	GetCapturable(index int) Capturable
	GetRank(index int) int
}

// MoveBuilder scans the board for legal moves from a given square.
type MoveBuilder[T any] interface {
	Build(moves *[]T, board Explorable, from int, moveBuilder MoveConstructor[T])
}
