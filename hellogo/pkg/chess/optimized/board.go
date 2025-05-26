// Package optimized provides optimized data structures for chess engine implementation.
package optimized

import "strconv"

// Board represents a chess board with piece placement and move history.
type Board struct {
	pieces      []int
	playedMoves *IntList
}

// NewBoard creates a new board from a FEN string.
func NewBoard(fen string) *Board {
	b := &Board{
		playedMoves: NewIntList(),
		pieces:      make([]int, 120),
	}
	b.fillBlockers()

	rank := 7
	file := 0
	for _, c := range fen {
		if c >= '1' && c <= '8' {
			count := int(c - '0')
			if file+count > 8 {
				panic("Too many pieces on rank " + strconv.Itoa(rank+1))
			}
			for i := 0; i < count; i++ {
				b.pieces[21+rank*10+file] = None
				file++
			}
		} else if c == '/' {
			if file != 8 {
				panic("Invalid FEN: incomplete rank " + strconv.Itoa(rank+1))
			}
			rank--
			file = 0
		} else {
			piece := FromCode(c)
			if piece == None {
				panic("Invalid FEN: unknown piece " + string(c))
			}
			b.pieces[21+rank*10+file] = piece
			file++
		}
	}
	return b
}

// NewBoardCopy creates a copy of an existing board.
func NewBoardCopy(original *Board) *Board {
	pieces := make([]int, len(original.pieces))
	copy(pieces, original.pieces)
	return &Board{
		pieces:      pieces,
		playedMoves: NewIntListFrom(original.playedMoves),
	}
}

// fillBlockers initializes the board edges with blocker pieces.
func (b *Board) fillBlockers() {
	// Fill top and bottom borders
	for i := 0; i < 20; i++ {
		b.pieces[i] = Blocker
		b.pieces[100+i] = Blocker
	}
	// Fill left and right borders
	for i := 2; i < 10; i++ {
		startRank := i * 10
		b.pieces[startRank] = Blocker
		b.pieces[startRank+9] = Blocker
	}
}

// GetSquare converts a UCI square string to an internal square index.
func GetSquare(uciSquare string) int {
	if len(uciSquare) != 2 {
		panic("Invalid UCI square: " + uciSquare)
	}
	file := int(uciSquare[0] - 'a')
	rank := int(uciSquare[1] - '1')
	return 21 + 10*rank + file
}

// GetRank returns the rank (0-7) of a square.
func GetRank(square int) int {
	return (square - 21) / 10
}

// GetUCI converts an internal square index to a UCI string.
func GetUCI(square int) string {
	square -= 21
	file := byte(square%10 + 'a')
	rank := strconv.Itoa(square/10 + 1)
	return string([]byte{file, rank[0]})
}

// GetPiece returns the piece at a given UCI square.
func (b *Board) GetPiece(uciSquare string) int {
	return b.pieces[GetSquare(uciSquare)]
}

// getPiece returns the piece at a given internal square index.
func (b *Board) getPiece(square int) int {
	return b.pieces[square]
}

// MakeMove applies a move to the board.
func (b *Board) MakeMove(move int) {
	from := From(move)
	to := To(move)

	// Move the piece
	b.pieces[to] = b.pieces[from]
	b.pieces[from] = None

	// Record the move with the captured piece
	b.playedMoves.Add(move)
}

// UnmakeMove undoes the last move made on the board.
func (b *Board) UnmakeMove() {
	// Get the last move and remove it from history
	move := b.playedMoves.RemoveLast()
	from := From(move)
	to := To(move)

	// Restore the moved piece
	b.pieces[from] = b.pieces[to]

	// Restore the captured piece (if any)
	b.pieces[to] = Capture(move)
}

// String returns a string representation of the board.
func (b *Board) String() string {
	var result []byte
	for rank := 7; rank >= 0; rank-- {
		result = append(result, byte('1'+rank), ' ')
		for file := 0; file < 8; file++ {
			square := 21 + rank*10 + file
			result = append(result, byte(GetCode(b.pieces[square])))
		}
		result = append(result, '\n')
	}
	result = append(result, "  abcdefgh"...)
	return string(result)
}
