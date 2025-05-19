package com.fathzer.hellojava;

import static com.fathzer.pi.PiCalculator.computePi;

import java.math.BigDecimal;
import java.math.RoundingMode;
import java.util.concurrent.atomic.AtomicReference;
import java.util.stream.IntStream;

import com.fathzer.oop.ACoolThing;
import com.fathzer.oop.AbstractThing;
import com.fathzer.oop.CoolInterface;

public class Main {
    public static void main(String[] args) {
        try {
            int piLoops = Flags.parsePiLoops(args);

            System.out.println("Starting OOP concepts test...");
            oopConceptsTest();
            System.out.println("Starting concurrency test with " + piLoops + " loops...");
            concurrencyTest(piLoops);
        } catch (IllegalArgumentException e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }

    private static void oopConceptsTest() {
        AbstractThing aThing = new ACoolThing("my abstract thing");
        ((CoolInterface)aThing).doSomethingCool();
        aThing.doSomething();

        CoolInterface iThing = new ACoolThing("my cool interface");
        ((AbstractThing)iThing).doSomething();
        iThing.doSomethingCool();

        ACoolThing acoolThing = new ACoolThing("my cool interface");
        acoolThing.doSomething();
        acoolThing.doSomethingCool();
    }

    private static void concurrencyTest(int nbLoops) {
        if (nbLoops < 1) {
            throw new IllegalArgumentException("Number of loops must be at least 1");
        }
        final AtomicReference<BigDecimal> result = new AtomicReference<>(new BigDecimal(0));
        long start = System.currentTimeMillis();
        IntStream.range(0, nbLoops).parallel().forEach(i -> {
            var pi = computePi(1000+i);
            synchronized(result) {
                result.set(result.get().add(pi));
            }
        });
        System.out.println("duration ("+nbLoops+" loops): " + (System.currentTimeMillis() - start)+" ms");
        System.out.println("result: " + result.get().divide(BigDecimal.valueOf(nbLoops), 10, RoundingMode.HALF_UP));
    }
}
