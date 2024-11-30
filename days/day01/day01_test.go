package day01

import (
	"testing"
)

func TestDay01Part01(t *testing.T) {
	expectedSum := 142
	actualSum, err := Part1("test_part1.txt")

	if err != nil {
		t.Fatalf("Error calling Day01_Part1: %v", err)
	}

	if actualSum != expectedSum {
		t.Errorf("Expected sum %d but got %d", expectedSum, actualSum)
	} else {
		t.Logf("Success! Expected sum %d and got %d", expectedSum, actualSum)
	}
}

func TestDay01Part02(t *testing.T) {
	expectedSum := 281
	actualSum, err := Part2("test_part2.txt")

	if err != nil {
		t.Fatalf("Error calling Day01_Part2: %v", err)
	}

	if actualSum != expectedSum {
		t.Errorf("Expected sum %d but got %d", expectedSum, actualSum)
	} else {
		t.Logf("Success! Expected sum %d and got %d", expectedSum, actualSum)
	}
}
