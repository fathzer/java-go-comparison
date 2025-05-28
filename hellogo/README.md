# The Golang version of the Hello (java-go-comparison) project

## Prerequisites

- Go 1.24.3

## Setup and Running

1. First, ensure all dependencies are properly set up:

   ```bash
   go mod tidy
   ```

2. To run the program:

   ```bash
   go run cmd/main.go
   ```

3. To build the program:

   ```bash
   go build cmd/main.go
   ```

## Program Arguments

The program accepts the following command-line arguments:

- `--piLoops` or `-pl`: Number of Pi calculations to perform (default: 2000)
- `--fen` or `-f`: FEN position to use for Perft (default: rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR)
- `--blackPlaying` or `-b`: If the black player is playing (default: false)
- `--perftDepth` or `-pd`: Depth for the Perft test (default: 6)
- `--perftLoops` or `-ptl`: Number of Perft (Performance Test) to perform (default: 5)
- `--perftCoolingTime` or `-pct`: Cooling time between Perft (Performance Test) computations (default: 0)
- `--optimizedPerft` or `-o`: Use optimized Perft (Performance Test) (default: false)

### Example Usage

```bash
go run cmd/main.go -pl=1000 -ptl=10 -pct=5000 -o
```
