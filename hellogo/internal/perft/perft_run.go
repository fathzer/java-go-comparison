package perft

import (
	"fmt"
	"time"

	"hellogo/pkg/chess"
	"hellogo/pkg/chess/optimized"
	chesscommon "hellogo/pkg/chess/common"
)

// ResultWithDuration holds the result of a Perft test along with its execution duration
type ResultWithDuration struct {
	Result   *chesscommon.PerftResult
	Duration time.Duration
}

// Runner defines the interface for running Perft tests
type Runner interface {
	run(fen string, whitePlaying bool, depth int) *chesscommon.PerftResult
}

type standardPerft struct{}
type optimizedPerft struct{}

var (
	standardPerftInstance = &standardPerft{}
	optimizedPerftInstance = &optimizedPerft{}
)

// Run executes a Perft test with the given parameters and returns the result with timing
func Run(fen string, whitePlaying bool, depth int, useOptimized bool) (ResultWithDuration, error) {
	var runner Runner
	if useOptimized {
		runner = optimizedPerftInstance
	} else {
		runner = standardPerftInstance
	}

	start := time.Now()
	result := runner.run(fen, whitePlaying, depth)
	elapsed := time.Since(start)

	return ResultWithDuration{
		Result:   result,
		Duration: elapsed,
	}, nil
}

func (p *standardPerft) run(fen string, whitePlaying bool, depth int) *chesscommon.PerftResult {
	board, err := chess.NewBoard(fen)
	if err != nil {
		panic(fmt.Sprintf("invalid FEN: %v", err))
	}
	perft := chess.NewPerft()
	// Using NonBulk as the default type to match Java implementation
	result, _ := perft.PerftWithType(board, depth, chesscommon.NonBulk, whitePlaying)
	return result
}

func (p *optimizedPerft) run(fen string, whitePlaying bool, depth int) *chesscommon.PerftResult {
	board := optimized.NewBoard(fen)
	perft := optimized.NewPerft()
	// Using NonBulk as the default type to match Java implementation
	return perft.PerftWithType(board, depth, chesscommon.NonBulk, whitePlaying)
}
