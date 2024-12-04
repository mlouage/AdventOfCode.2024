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
	rows int
	cols int
}

func NewPuzzleProcessor(lines int, chars int) *PuzzleProcessor {
	arrayOfSlices := make([][]string, lines)
	for i := range arrayOfSlices {
		arrayOfSlices[i] = make([]string, chars)
	}
	return &PuzzleProcessor{
		data: arrayOfSlices,
		rows: lines,
		cols: chars,
	}
}

func Part1(ctx context.Context, filename string) (int, error) {
	puzzle, err := createPuzzle(ctx, filename)
	if err != nil {
		return 0, fmt.Errorf("processing file: %w", err)
	}

	sum := puzzle.countXmas()

	return sum, nil
}

func Part2(ctx context.Context, filename string) (int, error) {
	puzzle, err := createPuzzle(ctx, filename)
	if err != nil {
		return 0, fmt.Errorf("processing file: %w", err)
	}

	sum := puzzle.countMasX()

	return sum, nil
}

func createPuzzle(ctx context.Context, filename string) (*PuzzleProcessor, error) {
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

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			line := scanner.Text()
			for j, value := range line {
				processor.data[i][j] = string(value)
			}
			i++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanning file: %w", err)
	}

	return processor, nil
}

func (pp *PuzzleProcessor) countXmas() int {
	totalCount := 0

	for i := 0; i < pp.rows; i++ {
		for j := 0; j < pp.cols; j++ {
			if pp.data[i][j] == "X" {
				totalCount = totalCount + pp.hasXmas(i, j)
			}
		}
	}

	return totalCount
}

func (pp *PuzzleProcessor) countMasX() int {
	totalCount := 0

	for i := 0; i < pp.rows; i++ {
		for j := 0; j < pp.cols; j++ {
			if pp.data[i][j] == "A" {
				if pp.hasMasX(i, j) {
					totalCount++
				}
			}
		}
	}

	return totalCount
}

func (pp *PuzzleProcessor) hasXmas(row int, col int) int {
	totalCount := 0

	if pp.checkHorizontal(row, col, col+3) {
		totalCount++
	}

	if pp.checkHorizontal(row, col, col-3) {
		totalCount++
	}

	if pp.checkVertical(row, row+3, col) {
		totalCount++
	}

	if pp.checkVertical(row, row-3, col) {
		totalCount++
	}

	if pp.checkDiagonal(row, col, row+3, col+3) {
		totalCount++
	}

	if pp.checkDiagonal(row, col, row-3, col-3) {
		totalCount++
	}

	if pp.checkDiagonal(row, col, row+3, col-3) {
		totalCount++
	}

	if pp.checkDiagonal(row, col, row-3, col+3) {
		totalCount++
	}

	return totalCount
}

func (pp *PuzzleProcessor) hasMasX(row int, col int) bool {
	if !pp.isValidX(row, col) {
		return false
	}

	if pp.data[row+1][col+1] == "M" && pp.data[row-1][col-1] == "S" &&
		pp.data[row-1][col+1] == "M" && pp.data[row+1][col-1] == "S" ||
		pp.data[row+1][col+1] == "M" && pp.data[row-1][col-1] == "S" &&
			pp.data[row+1][col-1] == "M" && pp.data[row-1][col+1] == "S" ||
		pp.data[row-1][col-1] == "M" && pp.data[row+1][col+1] == "S" &&
			pp.data[row-1][col+1] == "M" && pp.data[row+1][col-1] == "S" ||
		pp.data[row-1][col-1] == "M" && pp.data[row+1][col+1] == "S" &&
			pp.data[row+1][col-1] == "M" && pp.data[row-1][col+1] == "S" {
		return true
	}

	return false
}

func (pp *PuzzleProcessor) checkHorizontal(row int, colStart int, colEnd int) bool {
	if !pp.isValid(row, colStart) || !pp.isValid(row, colEnd) {
		return false
	}

	if colStart < colEnd {
		if pp.data[row][colStart] == "X" &&
			pp.data[row][colStart+1] == "M" &&
			pp.data[row][colStart+2] == "A" &&
			pp.data[row][colStart+3] == "S" {
			return true
		}
	} else {
		if pp.data[row][colStart] == "X" &&
			pp.data[row][colStart-1] == "M" &&
			pp.data[row][colStart-2] == "A" &&
			pp.data[row][colStart-3] == "S" {
			return true
		}
	}

	return false
}

func (pp *PuzzleProcessor) checkDiagonal(rowStart int, colStart int, rowEnd int, colEnd int) bool {
	if !pp.isValid(rowStart, colStart) || !pp.isValid(rowEnd, colEnd) {
		return false
	}

	if rowStart < rowEnd && colStart < colEnd {
		if pp.data[rowStart][colStart] == "X" &&
			pp.data[rowStart+1][colStart+1] == "M" &&
			pp.data[rowStart+2][colStart+2] == "A" &&
			pp.data[rowStart+3][colStart+3] == "S" {
			return true
		}
	} else if rowStart < rowEnd && colStart > colEnd {
		if pp.data[rowStart][colStart] == "X" &&
			pp.data[rowStart+1][colStart-1] == "M" &&
			pp.data[rowStart+2][colStart-2] == "A" &&
			pp.data[rowStart+3][colStart-3] == "S" {
			return true
		}
	} else if rowStart > rowEnd && colStart > colEnd {
		if pp.data[rowStart][colStart] == "X" &&
			pp.data[rowStart-1][colStart-1] == "M" &&
			pp.data[rowStart-2][colStart-2] == "A" &&
			pp.data[rowStart-3][colStart-3] == "S" {
			return true
		}
	} else if rowStart > rowEnd && colStart < colEnd {
		if pp.data[rowStart][colStart] == "X" &&
			pp.data[rowStart-1][colStart+1] == "M" &&
			pp.data[rowStart-2][colStart+2] == "A" &&
			pp.data[rowStart-3][colStart+3] == "S" {
			return true
		}
	}

	return false
}

func (pp *PuzzleProcessor) checkVertical(rowStart int, rowEnd int, col int) bool {
	if !pp.isValid(rowStart, col) || !pp.isValid(rowEnd, col) {
		return false
	}

	if rowStart < rowEnd {
		if pp.data[rowStart][col] == "X" &&
			pp.data[rowStart+1][col] == "M" &&
			pp.data[rowStart+2][col] == "A" &&
			pp.data[rowStart+3][col] == "S" {
			return true
		}
	} else {
		if pp.data[rowStart][col] == "X" &&
			pp.data[rowStart-1][col] == "M" &&
			pp.data[rowStart-2][col] == "A" &&
			pp.data[rowStart-3][col] == "S" {
			return true
		}
	}

	return false
}

func (pp *PuzzleProcessor) isValid(row int, col int) bool {
	if row >= 0 && row <= pp.rows-1 && col >= 0 && col <= pp.cols-1 {
		return true
	}

	return false
}

func (pp *PuzzleProcessor) isValidX(row int, col int) bool {
	if pp.isValid(row-1, col-1) && pp.isValid(row+1, col+1) && pp.isValid(row-1, col+1) && pp.isValid(row+1, col-1) {
		return true
	}

	return false
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
