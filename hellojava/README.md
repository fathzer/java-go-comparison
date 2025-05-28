# The Java version of the Hello (java-go-comparison) project

## Prerequisites

- Java 17
- Having Maven installed

## Setup and Running

1. First, build the package:

   ```bash
   mvn clean package
   ```

2. To run the program:

   ```bash
   java -jar target/hellojava-0.0.1-SNAPSHOT.jar
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
java -jar target/hellojava-0.0.1-SNAPSHOT.jar -pl=1000 -ptl=10 -pct=5000 -o
```
