package chess

// MoveConstructor creates a move from two indices.
type MoveConstructor interface {
	Create(from, to int) interface{}
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
type MoveBuilder interface {
	Build(moves interface{}, board Explorable, from int, moveBuilder MoveConstructor)
}
