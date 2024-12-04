package day04

import (
	"bufio"
	"context"
	"fmt"
	"github.com/mlouage/AdventOfCode.2024/internal/utils"
	"os"
)

type PuzzleProcessor struct {
	data [][]string
}

func NewPuzzleProcessor(lines int, chars int) *PuzzleProcessor {
	arrayOfSlices := make([][]string, lines)
	for i := range arrayOfSlices {
		arrayOfSlices[i] = make([]string, chars)
	}
	return &PuzzleProcessor{
		data: arrayOfSlices,
	}
}

func Part1(ctx context.Context, filename string) (int, error) {
	puzzle, err := processPuzzle(ctx, filename)
	if err != nil {
		return 0, fmt.Errorf("processing file: %w", err)
	}

	_ = puzzle

	return 0, nil
}

func Part2(ctx context.Context, filename string) (int, error) {
	return 0, nil
}

func processPuzzle(ctx context.Context, filename string) (*PuzzleProcessor, error) {
	numLines, err := countLines(filename)
	if err != nil {
		return nil, fmt.Errorf("counting lines: %w", err)
	}
	numChar, err := countCharacters(filename)
	if err != nil {
		return nil, fmt.Errorf("counting characters: %w", err)
	}

	processor := NewPuzzleProcessor(numLines, numChar)
	if err != nil {
		return nil, fmt.Errorf("reading number pairs: %w", err)
	}

	return processor, nil
}

func countLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	counter := utils.NewLineCounter(0, "")
	numLines, err := counter.Count(file)
	if err != nil {
		return 0, err
	}

	return numLines, err
}

func countCharacters(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numChar = 0
	if scanner.Scan() {
		firstLine := scanner.Text()
		numChar = len(firstLine)
	}

	if err := scanner.Err(); err != nil {
		return numChar, err
	}

	return numChar, err
}
