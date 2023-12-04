package Day03

import (
	"testing"
)

func TestPart01(t *testing.T) {
	testString := "4361"
	msg, err := Part01("C:\\Users\\minec\\Documents\\repo\\AoC-2023\\Day03\\input\\exercise")
	if msg != testString || err != nil {
		t.Fatalf(`part01() = %q, %v, want %q, error`, msg, err, testString)
	}
}

func TestPart02(t *testing.T) {
	testString := "84289137"
	msg, err := Part02("C:\\Users\\minec\\Documents\\repo\\AoC-2023\\Day03\\input\\exercise")
	if msg != testString || err != nil {
		t.Fatalf(`part02() = %q, %v, want %q, error`, msg, err, testString)
	}
}
