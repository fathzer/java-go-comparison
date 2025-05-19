package main

import (
	"fmt"
	"math/big"
	"os"
	"sync"
	"time"

	"hellogo/internal/flags"
	hellogo "hellogo/pkg/oop"
	"hellogo/pkg/pi"
)

func main() {
	loops, err := flags.ParseLoopsFlag(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Starting OOP concepts test...")
	oopConceptsTest()
	fmt.Println("Starting concurrency test with", *loops, "loops...")
	concurrencyTest(*loops)
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
