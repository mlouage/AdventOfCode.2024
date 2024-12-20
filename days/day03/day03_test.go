package day03

import (
	"testing"
)

func TestDay03Part1(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int64
	}{
		{"sample input", "test_part1.txt", 161},
		{"real input", "input.txt", 175015740},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Part1(tt.filename)
			if err != nil {
				t.Fatalf("Part1(%q) error = %v", tt.filename, err)
			}
			if got != tt.want {
				t.Errorf("Part1(%q) = %v, want %v", tt.filename, got, tt.want)
			}
		})
	}
}

func TestDay03Part2(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     int64
	}{
		{"sample input", "test_part2.txt", 48},
		{"real input", "input.txt", 112272912},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Part2(tt.filename)
			if err != nil {
				t.Fatalf("Part1(%q) error = %v", tt.filename, err)
			}
			if got != tt.want {
				t.Errorf("Part1(%q) = %v, want %v", tt.filename, got, tt.want)
			}
		})
	}
}
