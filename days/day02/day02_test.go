package day02

import (
	"context"
	"testing"
)

func TestDay02Part01(t *testing.T) {
	ctx := context.Background()

	var expectedSum int64 = 2
	actualSum, err := Part1(ctx, "test_part1.txt")
	//var expectedSum int64 = 3574690
	//actualSum, err := Part1(ctx, "input.txt")

	if err != nil {
		t.Fatalf("Error calling Day01_Part1: %v", err)
	}

	if actualSum != expectedSum {
		t.Errorf("Expected sum %d but got %d", expectedSum, actualSum)
	} else {
		t.Logf("Success! Expected sum %d and got %d", expectedSum, actualSum)
	}
}

func TestDay02Part02(t *testing.T) {
	ctx := context.Background()

	// var expectedSum int64 = 31
	// actualSum, err := Part2(ctx, "test_part2.txt")
	var expectedSum int64 = 22565391
	actualSum, err := Part2(ctx, "input.txt")

	if err != nil {
		t.Fatalf("Error calling Day01_Part2: %v", err)
	}

	if actualSum != expectedSum {
		t.Errorf("Expected sum %d but got %d", expectedSum, actualSum)
	} else {
		t.Logf("Success! Expected sum %d and got %d", expectedSum, actualSum)
	}
}
