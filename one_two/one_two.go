package one_two

func CountIncreasingSums(input []int) int {
	start := 0
	stop := 2
	counter := 0
	var previousSum int
	for i := range input {
		start = i
		end := i + stop
		if end < len(input) {
			sum := input[start] + input[start+1] + input[end]

			if i != 0 && sum > previousSum {
				counter += 1
			}
			previousSum = sum
		}
	}

	return counter
}
