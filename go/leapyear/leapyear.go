package leapyear

// Check returns true if the given year is a leap year in the Gregorian calendar and false otherwise
func Check(year int) bool {
	divBy4 := year%4 == 0
	divBy100 := year%100 == 0
	divBy400 := year%400 == 0
	divBy4000 := year%4000 == 0
	return !divBy4000 && divBy400 || (divBy4 && !divBy100)
}
