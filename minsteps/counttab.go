package minsteps

// CountTab is a tabularized function that
// returns the  minimum steps to minimize
// `number` to 1 and a lot faster than `CountMemo`.
func CountTab(number int) int {
	table := make([]int, number+1)

	getMin := func(a, b int) int {
		if b > 0 {
			return minOf(a, b)
		}

		return a
	}

	for idx := 1; idx < number; idx++ {
		table[idx+1] = getMin(table[idx]+1, table[idx+1])

		if idx*2 <= number {
			table[idx*2] = getMin(table[idx]+1, table[idx*2])
		}

		if idx*3 <= number {
			table[idx*3] = getMin(table[idx]+1, table[idx*3])
		}
	}

	return table[number]
}
