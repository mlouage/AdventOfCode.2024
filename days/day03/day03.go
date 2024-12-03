package day03

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	mulPattern = `(?:mul\((\d{1,3}),(\d{1,3})\)|(?:do|don't)\(\))`
	doCmd      = "do("
	dontCmd    = "don't("
	mulCmd     = "mul"
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

func Part1(filename string) (int64, error) {
	processor, err := processFile(filename, false)
	if err != nil {
		return 0, fmt.Errorf("processing file: %w", err)
	}
	return processor.sumOfMultiplications(), nil
}

func Part2(filename string) (int64, error) {
	processor, err := processFile(filename, true)
	if err != nil {
		return 0, fmt.Errorf("processing file: %w", err)
	}
	return processor.sumOfMultiplications(), nil
}

func processFile(filename string, strict bool) (*FileProcessor, error) {
	processor := NewFileProcessor()
	pairs, err := parseNumberPairs(filename, strict)
	if err != nil {
		return nil, fmt.Errorf("reading number pairs: %w", err)
	}

	processor.pairs = pairs
	return processor, nil
}

func parseNumberPairs(filename string, strict bool) ([]NumberPair, error) {

	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}

	input := string(content)

	pattern := mulPattern
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(input, -1)

	pairs := make([]NumberPair, 0, len(matches))
	canAppend := true

	for i, match := range matches {
		switch {
		case strings.HasPrefix(match[0], doCmd):
			canAppend = true
			continue
		case strings.HasPrefix(match[0], dontCmd):
			canAppend = false
			continue
		case !strings.HasPrefix(match[0], mulCmd):
			continue
		}

		if strict && !canAppend {
			continue
		}

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

	return pairs, nil
}

func (fp *FileProcessor) sumOfMultiplications() int64 {
	var totalSum int64
	for _, pair := range fp.pairs {
		totalSum = totalSum + (pair.First * pair.Second)
	}

	return totalSum
}
