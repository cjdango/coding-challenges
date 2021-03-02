package minsteps

// CountMemo is a memoized function that
// returns the  minimum steps to minimize
// `number` to 1 but a lot faster than `Count`
func CountMemo(number int, memo []int) int {
	if number == 1 {
		return 0
	}

	if memo[number] != 0 {
		return memo[number]
	}

	result := CountMemo(number-1, memo)

	if number%2 == 0 {
		result = minOf(result, CountMemo(number/2, memo))
	}

	if number%3 == 0 {
		result = minOf(result, CountMemo(number/3, memo))
	}

	memo[number] = result + 1

	return memo[number]
}
