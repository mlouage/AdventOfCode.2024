package day01

import (
	"bufio"
	"context"
	"fmt"
	"github.com/mlouage/AdventOfCode.2024/internal/utils"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type NumberPair struct {
	First  int64
	Second int64
}

type FileProcessor struct {
	pairs []NumberPair
	lists [][]int64
}

func NewFileProcessor(capacity int) *FileProcessor {
	return &FileProcessor{
		pairs: make([]NumberPair, 0, capacity),
		lists: make([][]int64, 2),
	}
}

func Part1(ctx context.Context, filename string) (sum int64, err error) {
	processor, err := processFile(ctx, filename)
	if err != nil {
		return 0, fmt.Errorf("processing file: %w", err)
	}

	return processor.calculateAbsoluteDifferences(), nil
}

func Part2(ctx context.Context, filename string) (sum int64, err error) {
	processor, err := processFile(ctx, filename)
	if err != nil {
		return 0, fmt.Errorf("processing file: %w", err)
	}

	return processor.calculateMatchingNumbers(), nil
}

func processFile(ctx context.Context, filename string) (*FileProcessor, error) {
	numLines, err := countLines(filename)
	if err != nil {
		return nil, fmt.Errorf("counting lines: %w", err)
	}

	processor := NewFileProcessor(numLines)
	pairs, err := readNumberPairs(ctx, filename)
	lists := transformToLists(pairs)
	if err != nil {
		return nil, fmt.Errorf("reading number pairs: %w", err)
	}

	processor.pairs = pairs
	processor.lists = lists
	return processor, nil
}

func transformToLists(pairs []NumberPair) [][]int64 {
	lists := make([][]int64, 2)

	var capacity = len(pairs)

	lists[0] = make([]int64, capacity)
	lists[1] = make([]int64, capacity)

	for i := range pairs {
		lists[0][i] = pairs[i].First
		lists[1][i] = pairs[i].Second
	}

	return lists
}

func (fp *FileProcessor) calculateAbsoluteDifferences() int64 {
	lists := make([][]int64, 2)
	copy(lists, fp.lists)

	sort.Slice(lists[0], func(i, j int) bool {
		return lists[0][i] < lists[0][j]
	})

	sort.Slice(lists[1], func(i, j int) bool {
		return lists[1][i] < lists[1][j]
	})

	var capacity = len(lists[0])
	var totalSum int64
	for i := 0; i < capacity; i++ {
		distance := lists[0][i] - lists[1][i]
		totalSum += int64(math.Abs(float64(distance)))
	}

	return totalSum
}

func (fp *FileProcessor) calculateMatchingNumbers() int64 {
	var totalSum int64

	for i := range fp.lists[0] {
		number := fp.lists[0][i]
		count := count(
			fp.lists[1],
			func(x int64) bool {
				return x == number
			})

		totalSum = totalSum + (number * count)
	}

	return totalSum
}

func readNumberPairs(ctx context.Context, filename string) ([]NumberPair, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	var pairs []NumberPair
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			pair, err := parseNumberPair(scanner.Text())
			if err != nil {
				return nil, fmt.Errorf("line %d: %w", len(pairs)+1, err)
			}
			pairs = append(pairs, pair)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanning file: %w", err)
	}

	return pairs, nil
}

func parseNumberPair(s string) (NumberPair, error) {
	numbers := strings.Fields(s)
	if len(numbers) != 2 {
		return NumberPair{}, fmt.Errorf("expected 2 numbers, got %d", len(numbers))
	}

	first, err := strconv.ParseInt(numbers[0], 10, 64)
	if err != nil {
		return NumberPair{}, fmt.Errorf("parsing first number: %w", err)
	}

	second, err := strconv.ParseInt(numbers[1], 10, 64)
	if err != nil {
		return NumberPair{}, fmt.Errorf("parsing second number: %w", err)
	}

	return NumberPair{First: first, Second: second}, nil
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

func count[T any](slice []T, f func(T) bool) int64 {
	var count int64 = 0
	for _, s := range slice {
		if f(s) {
			count++
		}
	}

	return count
}
