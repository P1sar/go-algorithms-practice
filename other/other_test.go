package other

import (
	"testing"
)

func TestGCD(t *testing.T) {
	if gcd(874, 1994) != 2 {
		t.Fail()
	}
	if gcd(21212121, 12121212) != 3030303 {
		t.Fail()
	}
}

func TestLCM(t *testing.T) {
	if lcm(8,14) != 56  {
		t.Fail()
	}
}