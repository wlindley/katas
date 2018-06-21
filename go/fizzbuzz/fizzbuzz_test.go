package fizzbuzz_test

import (
	"testing"

	"github.com/wlindley/katas/go/fizzbuzz"
)

func TestGenerate(t *testing.T) {
	tt := []struct {
		Name     string
		Range    []int
		Expected string
	}{
		{Name: "Three", Range: []int{1, 3}, Expected: "1\n2\nFizz\n"},
		{Name: "Five", Range: []int{4, 5}, Expected: "4\nBuzz\n"},
		{Name: "Fifteen", Range: []int{10, 15}, Expected: "Buzz\n11\nFizz\nFizz\n14\nFizzBuzz\n"},
		{Name: "Thirties", Range: []int{30, 35}, Expected: "FizzBuzz\nFizz\nFizz\nFizz\nFizz\nFizzBuzz\n"},
		{Name: "Fifties", Range: []int{50, 53}, Expected: "Buzz\nFizzBuzz\nBuzz\nFizzBuzz\n"},
	}
	for _, x := range tt {
		t.Run(x.Name, func(t *testing.T) {
			actual := fizzbuzz.Generate(x.Range[0], x.Range[1])
			if actual != x.Expected {
				t.Errorf("Expected something matching %s, but got %s", x.Expected, actual)
			}
		})
	}
}
