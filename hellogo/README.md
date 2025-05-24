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
- `--perftDepth` or `-pd`: Depth for the Perft test (default: 6)

### Example Usage

```bash
go run cmd/main.go -pl=1000 -pd=4
```
