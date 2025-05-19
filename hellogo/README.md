# The Golang version of the Hello (java-go-comparison) project.

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

The program has the following arguments:

- `--piLoops` or `-pl`: Number of Pi calculations to perform (default: 2000)
