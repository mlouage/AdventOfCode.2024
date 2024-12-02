package day02

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Report struct {
	Levels []int
}
type Direction int

const (
	Increasing Direction = iota
	Decreasing
	Levelling
)

func Part1(ctx context.Context, filename string) (sum int64, err error) {
	reports, err := readReports(ctx, filename)
	if err != nil {
		return 0, fmt.Errorf("processing file: %w", err)
	}

	_ = reports
	return 0, nil
}

func Part2(ctx context.Context, filename string) (sum int64, err error) {
	return 0, nil
}

func (report *Report) isSafe() bool {
	var isSafe bool
	var capacity = len(report.Levels)

	if capacity < 2 {
		return false
	}

	direction := isIncreasing(report.Levels[0], report.Levels[1])

	return isSafe
}

func isIncreasing(level1, level2 int) Direction {
	if level1 == level2 {
		return Levelling
	}

	if level1 < level2 {
		return Increasing
	}

	return Decreasing
}

func readReports(ctx context.Context, filename string) ([]Report, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	var reports []Report
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			levels, err := parseLevels(scanner.Text())
			if err != nil {
				return nil, fmt.Errorf("line %d: %w", len(pairs)+1, err)
			}
			reports = append(reports, Report{Levels: levels})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanning file: %w", err)
	}

	return reports, nil
}

func parseLevels(s string) ([]int, error) {
	input := strings.Fields(s)
	var capacity = len(input)

	levels := make([]int, capacity)

	for i := range input {
		level, err := strconv.Atoi(input[i])
		if err != nil {
			return []int{}, fmt.Errorf("parsing level: %w", err)
		}
		levels = append(levels, level)
	}

	return levels, nil
}
