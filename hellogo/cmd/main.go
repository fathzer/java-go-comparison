package main

import (
	"fmt"
	"math/big"
	"os"
	"sync"
	"time"

	"hellogo/internal/flags"
	"hellogo/internal/perft"
	hellogo "hellogo/pkg/oop"
	"hellogo/pkg/pi"
)

func main() {
	piLoops, err := flags.ParseIntFlag("pl", "piLoops", 2000, os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	perftDepth, err := flags.ParseIntFlag("pd", "perftDepth", 6, os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fen, err := flags.ParseStringFlag("f", "fen", "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR", os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	perftLoops, err := flags.ParseIntFlag("ptl", "perftLoops", 5, os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	blackPlaying, err := flags.ParseBoolFlag("b", "blackPlaying", os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	optimizedPerft, err := flags.ParseBoolFlag("o", "optimizedPerft", os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	coolingTime, err := flags.ParseIntFlag("pct", "perftCoolingTime", 0, os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	oopConceptsTest()
	concurrencyTest(*piLoops)
	if *perftLoops > 0 {
		perftTest(*fen, !blackPlaying, *perftDepth, *perftLoops, time.Duration(*coolingTime)*time.Millisecond, optimizedPerft)
	}
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

func perftTest(fen string, whitePlaying bool, depth int, nbLoops int, pauseDuration time.Duration, optimized bool) {
	optimizedStr := ""
	if optimized {
		optimizedStr = "optimized "
	}
	fmt.Printf("Starting %sPerft test with %dms cooling time before every perft computation ...\n", optimizedStr, pauseDuration.Milliseconds())

	// First execution with detailed output
	time.Sleep(pauseDuration)
	firstResult, err := perft.Run(fen, whitePlaying, depth, optimized)
	if err != nil {
		fmt.Printf("Error running perft test: %v\n", err)
		return
	}

	durationMs := firstResult.Duration.Milliseconds()
	leafNodes := firstResult.Result.LeafNodesCount()
	searchedNodes := firstResult.Result.SearchedNodesCount()
	fmt.Printf("First execution found: %d leaf nodes with %d move generation at depth %d in %d ms for %s\n",
		leafNodes, searchedNodes, depth, durationMs, fen)

	if nbLoops <= 1 {
		return
	}

	// Subsequent executions with progress tracking
	var totalDuration int64 = 0
	nbLoops--
	for i := 1; i <= nbLoops; i++ {
		fmt.Print(".")
		time.Sleep(pauseDuration)

		loopResult, err := perft.Run(fen, whitePlaying, depth, optimized)
		if err != nil {
			fmt.Printf("\nError running perft test on loop %d: %v\n", i, err)
			return
		}

		// Verify results match using the Equals method for consistency with Java version
		if !loopResult.Result.Equals(firstResult.Result) {
			fmt.Printf("\nResult mismatch at loop %d\n", i)
			return
		}

		totalDuration += loopResult.Duration.Milliseconds()
	}

	// Print final summary
	if nbLoops > 0 {
		avgDuration := totalDuration / int64(nbLoops)
		fmt.Printf("\nduration (%d loops): %d ms (average: %d ms)\n", nbLoops, totalDuration, avgDuration)
	}
}
