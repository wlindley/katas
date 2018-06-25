package lcd_test

import (
	"testing"

	"github.com/wlindley/katas/go/lcd"
)

func TestSPrint(t *testing.T) {
	tt := []struct {
		Name     string
		Input    uint
		Expected string
	}{
		{Name: "Zero", Input: 0, Expected: " _ \n| |\n|_|\n"},
		{Name: "One", Input: 1, Expected: "   \n  |\n  |\n"},
		{Name: "Two", Input: 2, Expected: " _ \n _|\n|_ \n"},
		{Name: "Three", Input: 3, Expected: " _ \n _|\n _|\n"},
		{Name: "Four", Input: 4, Expected: "   \n|_|\n  |\n"},
		{Name: "Five", Input: 5, Expected: " _ \n|_ \n _|\n"},
		{Name: "Six", Input: 6, Expected: " _ \n|_ \n|_|\n"},
		{Name: "Seven", Input: 7, Expected: " _ \n  |\n  |\n"},
		{Name: "Eight", Input: 8, Expected: " _ \n|_|\n|_|\n"},
		{Name: "Nine", Input: 9, Expected: " _ \n|_|\n _|\n"},
		{Name: "Ten", Input: 10, Expected: "    _ \n  || |\n  ||_|\n"},
		{Name: "Large Number", Input: 1429, Expected: "       _  _ \n  ||_| _||_|\n  |  ||_  _|\n"},
	}

	for _, x := range tt {
		t.Run(x.Name, func(t *testing.T) {
			actual := lcd.Display(x.Input).SPrint()
			if actual != x.Expected {
				t.Errorf("for %d got\n%s, but expected\n%s", x.Input, actual, x.Expected)
			}
		})
	}
}
