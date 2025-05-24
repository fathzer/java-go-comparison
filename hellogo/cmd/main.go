package main

import (
	"fmt"
	"math/big"
	"os"
	"sync"
	"time"

	"hellogo/internal/flags"
	"hellogo/pkg/chess"
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

	fmt.Println("Starting OOP concepts test...")
	oopConceptsTest()
	fmt.Println("Starting concurrency test with", *loops, "loops...")
	concurrencyTest(*loops)
	fmt.Println("Starting Perft test...")
	perftTest(*perftDepth)
}

func oopConceptsTest() {
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
	for i := 1; i <= 5; i++ {
		board1, _ := chess.NewBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR")
		doPerft(depth, board1, true)

		board2, _ := chess.NewBoard("rnbqkbnr/pp1ppppp/2p5/8/6P1/2P5/PP1PPP1P/RNBQKBNR")
		doPerft(depth, board2, false)
	}
}

func doPerft(depth int, board *chess.Board, whitePlaying bool) {
	const durationFormat = "duration (%d depth): %v"
	const foundFormat = "Found: %d leaf nodes. Generated: %d"

	perft := chess.NewPerft()
	start := time.Now()
	result, err := perft.Perft(board, depth, whitePlaying)
	if err != nil {
		fmt.Printf("Error in perft: %v\n", err)
		return
	}
	duration := time.Since(start)
	fmt.Printf(durationFormat+"\n", depth, duration)
	fmt.Printf(foundFormat+"\n", result.LeafNodes, result.SearchedNodes)
}
