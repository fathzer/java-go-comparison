package pi

import (
	"strings"
	"testing"
)

func TestComputePi(t *testing.T) {
	// Test with a small precision
	piVal, err := ComputePi(10)
	if err != nil {
		t.Fatalf("ComputePi returned error: %v", err)
	}
	if piVal == nil {
		t.Fatal("ComputePi returned nil value")
	}

	// Convert to string with 10 digits after decimal point
	piStr := piVal.Text('f', 10)
	if len(piStr) < 12 { // "3." + 10 digits
		t.Errorf("Pi string too short: %s", piStr)
	}
	parts := strings.SplitN(piStr, ".", 2)
	if len(parts) != 2 || len(parts[1]) != 10 {
		t.Errorf("Expected 10 digits after decimal point, got: %s", piStr)
	}
	if !strings.HasPrefix(piStr, "3.14159") {
		t.Errorf("Pi string should start with 3.14159, got: %s", piStr)
	}

	// Pi should be close to 3.14 for low precision
	piVal3, err := ComputePi(3)
	if err != nil {
		t.Fatalf("ComputePi(3) returned error: %v", err)
	}
	piStr3 := piVal3.Text('f', 3)
	if piStr3 != "3.142" {
		// Note: 3.141 contains the right decimals, but 3.142 is closer to pi
		t.Errorf("Expected Pi string '3.142', got: %s", piStr3)
	}

	// Error cases
	_, err = ComputePi(0)
	if err == nil {
		t.Error("Expected error for precision 0, got nil")
	}
	_, err = ComputePi(-5)
	if err == nil {
		t.Error("Expected error for negative precision, got nil")
	}
}
