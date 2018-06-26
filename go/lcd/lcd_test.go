package lcd_test

import (
	"testing"

	"github.com/wlindley/katas/go/lcd"
)

func TestSPrint(t *testing.T) {
	tt := []testCase{
		defaultTestCase("Zero", 0, " _ \n| |\n|_|\n"),
		defaultTestCase("One", 1, "   \n  |\n  |\n"),
		defaultTestCase("Two", 2, " _ \n _|\n|_ \n"),
		defaultTestCase("Three", 3, " _ \n _|\n _|\n"),
		defaultTestCase("Four", 4, "   \n|_|\n  |\n"),
		defaultTestCase("Five", 5, " _ \n|_ \n _|\n"),
		defaultTestCase("Six", 6, " _ \n|_ \n|_|\n"),
		defaultTestCase("Seven", 7, " _ \n  |\n  |\n"),
		defaultTestCase("Eight", 8, " _ \n|_|\n|_|\n"),
		defaultTestCase("Nine", 9, " _ \n|_|\n _|\n"),
		defaultTestCase("Ten", 10, "    _ \n  || |\n  ||_|\n"),
		defaultTestCase("Large Number", 1429, "       _  _ \n  ||_| _||_|\n  |  ||_  _|\n"),
		sizedTestCase("Big Zero", 0, lcd.Dimension{Width: 7, Height: 7}, " _____ \n|     |\n|     |\n|     |\n|     |\n|     |\n|_____|\n"),
		sizedTestCase("Non-square", 8, lcd.Dimension{Width: 3, Height: 5}, " _ \n| |\n|_|\n| |\n|_|\n"),
		sizedTestCase("Non-odd", 2, lcd.Dimension{Width: 6, Height: 4}, " ____ \n     |\n ____|\n|____ \n"),
	}

	for _, x := range tt {
		t.Run(x.Name, func(t *testing.T) {
			actual := lcd.Display(x.Input).SPrint(x.Dimension)
			if actual != x.Expected {
				t.Errorf("for %d got\n%s, but expected\n%s", x.Input, actual, x.Expected)
			}
		})
	}
}

type testCase struct {
	Name      string
	Input     uint
	Expected  string
	Dimension lcd.Dimension
}

func defaultTestCase(name string, input uint, expected string) testCase {
	return sizedTestCase(name, input, lcd.Dimension{Width: 3, Height: 3}, expected)
}

func sizedTestCase(name string, input uint, dimension lcd.Dimension, expected string) testCase {
	return testCase{
		Name:      name,
		Input:     input,
		Expected:  expected,
		Dimension: dimension,
	}
}
