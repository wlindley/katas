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
	}

	for _, x := range tt {
		t.Run(x.Name, func(t *testing.T) {
			actual := harrypotter.CalculatePrice(x.booksForTest())
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
}

func testCase(name string, books []int, expected float64) calculateTestCase {
	return calculateTestCase{
		Name:     name,
		Books:    books,
		Expected: expected,
	}
}

func (tc *calculateTestCase) booksForTest() []int {
	bookCounts := make([]int, len(tc.Books))
	copy(bookCounts, tc.Books)
	return bookCounts
}
