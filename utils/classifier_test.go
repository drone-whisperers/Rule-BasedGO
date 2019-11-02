package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testClassifyStruct struct {
	name           string
	inputString    string
	expectedResult map[string]interface{}
}

func TestClassify(t *testing.T) {
	tests := []testClassifyStruct{
		{
			name:        "test should work",
			inputString: "big jet 345 metro ground taxi to holding point a 1 hold short of runway 18",
			expectedResult: map[string]interface{}{
				"big jet 345":  "drone",
				"metro ground": "targer",
				"taxi": Taxi{
					holdingPoint: "a 1",
					holdPostion:  "hold short of",
					runway: Runway{
						number: 18,
					},
				},
			},
		},
		{
			name:        "test should work",
			inputString: "big jet 345, metro ground taxi to holding point c runway 27",
			expectedResult: map[string]interface{}{
				"big jet 345":  "drone",
				"metro ground": "targer",
				"taxi": Taxi{
					holdingPoint: "c",
					holdPostion:  "",
					runway: Runway{
						number: 27,
					},
				},
			},
		},
		{
			name:        "test should work",
			inputString: "big jet 345 contact metro tower 119.2",
			expectedResult: map[string]interface{}{
				"big jet 345":  "drone",
				"metro ground": "targer",
				"taxi": Taxi{
					holdingPoint: "c",
					holdPostion:  "",
					runway: Runway{
						number: 27,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := classify(test.inputString)
			require.Equal(t, test.expectedResult, s)
		})
	}
}
