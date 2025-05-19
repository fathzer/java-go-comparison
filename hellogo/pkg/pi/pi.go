package pi

import (
	"fmt"
	"math"
	"math/big"
)

// computePi calculates Pi to the specified number of decimal digits
func ComputePi(digits int) (*big.Float, error) {
	if digits <= 0 {
		return nil, fmt.Errorf("number of digits must be greater than 0: %d", digits)
	}
	precision := uint(digits * 4) // bits of precision; estimate conservatively

	// constants
	c := new(big.Float).SetPrec(precision)
	c.SetFloat64(426880)
	sqrt10005 := sqrtBigFloat(big.NewFloat(10005), precision)
	c.Mul(c, sqrt10005)

	sum := new(big.Float).SetPrec(precision)
	sum.SetFloat64(13591409)

	M := big.NewInt(1)
	L := big.NewInt(13591409)
	X := big.NewInt(1)

	for k := 1; k < digits/14+1; k++ {
		kBig := big.NewInt(int64(k))

		// Compute (6k - 5)
		a := big.NewInt(6*int64(k) - 5)
		// Compute (2k - 1)
		b := big.NewInt(2*int64(k) - 1)
		// Compute (6k - 1)
		c := big.NewInt(6*int64(k) - 1)

		// numerator = M * a * b * c
		numerator := new(big.Int).Mul(M, a)
		numerator.Mul(numerator, b)
		numerator.Mul(numerator, c)

		// denominator = k^3
		kCubed := new(big.Int).Mul(kBig, kBig)
		kCubed.Mul(kCubed, kBig)

		M = new(big.Int).Quo(numerator, kCubed)

		// L = L + 545140134
		L.Add(L, big.NewInt(545140134))

		// X = X * -262537412640768000
		X.Mul(X, big.NewInt(-262537412640768000))

		// term = (M * L) / X
		termNumer := new(big.Int).Mul(M, L)
		term := new(big.Float).SetPrec(precision).SetInt(termNumer)
		term.Quo(term, new(big.Float).SetPrec(precision).SetInt(X))

		sum.Add(sum, term)
	}

	pi := new(big.Float).SetPrec(precision)
	pi.Quo(c, sum)
	// Set the final precision
	pi.SetPrec(uint(digits * 4))
	return pi, nil
}

// sqrtBigFloat computes the square root using Newton-Raphson method
func sqrtBigFloat(x *big.Float, prec uint) *big.Float {
	tolerance := new(big.Float).SetPrec(prec).SetFloat64(1e-20)
	guess := new(big.Float).SetPrec(prec).SetFloat64(math.Sqrt(getFloat64(x)))
	diff := new(big.Float).SetPrec(prec)

	for {
		// next = (guess + x / guess) / 2
		next := new(big.Float).SetPrec(prec)
		next.Quo(x, guess)
		next.Add(next, guess)
		next.Quo(next, big.NewFloat(2))

		diff.Sub(next, guess)
		if diff.Cmp(tolerance) < 0 && diff.Cmp(new(big.Float).Neg(tolerance)) > 0 {
			break
		}
		guess = next
	}
	return guess
}

// Helper to convert big.Float to float64 safely
func getFloat64(f *big.Float) float64 {
	val, _ := f.Float64()
	return val
}
