package lcd

// Display is a uint that can be converted into a string representation
type Display uint

// Dimension defines the width and height of each digit as it's printed
type Dimension struct {
	Width  uint
	Height uint
}

// SPrint converts a Display to a string representation using LCD-style digits
func (l Display) SPrint(dimension Dimension) string {
	digits := digitize(uint(l))
	output := ""
	for row := 0; row < int(dimension.Height); row++ {
		output += printRow(digits, dimension, row)
	}
	return output
}

func digitize(number uint) []uint {
	if number == 0 {
		return []uint{0}
	}
	var digits []uint
	for number > 0 {
		digits = append(digits, number%10)
		number /= 10
	}
	return reverse(digits)
}

func reverse(digits []uint) []uint {
	numDigits := len(digits)
	for i := 0; i < numDigits/2; i++ {
		opp := numDigits - 1 - i
		digits[i], digits[opp] = digits[opp], digits[i]
	}
	return digits
}

func printRow(digits []uint, dimension Dimension, row int) string {
	output := ""
	srcRow := dstToSrc(row, dimension.Height, digitHeight)
	for _, digit := range digits {
		for col := 0; col < int(dimension.Width); col++ {
			srcCol := (srcRow * digitWidth) + dstToSrc(col, dimension.Width, digitWidth)
			output += string(digitStrings[digit][srcCol])
		}
	}
	return output + "\n"
}

func dstToSrc(dstValue int, dstSize, srcSize uint) int {
	dstMidpoint := int(dstSize / 2)
	srcMidpoint := int(srcSize / 2)
	switch {
	case dstValue == 0:
		return 0
	case dstValue == int(dstSize)-1:
		return int(srcSize) - 1
	case dstValue == dstMidpoint:
		return srcMidpoint
	case dstValue < dstMidpoint:
		return srcMidpoint - 1
	case dstValue > dstMidpoint:
		return srcMidpoint + 1
	}
	return 0
}

const digitWidth = 5
const digitHeight = 5

var digitStrings = [...]string{
	" ___ |   ||   ||   ||___|", // 0
	"         |    |    |    |", // 1
	" ___     | ___||    |___ ", // 2
	" ___     | ___|    | ___|", // 3
	"     |   ||___|    |    |", // 4
	" ___ |    |___     | ___|", // 5
	" ___ |    |___ |   ||___|", // 6
	" ___     |    |    |    |", // 7
	" ___ |   ||___||   ||___|", // 8
	" ___ |   ||___|    | ___|", // 9
}
