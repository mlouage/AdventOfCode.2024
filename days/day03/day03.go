package day03

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type NumberPair struct {
	First  int64
	Second int64
}

type FileProcessor struct {
	pairs []NumberPair
}

func NewFileProcessor() *FileProcessor {
	return &FileProcessor{
		pairs: []NumberPair{},
	}
}

func Part1(ctx context.Context, filename string) (int64, error) {
	processor, err := processFile(ctx, filename, false)
	if err != nil {
		return 0, fmt.Errorf("processing file: %w", err)
	}
	return processor.calc(), nil
}

func Part2(ctx context.Context, filename string) (int64, error) {
	processor, err := processFile(ctx, filename, true)
	if err != nil {
		return 0, fmt.Errorf("processing file: %w", err)
	}
	return processor.calc(), nil
}

func processFile(ctx context.Context, filename string, strict bool) (*FileProcessor, error) {
	processor := NewFileProcessor()
	pairs, err := parseNumberPairs(ctx, filename, strict)
	if err != nil {
		return nil, fmt.Errorf("reading number pairs: %w", err)
	}

	processor.pairs = pairs
	return processor, nil
}

func parseNumberPairs(ctx context.Context, filename string, strict bool) ([]NumberPair, error) {
	canAppend := true

	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	input := string(content)

	pattern := `(?:mul\((\d{1,3}),(\d{1,3})\)|(?:do|don't)\(\))`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(input, -1)

	pairs := make([]NumberPair, 0, len(matches))

	for i, match := range matches {
		if strings.HasPrefix(match[0], "do(") {
			canAppend = true
		}
		if strings.HasPrefix(match[0], "don't(") {
			canAppend = false
		}
		if strict && !canAppend {
			continue
		}
		if strings.HasPrefix(match[0], "mul") {
			first, err1 := strconv.ParseInt(match[1], 10, 64)
			second, err2 := strconv.ParseInt(match[2], 10, 64)

			if err1 != nil || err2 != nil {
				fmt.Printf("Error parsing numbers from match %d: %s\n", i+1, match[0])
				continue
			}

			pair := NumberPair{
				First:  first,
				Second: second,
			}
			pairs = append(pairs, pair)
		}
	}

	return pairs, nil
}

func (fp *FileProcessor) calc() int64 {
	var totalSum int64
	for _, pair := range fp.pairs {
		totalSum = totalSum + (pair.First * pair.Second)
	}

	return totalSum
}
