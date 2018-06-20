package primefactors_test

import (
	"testing"

	"github.com/wlindley/katas/go/primefactors"
)

func TestGenerate(t *testing.T) {
	tt := []struct {
		Name     string
		Input    int
		Expected []int
	}{
		{Name: "Zero returns empty list", Input: 0, Expected: []int{}},
		{Name: "One returns empty list", Input: 1, Expected: []int{}},
		{Name: "Two returns 2", Input: 2, Expected: []int{2}},
		{Name: "Three returns 3", Input: 3, Expected: []int{3}},
		{Name: "Four returns 2,2", Input: 4, Expected: []int{2, 2}},
		{Name: "Six returns 2,3", Input: 6, Expected: []int{2, 3}},
		{Name: "Eight returns 2,2,2", Input: 8, Expected: []int{2, 2, 2}},
		{Name: "Nine returns 3,3", Input: 9, Expected: []int{3, 3}},
	}

	for _, x := range tt {
		t.Run(x.Name, func(st *testing.T) {
			actual := primefactors.Generate(x.Input)
			if !areEqual(actual, x.Expected) {
				st.Errorf("Expected %d to factor to %v, but instead got %v", x.Input, x.Expected, actual)
			}
		})
	}
}

func areEqual(one, two []int) bool {
	if len(one) != len(two) {
		return false
	}
	for i, _ := range one {
		if one[i] != two[i] {
			return false
		}
	}
	return true
}
