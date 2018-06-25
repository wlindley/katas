package lcd

// Display is a uint that can be converted into a string representation
type Display uint

// SPrint converts a Display to a string representation using LCD-style digits
func (l Display) SPrint() string {
	digits := digitize(uint(l))
	output := ""
	for row := 0; row < digitHeight; row++ {
		output += printRow(digits, row)
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

func printRow(digits []uint, row int) string {
	output := ""
	for _, digit := range digits {
		start := row * digitWidth
		end := start + digitWidth
		output += digitStrings[digit][start:end]
	}
	return output + "\n"
}

const digitWidth = 3
const digitHeight = 3

var digitStrings = [...]string{
	" _ | ||_|",
	"     |  |",
	" _  _||_ ",
	" _  _| _|",
	"   |_|  |",
	" _ |_  _|",
	" _ |_ |_|",
	" _   |  |",
	" _ |_||_|",
	" _ |_| _|",
}
