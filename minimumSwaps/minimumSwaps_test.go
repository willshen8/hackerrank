package minimumSwaps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumSwaps(t *testing.T) {
	tests := []struct {
		testName string
		input    []int32
		expected int32
	}{
		{
			testName: "Simple minimum swap",
			input:    []int32{2, 3, 4, 1, 5},
			expected: 3,
		},
		{
			testName: "Simple minimum swap 2",
			input:    []int32{1, 3, 5, 2, 4, 6, 7},
			expected: 3,
		},
	}
	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			actual := minimumSwaps(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestSwap(t *testing.T) {
	tests := []struct {
		testName string
		input    []int32
		indexA   int
		indexB   int
		expected []int32
	}{
		{
			testName: "Simple minimum swap",
			input:    []int32{2, 3},
			indexA:   0,
			indexB:   1,
			expected: []int32{3, 2},
		},
	}
	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			actual := swap(tc.input, tc.indexA, tc.indexB)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
