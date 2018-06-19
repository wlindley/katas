package stringcalc_test

import (
	"regexp"
	"testing"

	"github.com/wlindley/katas/go/stringcalc"
)

func TestAdd(t *testing.T) {
	tt := []struct {
		Name     string
		Input    string
		Expected int
		Error    string
	}{
		{Name: "Empty String", Input: "", Expected: 0},
		{Name: "One", Input: "1", Expected: 1},
		{Name: "Single Invalid", Input: "a", Error: ".*"},
		{Name: "Two Numbers", Input: "1,2", Expected: 3},
		{Name: "Many", Input: "1,1,3,5,2,6", Expected: 18},
		{Name: "Newline", Input: "1\n2,3", Expected: 6},
		{Name: "Custom Separator", Input: "//;\n1;1;1", Expected: 3},
		{Name: "Reject Negatives", Input: "2,3,-1", Error: "negatives not allowed -1"},
		{Name: "Reject Multiple Negatives", Input: "-1,3,-2", Error: "negatives not allowed -1 -2"},
		{Name: "Ignore Numbers Over 1000", Input: "2,1001", Expected: 2},
	}

	for _, x := range tt {
		t.Run(x.Name, func(st *testing.T) {
			actual, err := stringcalc.Add(x.Input)

			if x.Error != "" {
				if err == nil {
					st.Error(`Expected an error, but did not receive one`)
				} else if matches, matchErr := regexp.MatchString(x.Error, err.Error()); matchErr != nil || !matches {
					st.Errorf(`Expected an error matching "%s", but received "%s"`, x.Error, err.Error())
				}
			}

			if x.Error == "" && err != nil {
				st.Errorf(`Encountered unexpected error: %s`, err.Error())
			}

			if actual != x.Expected {
				st.Errorf(`Expected "%s" to be %d, but was %d instead`, x.Input, x.Expected, actual)
			}
		})
	}
}

func TestConcurrentAdd(t *testing.T) {
	tt := []struct {
		Name     string
		Input    string
		Expected int
		Error    string
	}{
		{Name: "Empty String", Input: "", Expected: 0},
		{Name: "One", Input: "1", Expected: 1},
		{Name: "Single Invalid", Input: "a", Error: ".*"},
		{Name: "Two Numbers", Input: "1,2", Expected: 3},
		{Name: "Many", Input: "1,1,3,5,2,6", Expected: 18},
		{Name: "Newline", Input: "1\n2,3", Expected: 6},
		{Name: "Custom Separator", Input: "//;\n1;1;1", Expected: 3},
		{Name: "Reject Negatives", Input: "2,3,-1", Error: "negatives not allowed -1"},
		{Name: "Reject Multiple Negatives", Input: "-1,3,-2", Error: "negatives not allowed -1 -2"},
		{Name: "Ignore Numbers Over 1000", Input: "2,1001", Expected: 2},
	}

	for _, x := range tt {
		t.Run(x.Name, func(st *testing.T) {
			actual, err := stringcalc.ConcurrentAdd(x.Input)

			if x.Error != "" {
				if err == nil {
					st.Error(`Expected an error, but did not receive one`)
				} else if matches, matchErr := regexp.MatchString(x.Error, err.Error()); matchErr != nil || !matches {
					st.Errorf(`Expected an error matching "%s", but received "%s"`, x.Error, err.Error())
				}
			}

			if x.Error == "" && err != nil {
				st.Errorf(`Encountered unexpected error: %s`, err.Error())
			}

			if actual != x.Expected {
				st.Errorf(`Expected "%s" to be %d, but was %d instead`, x.Input, x.Expected, actual)
			}
		})
	}
}
