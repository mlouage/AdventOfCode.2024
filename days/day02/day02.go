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

func Part1(ctx context.Context, filename string) (int, error) {
	reports, err := readReports(ctx, filename)
	if err != nil {
		return 0, fmt.Errorf("processing file: %w", err)
	}

	count := 0
	for _, report := range reports {
		if report.safe() {
			count++
		}
	}
	return count, nil
}

func Part2(ctx context.Context, filename string) (int, error) {
	reports, err := readReports(ctx, filename)
	if err != nil {
		return 0, fmt.Errorf("processing file: %w", err)
	}

	count := 0
	for _, report := range reports {
		if report.safe() {
			count++
		} else {
			if report.safeWithDamp() {
				count++
			}
		}
	}
	return count, nil
}

func (r *Report) safe() bool {
	if len(r.Levels) < 2 {
		return false
	}

	dir := getDirection(r.Levels[0], r.Levels[1])
	diff := math.Abs(float64(r.Levels[0]) - float64(r.Levels[1]))
	if dir == Levelling || diff > 3 {
		return false
	}

	for i := 1; i < len(r.Levels)-1; i++ {
		currDir := getDirection(r.Levels[i], r.Levels[i+1])
		if currDir == Levelling {
			return false
		}

		diff = math.Abs(float64(r.Levels[i]) - float64(r.Levels[i+1]))
		if dir != currDir || diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func (r *Report) safeWithDamp() bool {
	levels := r.Levels

	// Pre-allocate the slice with capacity = len-1
	newLevels := make([]int, 0, len(levels)-1)

	// Try removing each element one at a time
	for i := range levels {
		// Reset the slice but keep the capacity
		newLevels = newLevels[:0]

		// Build new slice excluding index i
		newLevels = append(newLevels, levels[:i]...)
		newLevels = append(newLevels, levels[i+1:]...)

		if (&Report{Levels: newLevels}).safe() {
			return true
		}
	}

	return false
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
	levels := make([]int, len(input))

	for i := range input {
		level, err := strconv.Atoi(input[i])
		if err != nil {
			return []int{}, fmt.Errorf("parsing level: %w", err)
		}
		levels[i] = level
	}

	return levels, nil
}
