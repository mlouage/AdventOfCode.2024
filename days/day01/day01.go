package day01

import (
	"bufio"
	"bytes"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Part1(filename string) (int64, error) {
	numLines, err := countLines(filename)

	if err != nil {
		return 0, err
	}

	list1, list2, err := createLists(filename, numLines)

	if err != nil {
		return 0, err
	}

	sortLists(list1, list2)

	var totalSum int64

	for i, _ := range list1 {
		distance := list1[i] - list2[i]
		totalSum = totalSum + int64(math.Abs(float64(distance)))
	}

	return totalSum, nil

}

func Part2(filename string) (int64, error) {
	numLines, err := countLines(filename)

	if err != nil {
		return 0, err
	}

	list1, list2, err := createLists(filename, numLines)

	if err != nil {
		return 0, err
	}

	var totalSum int64

	for i, _ := range list1 {
		number := list1[i]
		count := count(
			list2,
			func(x int64) bool {
				return x == number
			})

		totalSum = totalSum + (number * count)
	}

	return totalSum, err
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

func sortLists(list1 []int64, list2 []int64) {
	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})

	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})
}

func createLists(filename string, numLines int) ([]int64, []int64, error) {
	var count int

	file, err := os.Open(filename)

	if err != nil {
		return []int64{}, []int64{}, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	column1 := make([]int64, numLines)
	column2 := make([]int64, numLines)

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)

		if len(numbers) != 2 {
			continue
		}

		num1, _ := strconv.ParseInt(numbers[0], 10, 64)
		num2, _ := strconv.ParseInt(numbers[1], 10, 64)

		column1[count] = num1
		column2[count] = num2

		count++
	}

	if err := scanner.Err(); err != nil {
		return []int64{}, []int64{}, err
	}

	return column1, column2, err
}

func countLines(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	counter := &JimBLineCounter{}

	numLines, err := counter.Count(file)

	if err != nil {
		return 0, err
	}

	return numLines, err
}

type JimBLineCounter struct {
	Size int    // Size of the buffer
	Sep  string // End of line character
}

func (b *JimBLineCounter) Count(r io.Reader) (int, error) {
	defaultSize := 32 * 1024
	defaultEndLine := "\n"

	if b.Size == 0 {
		b.Size = defaultSize
	}

	if b.Sep == "" {
		b.Sep = defaultEndLine
	}

	buf := make([]byte, b.Size)
	var count int

	for {
		n, err := r.Read(buf)
		count += bytes.Count(buf[:n], []byte(b.Sep))

		if err != nil {
			if err == io.EOF {
				return count, nil
			}
			return count, err
		}

	}
}
