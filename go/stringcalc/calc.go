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
		if n <= 1000 {
			total += n
		}
	}
	return total
}

//ConcurrentAdd concurrently parses a string into numbers and returns the sum
func ConcurrentAdd(input string) (int, error) {
	if len(input) == 0 {
		return 0, nil
	}
	numerals := splitIntoNumerals(input)
	valid, invalid, errs := concurrentNumeralsToInts(numerals)

	total := 0
	errorFound := false
	errorMsg := "negatives not allowed"
	var error error
	for {
		select {
		case n, ok := <-valid:
			if !ok {
				valid = nil
			} else {
				total += n
			}
		case n, ok := <-invalid:
			if !ok {
				invalid = nil
			} else {
				errorFound = true
				errorMsg += fmt.Sprintf(" %d", n)
			}
		case err, ok := <-errs:
			if !ok {
				errs = nil
			} else if err != nil && error == nil {
				error = err
			}
		}

		if valid == nil && invalid == nil && errs == nil {
			break
		}
	}

	if error != nil {
		return 0, error
	}
	if errorFound {
		return 0, errors.New(errorMsg)
	}
	return total, nil
}

func concurrentNumeralsToInts(numerals []string) (<-chan int, <-chan int, <-chan error) {
	numeralChan := numeralsToChannel(numerals)
	numberChan, errs := numeralsToNumbers(numeralChan)
	smallNumberChan := ignoreLargeNumbers(numberChan)
	valid, invalid := validateAllNumbers(smallNumberChan)
	return valid, invalid, errs
}

func numeralsToChannel(numerals []string) <-chan string {
	out := make(chan string)
	go func() {
		for _, n := range numerals {
			out <- n
		}
		close(out)
	}()
	return out
}

func numeralsToNumbers(numerals <-chan string) (<-chan int, <-chan error) {
	out := make(chan int)
	errs := make(chan error)
	go func() {
		for n := range numerals {
			val, err := strconv.Atoi(n)
			if err != nil {
				errs <- err
			} else {
				out <- val
			}
		}
		close(out)
		close(errs)
	}()
	return out, errs
}

func ignoreLargeNumbers(numbers <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range numbers {
			if n <= 1000 {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

func validateAllNumbers(numbers <-chan int) (<-chan int, <-chan int) {
	valid := make(chan int)
	invalid := make(chan int)
	go func() {
		for n := range numbers {
			if n < 0 {
				invalid <- n
			} else {
				valid <- n
			}
		}
		close(valid)
		close(invalid)
	}()
	return valid, invalid
}
