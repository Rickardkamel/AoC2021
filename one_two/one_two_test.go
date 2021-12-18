package one_two

import "testing"

type testInput struct {
	name           string
	data           []int
	expectedResult int
}

func setInputs() []testInput {
	inputs := []testInput{
		{
			name:           "test_one_increment",
			data:           []int{0, 0, 0, 0, 1},
			expectedResult: 1,
		},
		{
			name:           "test_three_increments",
			data:           []int{0, 1, 2, 3, 4, 5},
			expectedResult: 3,
		},
		{
			name:           "test_four_increments",
			data:           []int{0, 1, 2, 3, 4, 5, 6},
			expectedResult: 4,
		},
		{
			name:           "test_two_increments",
			data:           []int{0, 1, 2, 3, 4, 1, 0},
			expectedResult: 2,
		},
	}
	return inputs
}

func TestCountIncreasingSums(t *testing.T) {
	inputs := setInputs()
	for i := range inputs {
		t.Run(inputs[i].name, func(t *testing.T) {
			result := CountIncreasingSums(inputs[i].data)
			if inputs[i].expectedResult != result {
				t.Errorf("Increased sums were %d; expected %d", result, inputs[i].expectedResult)
			}
		})
	}
}