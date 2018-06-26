package leapyear_test

import (
	"fmt"
	"testing"

	"github.com/wlindley/katas/go/leapyear"
)

func TestCheck(t *testing.T) {
	tt := []struct {
		Year   int
		IsLeap bool
	}{
		{Year: 600, IsLeap: false},
		{Year: 1700, IsLeap: false},
		{Year: 1800, IsLeap: false},
		{Year: 1900, IsLeap: false},
		{Year: 1999, IsLeap: false},
		{Year: 2000, IsLeap: true},
		{Year: 2002, IsLeap: false},
		{Year: 2008, IsLeap: true},
		{Year: 2012, IsLeap: true},
		{Year: 2016, IsLeap: true},
		{Year: 2017, IsLeap: false},
		{Year: 2018, IsLeap: false},
		{Year: 2019, IsLeap: false},
		{Year: 2100, IsLeap: false},
		{Year: 4000, IsLeap: false},
	}

	for _, x := range tt {
		t.Run(fmt.Sprintf("Year %d", x.Year), func(t *testing.T) {
			actual := leapyear.Check(x.Year)
			if actual != x.IsLeap {
				t.Errorf("for year %d got %v, but expected %v", x.Year, actual, x.IsLeap)
			}
		})
	}
}
