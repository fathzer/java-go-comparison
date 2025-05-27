package main

import (
	"fmt"
	"math/big"
	"os"
	"sync"
	"time"

	"hellogo/internal/flags"
	"hellogo/pkg/chess"
	"hellogo/pkg/chess/optimized"
	hellogo "hellogo/pkg/oop"
	"hellogo/pkg/pi"
)

func main() {
	loops, err := flags.ParseLoopsFlag(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	perftDepth, err := flags.ParsePerftDepthFlag(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	oopConceptsTest()
	concurrencyTest(*loops)
	perftTest(*perftDepth)
}

func oopConceptsTest() {
	fmt.Println("Starting OOP concepts test...")
	// Demonstrate abstract thing interface
	aThing := hellogo.NewACoolThing("my abstract thing")
	aThing.DoSomething()
	var coolThing hellogo.CoolInterface = aThing
	coolThing.DoSomethingCool()

	// Demonstrate cool interface interface
	iThing := hellogo.NewACoolThing("my cool interface")
	iThing.DoSomething()
	iThing.DoSomethingCool()

	// Demonstrate concrete type
	acoolThing := hellogo.NewACoolThing("my cool interface")
	acoolThing.DoSomething()
	acoolThing.DoSomethingCool()
}

func concurrencyTest(nbLoops int) {
	fmt.Println("Starting concurrency test with", nbLoops, "loops...")
	var wg sync.WaitGroup
	var mu sync.Mutex

	result := new(big.Float).SetPrec(10000).SetFloat64(0)

	start := time.Now()

	for i := 0; i < nbLoops; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			pi, err := pi.ComputePi(1000 + i)
			if err != nil {
				fmt.Printf("Error computing Pi: %v\n", err)
				return
			} // increase precision with i

			mu.Lock()
			result.Add(result, pi)
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	duration := time.Since(start)

	// Compute average
	average := new(big.Float).SetPrec(10000)
	average.Quo(result, big.NewFloat(float64(nbLoops)))

	// Round to 10 decimal places for display
	outputStr := fmt.Sprintf("%.10f", average)

	fmt.Printf("duration (%d loops): %v\n", nbLoops, duration)
	fmt.Println("result:", outputStr)
}

func perftTest(depth int) {
	fmt.Println("Starting Perft test...")
	perftTestWithType(depth, false)
	fmt.Println("Starting optimized perft tests")
	perftTestWithType(depth, true)
}

func perftTestWithType(depth int, fast bool) {
	for i := 1; i <= 3; i++ {
		startFEN := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
		otherFEN := "rnbqkbnr/pp1ppppp/2p5/8/6P1/2P5/PP1PPP1P/RNBQKBNR"
		if fast {
			doFastPerft(startFEN, true, depth)
			doFastPerft(otherFEN, false, depth)
		} else {
			doPerft(startFEN, true, depth)
			doPerft(otherFEN, false, depth)
		}
	}
}

func doPerft(fen string, whitePlaying bool, depth int) {
	start := time.Now()
	board, err := chess.NewBoard(fen)
	if err != nil {
		fmt.Printf("Error creating board: %v\n", err)
		return
	}
	perft := chess.NewPerft()
	result, err := perft.Perft(board, depth, whitePlaying)
	if err != nil {
		fmt.Printf("Error in perft: %v\n", err)
		return
	}
	formatResult(result, depth, time.Since(start), fen)
}

func doFastPerft(fen string, whitePlaying bool, depth int) {
	start := time.Now()
	board := optimized.NewBoard(fen)
	perft := optimized.NewPerft()
	result := perft.Perft(board, depth, whitePlaying)
	formatResult(result, depth, time.Since(start), fen)
}

func formatResult(result interface{ LeafNodesCount() int64; SearchedNodesCount() int64 }, depth int, duration time.Duration, fen string) {
	fmt.Printf("Found: %d leaf nodes with %d move generation at depth %d in %v for %s\n",
		result.LeafNodesCount(), result.SearchedNodesCount(), depth, duration, fen)
}
