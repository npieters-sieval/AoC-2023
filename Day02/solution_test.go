package Day02

import (
	"testing"
)

func TestPart01(t *testing.T) {
	testString := "8"
	msg, err := Part01("C:\\Users\\minec\\Documents\\repo\\AoC-2023\\Day02\\input\\test")
	if msg != testString || err != nil {
		t.Fatalf(`part01() = %q, %v, want %q, error`, msg, err, testString)
	}
}

func TestPart02(t *testing.T) {
	testString := "2286"
	msg, err := Part02("C:\\Users\\minec\\Documents\\repo\\AoC-2023\\Day02\\input\\testPart2")
	if msg != testString || err != nil {
		t.Fatalf(`part02() = %q, %v, want "", error`, msg, err)
	}
}
