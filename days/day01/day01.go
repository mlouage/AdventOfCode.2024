package day01

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part1(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var totalSum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := extractNumbers(line)

		if len(numbers) > 0 {
			first := numbers[0]
			last := numbers[len(numbers)-1]
			combinedStr := fmt.Sprintf("%d%d", first, last)
			num, _ := strconv.Atoi(combinedStr)
			totalSum += num
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return totalSum, nil

}

func Part2(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	sum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		digits := extractDigits(line)

		if len(digits) > 0 {
			firstNum := digits[0]
			lastNum := digits[len(digits)-1]
			combinedNumStr := firstNum + lastNum

			if combinedNum, err := strconv.Atoi(combinedNumStr); err == nil {
				sum += combinedNum
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return sum, nil
}

func extractNumbers(s string) []int {
	re := regexp.MustCompile("[0-9]+")
	matches := re.FindAllString(s, -1)

	var numbers []int
	for _, match := range matches {
		num, _ := strconv.Atoi(match)
		numbers = append(numbers, num)
	}

	return numbers
}

var wordToDigit = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func extractDigits(line string) []string {
	var digits []string
	numberRegex := regexp.MustCompile(`\d+|zero|one|two|three|four|five|six|seven|eight|nine`)

	matches := numberRegex.FindAllString(line, -1)
	for _, match := range matches {
		if digit, ok := wordToDigit[match]; ok {
			digits = append(digits, digit)
		} else if isNumeric(match) {
			digits = append(digits, match)
		}
	}
	return digits
}
