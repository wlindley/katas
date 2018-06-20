package primefactors

//Generate returns prime factors for given number
func Generate(value int) []int {
	var factors []int
	for divisor := 2; value > 1; divisor++ {
		for ; value%divisor == 0; value /= divisor {
			factors = append(factors, divisor)
		}
	}
	return factors
}
