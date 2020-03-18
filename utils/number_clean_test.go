package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testNewSlice struct {
	name           string
	originalS      []string
	s              []slicedInt
	expectedResult string
}

func BenchmarkNewSlice(t *testing.B) {
	tests := []testNewSlice{
		{
			name:      "test should work",
			originalS: []string{"this", "is", "test", "one", "hundread", "replace"},
			s: []slicedInt{slicedInt{
				start: 3,
				end:   4,
				num:   "100",
			}},
			expectedResult: "this is test 100 replace",
		},
		{
			name:      "test should work with decimal",
			originalS: []string{"this", "is", "test", "one", "hundread", "decimal", "one", "replace"},
			s: []slicedInt{slicedInt{
				start: 3,
				end:   4,
				num:   "100",
			},
				slicedInt{
					start: 6,
					end:   6,
					num:   "1",
				},
			},
			expectedResult: "this is test 100.1 replace",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.B) {
			s := newSlice(test.originalS, test.s)
			require.Equal(t, test.expectedResult, s)
		})
	}
}

type testIntSliceToNum struct {
	name           string
	slice          []int
	expectedResult string
}

func BenchmarkIntSliceToNum(t *testing.B) {
	tests := []testIntSliceToNum{
		{
			name:           "test should work",
			slice:          []int{1, 100},
			expectedResult: "100",
		},
		{
			name:           "test should work",
			slice:          []int{90, 1, 1000, 3, 100, 20, 3},
			expectedResult: "91323",
		},
		{
			name:           "test should work",
			slice:          []int{90, 1, 1000, 3, 100, 2, 3},
			expectedResult: "91323",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.B) {
			s := intSliceToNum(test.slice)
			require.Equal(t, test.expectedResult, s)
		})
	}
}

type testWordToNum struct {
	name           string
	s              string
	expectedResult string
}

func BenchmarkWordToNum(t *testing.B) {
	tests := []testWordToNum{
		{
			name:           "forty five thousand three hundred ninety five",
			s:              "forty five thousand three hundred ninety five",
			expectedResult: "45395",
		},
		{
			name:           "forty five thousand three hundred five five",
			s:              "forty five thousand three hundred five five",
			expectedResult: "45355",
		},
		{
			name:           "four five thousand three hundred five five",
			s:              "four five thousand three hundred five five",
			expectedResult: "45355",
		},
		{
			name:           "forty five thousand three five five",
			s:              "forty five thousand three five five",
			expectedResult: "45355",
		},
		{
			name:           "one one nine point five three",
			s:              "one one nine point five three",
			expectedResult: "119.53",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.B) {
			s := WordToNum(test.s)
			require.Equal(t, test.expectedResult, s)
		})
	}
}
