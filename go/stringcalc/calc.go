package stringcalc

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//Add parses a string into numbers and returns the sum
func Add(input string) (int, error) {
	if len(input) == 0 {
		return 0, nil
	}

	numerals := splitIntoNumerals(input)
	numbers, err := numeralsToInts(numerals)
	if err != nil {
		return 0, err
	}

	err = validateNumbers(numbers)
	if err != nil {
		return 0, err
	}

	return sum(numbers), nil
}

func splitIntoNumerals(input string) []string {
	delimeter, remainder := parseCustomDelimeter(input)
	return regexp.MustCompile(delimeter).Split(remainder, -1)
}

func parseCustomDelimeter(input string) (string, string) {
	if !strings.HasPrefix(input, "//") {
		return ",|\n", input
	}
	lineIndex := strings.Index(input, "\n")
	return input[2:lineIndex], input[lineIndex+1:]
}

func numeralsToInts(numerals []string) ([]int, error) {
	ints := make([]int, len(numerals))
	for i, n := range numerals {
		num, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		ints[i] = num
	}
	return ints, nil
}

func validateNumbers(numbers []int) error {
	errorMsg := "negatives not allowed"
	foundError := false
	for _, n := range numbers {
		if n < 0 {
			foundError = true
			errorMsg += fmt.Sprintf(" %d", n)
		}
	}
	if foundError {
		return errors.New(errorMsg)
	}
	return nil
}

func sum(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}
