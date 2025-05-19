package com.fathzer.pi;

import java.math.BigDecimal;
import java.math.MathContext;
import java.math.RoundingMode;

public abstract class PiCalculator {
    private PiCalculator() {
        // Private constructor to prevent instantiation
    }

    // Scale controls how many digits we calculate (precision)
    public static BigDecimal computePi(int digits) {
        if (digits < 1) {
            throw new IllegalArgumentException("Number of digits must be at least 1");
        }
        MathContext mc = new MathContext(digits + 10, RoundingMode.HALF_UP); // extra digits for accuracy
        final BigDecimal c = new BigDecimal("426880").multiply(sqrt(new BigDecimal("10005"), mc));
        BigDecimal m = BigDecimal.ONE;
        BigDecimal x = BigDecimal.ONE;
        BigDecimal l = BigDecimal.valueOf(13591409);
        BigDecimal s = l;

        for (int k = 1; k < digits / 14 + 1; k++) {
            m = m.multiply(BigDecimal.valueOf((6 * k - 5) * (2 * k - 1) * (6 * k - 1)))
                 .divide(BigDecimal.valueOf(k * k * k), mc);
            l = l.add(BigDecimal.valueOf(545140134));
            x = x.multiply(BigDecimal.valueOf(-262537412640768000L), mc);
            BigDecimal term = m.multiply(l).divide(x, mc);
            s = s.add(term);
        }

        BigDecimal pi = c.divide(s, mc);
        return pi.setScale(digits, RoundingMode.DOWN);
    }

    // Square root using Newton-Raphson method
    private static BigDecimal sqrt(BigDecimal A, MathContext mc) {
        BigDecimal x0 = BigDecimal.ZERO;
        BigDecimal x1 = BigDecimal.valueOf(Math.sqrt(A.doubleValue()));
        while (!x0.equals(x1)) {
            x0 = x1;
            x1 = A.divide(x0, mc);
            x1 = x1.add(x0);
            x1 = x1.divide(BigDecimal.valueOf(2), mc);
        }
        return x1;
    }

}
