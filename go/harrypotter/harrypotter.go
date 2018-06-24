package harrypotter

//NumBooks is the number of books in the series
const NumBooks = 5
const pricePerBook = 8.0
const optimalDiscountSize = 4

var discounts = [NumBooks + 1]float64{1.0, 1.0, .95, .9, .8, .75}

//CalculatePrice computes and returns the price for the specified books
func CalculatePrice(books []int) float64 {
	sets := mergeComplementarySets(buildSets(books))
	return totalPriceOf(sets)
}

type bookSet []bool

func (set bookSet) price() float64 {
	numBooks := set.count()
	return float64(numBooks) * pricePerBook * discounts[numBooks]
}

func (set bookSet) count() int {
	count := 0
	for _, present := range set {
		if present {
			count++
		}
	}
	return count
}

func buildSets(books []int) []bookSet {
	sets := []bookSet{}
	for {
		set, ok := extractSet(books, optimalDiscountSize)
		if !ok {
			break
		}
		sets = append(sets, set)
	}
	return sets
}

func extractSet(books []int, maxSize int) (bookSet, bool) {
	counts := make([]bool, NumBooks)
	setSize := 0
	for i, count := range books {
		if count > 0 && setSize < maxSize {
			counts[i] = true
			books[i]--
			setSize++
		}
	}
	if setSize == 0 {
		return counts, false
	}
	return counts, true
}

func mergeComplementarySets(sets []bookSet) []bookSet {
	fours := findSetsOf(sets, optimalDiscountSize)
	ones := findSetsOf(sets, NumBooks-optimalDiscountSize)
	for _, fourSet := range fours {
		for _, oneSet := range ones {
			if areComplementary(fourSet, oneSet) {
				mergeSets(fourSet, oneSet)
			}
		}
	}
	return sets
}

func totalPriceOf(sets []bookSet) float64 {
	totalPrice := 0.0
	for _, set := range sets {
		totalPrice += set.price()
	}
	return totalPrice
}

func findSetsOf(sets []bookSet, size int) []bookSet {
	startIndex := 0
	stopIndex := len(sets)
	for i, set := range sets {
		if set.count() > size {
			startIndex = i + 1
		}
		if set.count() < size {
			stopIndex = i
			break
		}
	}
	return sets[startIndex:stopIndex]
}

func areComplementary(bigSet, smallSet bookSet) bool {
	for i := range bigSet {
		if bigSet[i] && smallSet[i] {
			return false
		}
	}
	return true
}

func mergeSets(bigSet, smallSet bookSet) {
	for i := range bigSet {
		if smallSet[i] {
			bigSet[i] = smallSet[i]
			smallSet[i] = false
		}
	}
}
