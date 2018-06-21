package fizzbuzz

import (
	"strconv"
	"strings"
)

//Generate returns a string containing all numbers up to n, each on their own
//line, but with some replaced based on different rules
func Generate(begin, end int) string {
	result := ""
	for i := begin; i <= end; i++ {
		isFizz := isOfNumber(i, 3)
		isBuzz := isOfNumber(i, 5)
		if isFizz {
			result += "Fizz"
		}
		if isBuzz {
			result += "Buzz"
		}
		if !isFizz && !isBuzz {
			result += strconv.Itoa(i)
		}
		result += "\n"
	}
	return result
}

func isOfNumber(value, number int) bool {
	divisible := value%number == 0
	contains := strings.Contains(strconv.Itoa(value), strconv.Itoa(number))
	return divisible || contains
}
