package minsteps

// Count returns the  minimum steps
// to minimize `number` to 1
func Count(number int) int {
	if number == 1 {
		return 0
	}

	result := Count(number - 1)

	if number%2 == 0 {
		result = minOf(result, Count(number/2))
	}

	if number%3 == 0 {
		result = minOf(result, Count(number/3))
	}

	return result + 1
}
