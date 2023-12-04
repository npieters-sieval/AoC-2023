package Day01

import (
	"testing"
)

func TestPart01(t *testing.T) {
	testString := "142"
	msg, err := Part01("C:\\Users\\minec\\Documents\\repo\\AoC-2023\\Day01\\input\\test")
	if msg != testString || err != nil {
		t.Fatalf(`part01() = %q, %v, want %q, error`, msg, err, testString)
	}
}

func TestPart02(t *testing.T) {
	testString := "281"
	msg, err := Part02("C:\\Users\\minec\\Documents\\repo\\AoC-2023\\Day01\\input\\testPart2")
	if msg != testString || err != nil {
		t.Fatalf(`part02() = %q, %v, want "", error`, msg, err)
	}
}
