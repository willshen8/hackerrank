package appendDelete

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestAppendDelete(t *testing.T) {
	tests := []struct {
		testName string
		s        string
		t        string
		k        int32
		expected bool
	}{
		{
			testName: "Empty s and t",
			s:        "",
			t:        "",
			k:        4,
			expected: true,
		},
		{
			testName: "Delete initial string and append more",
			s:        "aaa",
			t:        "a",
			k:        5,
			expected: true,
		},
		{
			testName: "Simple delete and append",
			s:        "hackerhappy",
			t:        "hackerrank",
			k:        9,
			expected: true,
		},
		{
			testName: "Begin with small string and then append and delete",
			s:        "zzzzz",
			t:        "zzzzzzz",
			k:        4,
			expected: true,
		},
		{
			testName: "Different on first char",
			s:        "qwerty",
			t:        "zxcvbn",
			k:        100,
			expected: true,
		},
		{
			testName: "Begin and end with the same word",
			s:        "aba",
			t:        "aba",
			k:        7,
			expected: true,
		},
		{
			testName: "Not enough moves",
			s:        "ashley",
			t:        "ash",
			k:        2,
			expected: false,
		},
		{
			testName: "Same length but not enough moves",
			s:        "qwerasdf",
			t:        "qwerbsdf",
			k:        6,
			expected: false,
		},
		{
			testName: "Begin with less letters but not enough moves",
			s:        "y",
			t:        "yu",
			k:        2,
			expected: false,
		},
		{
			testName: "Same letters but different length",
			s:        "aaaaaaaaaa",
			t:        "aaaaa",
			k:        7,
			expected: true,
		},
		{
			testName: "Larger t",
			s:        "abcd",
			t:        "abcdert",
			k:        10,
			expected: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			actual := appendAndDelete(tc.s, tc.t, tc.k)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestfirstDifferentIndexPosition(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected int
	}{
		{
			s:        "",
			t:        "",
			expected: -1,
		},
		{
			s:        "hackerhappy",
			t:        "hackerrank",
			expected: 7,
		},
		{
			s:        "aba",
			t:        "aba",
			expected: -1,
		},
		{
			s:        "ashley",
			t:        "ash",
			expected: -1,
		},
		{
			s:        "qwerasdf",
			t:        "qwerbsdf",
			expected: 5,
		},
		{
			s:        "y",
			t:        "yu",
			expected: -1,
		},
		{
			s:        "aaaaaaaaaa",
			t:        "aaaaa",
			expected: -1,
		},
		{
			s:        "abcd",
			t:        "abcdert",
			expected: -1,
		},
	}
	for _, tc := range tests {
		actual := firstDifferentIndexPosition(tc.s, tc.t)
		assert.Equal(t, tc.expected, actual)
	}
}
