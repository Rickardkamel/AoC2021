package one

// CheckPreviousMeasurements goes through a slice of depth data points, and counts how many times we found depth point higher than its previous point.
func CheckPreviousMesurement(input []int) int {
	var counter int
	for i, k := range input {
		if i > 0 && k > input[i-1] {
			counter++
		}
	}
	return counter
}
