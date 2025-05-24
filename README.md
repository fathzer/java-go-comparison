# java-go-comparison
Some pieces of code in [Go](https://golang.org/) and [Java](https://www.java.com/)

The goal of this project was, for me, to compare the two languages in terms of:
- syntax
- performance
- ease of use

The [hellogo](hellogo) and [hellojava](hellojava) folders contain a simple program, respectively in Go and Java, that is supposed (except if some bugs remains) have the same behavior.

I precise that I'm Go newby, most of the GO code has been translated by AI (ChatGPT 4.1 and SWE-1), so it might be suboptimal. I develop in Java for a long time, so I might have done most things in a Java-like way.

## Contents

The `cmd/main.go` and `cmd/main.java` files are the entry points of the programs.

### Arguments parsing
Program argument parsing is implemented in the `internal/flags` go module and in the `java/com/fathzer/hellogo/Flags.java` java Class.
The `args.go` and `args.java` files are the arguments parsing code.

### OOP concepts
I implemented in both languages the same OOP concepts:
- Class (abstract and concrete) ... even if Go doesn't have classes, it has structs that can mimic classes
- Inheritance ... even if Go doesn't have inheritance, it has interfaces that can mimic inheritance
- Polymorphism ... No surprise, Go does not support this out of the box, but it can be mimicked with interfaces
- Encapsulation

This part is implemented in `pkg/oop` go module and in the `java/com/fathzer/hellogo/oop` java package.

### Error management
I'm not sure about the best way to handle errors in Go, but clearly Java and Go have totally different ways to manage errors.
The Java exceptions are thrown and caught wherever you want, while Go uses error values that can be checked only by the direct caller.
Fortunately, there are no complex error handling in this example.

### Performance throught Pi decimal computing
The programs compute more than the first 1000 decimal in a loop.

This part is implemented in `pkg/pi` go module and in the `java/com/fathzer/pi/PiCalculator.java` java class.

Here are the performance results for 1000 decimals to 3000 decimals (The program computes the decimals in a loop starting from 1000 decimals to 3000 decimals):

| Language | Time (ms) |
| --- | --- |
| Go | 540 |
| Java | 3100 |

A main difference between Java and Go is that the class that represent a big decimal number is quite different in both languages.
In Java, the `java.math.BigDecimal` class is immutable, while in Go, the `big.Float` type is mutable. I presume that this is the reason why the Java program is far slower (it should make far more memory allocations).

### Performance through chess Perft test

To have a more balanced performance comparison, I implemented a simplified chess Perft test in both languages. The Perft test is a standard way to verify the correctness of a chess engine's move generator by counting the number of nodes (possible positions) at a given depth.

This part is implemented in the `pkg/chess` go module and in the `java/com/fathzer/chess` java package.

Here are the performance results for Perft test at depth 6 from the initial position:

| Language | Time 1st run (ms) | Time (ms) | Nodes |
| --- | --- | --- | --- |
| Go | 1750 | 2350 | 51,004,063 |
| Java | 1343 | 1220 | 51,004,063 |

These results show that for this particular test case, Java is twice faster than Go! I was suprised by this result as I tried to make the Go program as close possible to Java (using references instead of copying objects, using the same data structure (for instance using ArrayList and not LinkedList), etc). Here are the possible reasons that SWE-1 AI gave to me:

1. Go's Slice vs Java's Array:

   - Go's slice has a small overhead (pointer + length + capacity)
   - Java's array is more lightweight in this case

2. Memory Allocation Patterns:

   - Go's garbage collector might behave differently under pressure
   - Java's JIT might optimize the hot paths better

3. Function Call Overhead:

   - Go's method calls might have slightly more overhead
   - Java's JIT can inline aggressively

4. Bounds Checking:

   - Go performs bounds checking on slice/array access
   - Java's JIT might optimize out some bounds checks

**Note:** The chess package is *simplified* in both languages, it does not implement the full chess rules. It lacks the king safety validation, the castling, the en passant capture and the pawn promotion.

### Testing

Parsing, Pi computing and chess packages have test classes in both languages.
In Go, the tests are in the folder as the source, while in Java, the tests are in the `src/test/java` folder and sources in the `src/main/java` folder.
