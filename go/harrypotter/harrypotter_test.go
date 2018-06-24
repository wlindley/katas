package harrypotter_test

import (
	"testing"

	"github.com/wlindley/katas/go/harrypotter"
)

func TestCalculatePrice(t *testing.T) {
	tt := []calculateTestCase{
		testCase("No Books", []int{}, 0),
		testCase("One Book", []int{1}, 8),
		testCase("Two Unique Books", []int{1, 1}, 16*.95),
		testCase("Three Unique Books", []int{1, 1, 1}, 24*.90),
		testCase("Four Unique Books", []int{1, 1, 1, 1}, 32*.80),
		testCase("Five Unique Books", []int{1, 1, 1, 1, 1}, 40*.75),
		testCase("Two Duplicate Books", []int{2}, 16),
		testCase("Zero Book Count", []int{0}, 0),
		testCase("Another Zero Book Count", []int{1, 0, 1}, 16*.95),
		testCase("Different Sized Sets", []int{2, 1, 2}, (24*.90)+(16*.95)),
		testCase("Two Sets of Four", []int{2, 2, 2, 1, 1}, 2*(32*.80)),
		errorCase("Too May Books", []int{1, 1, 1, 1, 1, 1}),
		errorCase("Negative Book Counts", []int{-1, -1}),
	}

	for _, x := range tt {
		t.Run(x.Name, func(t *testing.T) {
			basket, ok := x.basket()
			if !ok && !x.Error {
				t.Errorf("Invalid books for basket: %v\n", x.Books)
			} else if ok && x.Error {
				t.Errorf("Expected %v to produce invalid basket, but it did not", x.Books)
			}
			actual := harrypotter.CalculatePrice(basket)
			if actual != x.Expected {
				t.Errorf("Expected %v to cost $%.2f, but instead cost $%.2f", x.Books, x.Expected, actual)
			}
		})
	}
}

type calculateTestCase struct {
	Name     string
	Books    []int
	Expected float64
	Error    bool
}

func testCase(name string, books []int, expected float64) calculateTestCase {
	return calculateTestCase{
		Name:     name,
		Books:    books,
		Expected: expected,
	}
}

func errorCase(name string, books []int) calculateTestCase {
	return calculateTestCase{
		Name:  name,
		Books: books,
		Error: true,
	}
}

func (tc *calculateTestCase) basket() (harrypotter.BookBasket, bool) {
	bookCounts := make([]int, len(tc.Books))
	copy(bookCounts, tc.Books)
	return harrypotter.CreateBasket(bookCounts...)
}
