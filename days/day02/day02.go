package day02

import (
	"bufio"
	"context"
	"fmt"
	"math"
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

func Part1(ctx context.Context, filename string) (sum int, err error) {
	reports, err := readReports(ctx, filename)
	if err != nil {
		return 0, fmt.Errorf("processing file: %w", err)
	}

	count := 0
	for _, report := range reports {
		if report.isSafe() {
			count++
		}
	}
	return count, nil
}

func Part2(ctx context.Context, filename string) (sum int64, err error) {
	return 0, nil
}

func (report *Report) isSafe() bool {
	isSafe := true
	var capacity = len(report.Levels)

	if capacity < 2 {
		return false
	}

	direction := getDirection(report.Levels[0], report.Levels[1])
	difference := math.Abs(float64(report.Levels[0]) - float64(report.Levels[1]))
	if direction == Levelling || (difference > 3) {
		return false
	}

	for i := 1; i < capacity-1; i++ {
		currentDirection := getDirection(report.Levels[i], report.Levels[i+1])
		if currentDirection == Levelling {
			isSafe = false
			break
		}
		difference = math.Abs(float64(report.Levels[i]) - float64(report.Levels[i+1]))
		if direction == currentDirection && (difference >= 1 && difference <= 3) {
			direction = currentDirection
			continue
		} else {
			isSafe = false
			break
		}
	}

	return isSafe
}

func getDirection(level1, level2 int) Direction {
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
				return nil, fmt.Errorf("line %d: %w", len(levels)+1, err)
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
		levels[i] = level
	}

	return levels, nil
}
