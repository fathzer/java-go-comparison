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

The program has the following arguments:

- `--piLoops` or `-pl`: Number of Pi calculations to perform (default: 2000)
