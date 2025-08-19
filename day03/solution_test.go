package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	var testStr = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	var actual = part1(string(testStr))
	if actual != 161 {
		t.Errorf("Incorrect result")
	}
}

func TestPart2(t *testing.T) {
	var testStr = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	var actual = part2(string(testStr))
	var wanted = 48
	if actual != wanted {
		t.Errorf("Incorrect result. Got=%d Want=%d", actual, wanted)
	}
}
