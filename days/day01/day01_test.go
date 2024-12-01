package day01

import (
	"testing"
)

func TestDay01Part01(t *testing.T) {
	// var expectedSum int64 = 11
	var expectedSum int64 = 3574690
	actualSum, err := Part1("input.txt")

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
	// var expectedSum int64 = 31
	var expectedSum int64 = 22565391
	actualSum, err := Part2("input.txt")

	if err != nil {
		t.Fatalf("Error calling Day01_Part2: %v", err)
	}

	if actualSum != expectedSum {
		t.Errorf("Expected sum %d but got %d", expectedSum, actualSum)
	} else {
		t.Logf("Success! Expected sum %d and got %d", expectedSum, actualSum)
	}
}
