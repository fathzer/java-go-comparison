# java-go-comparison
Some pieces of code in [Go](https://golang.org/) and [Java](https://www.java.com/)

The goal of this project was, for me, to compare the two languages in terms of:
- syntax
- performance
- ease of use

The [hellogo](hellogo) and [hellojava](hellojava) folders contain a simple program, respectively in Go and Java, that is supposed (except if some bugs remains) have the same behavior.

I precise that I'm Go newby, so I might have done some things in a suboptimal way. I develop in Java for a long time, so I might have done most things in a Java-like way.

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

I will try to implement other tests with other algorithms to compare the performance of both languages without that kind of bias.

### Testing
Parsing and Pi computing have test classes in both languages.
In Go, the tests are in the folder as the source, while in Java, the tests are in the `src/test/java` folder and sources in the `src/main/java` folder.
