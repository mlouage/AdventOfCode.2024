package day05

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PageRule struct {
	Before int
	After  int
}

func NewPageRule(before, after int) PageRule {
	return PageRule{Before: before, After: after}
}

type PageUpdate struct {
	Updates []int
}

func NewPageUpdate(updates []int) PageUpdate {
	return PageUpdate{Updates: updates}
}

func Part1(ctx context.Context, filename string) (int, error) {
	pageRules, pageUpdates, err := parseRulesAndUpdates(ctx, filename)
	if err != nil {
		return 0, fmt.Errorf("processing file: %w", err)
	}

	sum := 0
	for _, update := range pageUpdates {
		if isValidOrder(update.Updates, pageRules) {
			sum += getMiddleNumber(update.Updates)
		}
	}

	return sum, nil
}

func Part2(ctx context.Context, filename string) (int, error) {
	pageRules, pageUpdates, err := parseRulesAndUpdates(ctx, filename)
	if err != nil {
		return 0, fmt.Errorf("processing file: %w", err)
	}

	sum := 0
	for _, update := range pageUpdates {
		if !isValidOrder(update.Updates, pageRules) {
			orderedUpdate := topologicalSort(update.Updates, pageRules)
			sum += getMiddleNumber(orderedUpdate)
		}
	}

	return sum, nil
}

func parseRulesAndUpdates(ctx context.Context, filename string) ([]PageRule, []PageUpdate, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var pageRules []PageRule
	var pageUpdates []PageUpdate

	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return nil, nil, ctx.Err()
		default:
			line := scanner.Text()
			result, found, err := convertToArray(line, "|")
			if found && err == nil {
				pageRules = append(pageRules, NewPageRule(result[0], result[1]))
			}
			result, found, err = convertToArray(line, ",")
			if found && err == nil {
				pageUpdates = append(pageUpdates, NewPageUpdate(result))
			}
		}
	}

	return pageRules, pageUpdates, nil
}

func isValidOrder(pages []int, rules []PageRule) bool {
	positions := make(map[int]int)
	for i, page := range pages {
		positions[page] = i
	}

	for _, rule := range rules {
		beforePos, beforeExists := positions[rule.Before]
		afterPos, afterExists := positions[rule.After]

		if beforeExists && afterExists && beforePos > afterPos {
			return false
		}
	}

	return true
}

func getMiddleNumber(nums []int) int {
	return nums[len(nums)/2]
}

func convertToArray(input string, sep string) ([]int, bool, error) {
	strSlice := strings.Split(input, sep)

	if len(strSlice) == 0 {
		return nil, false, nil
	}

	intSlice := make([]int, 0, len(strSlice))

	for _, str := range strSlice {
		if num, err := strconv.Atoi(str); err == nil {
			intSlice = append(intSlice, num)
		} else {
			return nil, false, err
		}
	}

	return intSlice, true, nil
}

func buildDependencyGraph(pages []int, rules []PageRule) map[int][]int {
	// Create a map of pages that must come before each page
	dependencies := make(map[int][]int)

	// Initialize empty dependencies for all pages
	for _, page := range pages {
		if _, exists := dependencies[page]; !exists {
			dependencies[page] = []int{}
		}
	}

	// Add dependencies based on rules
	for _, rule := range rules {
		if containsPage(pages, rule.Before) && containsPage(pages, rule.After) {
			dependencies[rule.After] = append(dependencies[rule.After], rule.Before)
		}
	}

	return dependencies
}

func containsPage(pages []int, page int) bool {
	for _, p := range pages {
		if p == page {
			return true
		}
	}
	return false
}

func topologicalSort(pages []int, rules []PageRule) []int {
	dependencies := buildDependencyGraph(pages, rules)
	visited := make(map[int]bool)
	temp := make(map[int]bool)
	var order []int

	var visit func(page int)
	visit = func(page int) {
		if temp[page] {
			return // Skip if already in temporary mark (handles cycles)
		}
		if visited[page] {
			return // Skip if already visited
		}
		temp[page] = true

		// Visit dependencies
		for _, dep := range dependencies[page] {
			visit(dep)
		}

		temp[page] = false
		visited[page] = true
		order = append([]int{page}, order...)
	}

	// Visit all pages
	for _, page := range pages {
		if !visited[page] {
			visit(page)
		}
	}

	return order
}
