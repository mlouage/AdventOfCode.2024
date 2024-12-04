package day04

import (
	"context"
	"testing"
)

func TestDay04Part1(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{"sample input", "test_part1.txt", 18},
		{"real input", "input.txt", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Part1(context.Background(), tt.filename)
			if err != nil {
				t.Fatalf("Part1(%q) error = %v", tt.filename, err)
			}
			if got != tt.want {
				t.Errorf("Part1(%q) = %v, want %v", tt.filename, got, tt.want)
			}
		})
	}
}

func TestDay04Part2(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int
	}{
		{"sample input", "test_part2.txt", 0},
		{"real input", "input.txt", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Part2(context.Background(), tt.filename)
			if err != nil {
				t.Fatalf("Part1(%q) error = %v", tt.filename, err)
			}
			if got != tt.want {
				t.Errorf("Part1(%q) = %v, want %v", tt.filename, got, tt.want)
			}
		})
	}
}
