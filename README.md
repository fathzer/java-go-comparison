# java-go-comparison

Some pieces of code in [Go](https://golang.org/) and [Java](https://www.java.com/)

The goal of this project was, for me, to compare the two languages in terms of:

- syntax
- performance (all performance tests have been made on a Minisforum UM790 with AMD Ryzen 9 7940HS and 64GB RAM)
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

To have a more balanced performance comparison, I implemented a simplified chess Perft test in both languages. The [Perft](https://www.chessprogramming.org/Perft) test is a standard way to verify the correctness of a chess engine's move generator by counting the number of nodes (possible positions) at a given depth.

This part is implemented in the `pkg/chess` go module and in the `java/com/fathzer/chess` java package.
I also implemented an optimized version, that makes no memory allocation during the move generation (in `pkg/chess/optimized` and `java/com/fathzer/chess/optimized`).

Here are the performance results for singled threaded Perft test at depth 6 from the initial position (51,004,063 move generation, 120,909,581 leaf nodes):

| Language | Time 1st run (ms) | average time after 1st (ms) |
| --- | --- | --- |
| Go | 1720 | 2160 |
| Java | 1250 | 1450 |
| Go (optimized) | 890 | 900 |
| Java optimized | 900 | 1080 |

These results show that there's is no absolute truth. In the non optimized perft test, Java is faster, and slower after optmization. Maybe Go's garbage collector is less efficient than java one.

**Note:** The chess package is *simplified* in both languages, it does not implement the full chess rules. It lacks the king safety validation, the castling, the en passant capture and the pawn promotion.

### Testing

Parsing, Pi computing and chess packages have test classes in both languages.
In Go, the tests are in the folder as the source, while in Java, the tests are in the `src/test/java` folder and sources in the `src/main/java` folder.
