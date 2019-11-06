package utils

import (
	"testing"

	"github.com/Rule-BasedGO/lexer"
	"github.com/Rule-BasedGO/structures"
	"github.com/stretchr/testify/require"
)

type testClassifyStruct struct {
	name           string
	inputString    string
	expectedResult []interface{}
}

func TestClassify(t *testing.T) {
	tests := []testClassifyStruct{
		{
			name:        "Taxi Request",
			inputString: "big jet 345 metro ground taxi to holding point a 1 hold short of runway 18",
			expectedResult: []interface{}{
				&structures.Taxi{
					HoldingPoint: &structures.HoldingPoint{
						Location: "a 1",
					},
					RunWay: &structures.RunWay{
						Number: 18,
					},
					HoldPoint: "hold short of",
				},
			},
		},
		{
			name:        "Contact Request",
			inputString: "big jet 345 contact metro ground 119.2",
			expectedResult: []interface{}{
				&structures.Contact{
					Frequency: 119.2,
					Target:    "metro ground",
				},
			},
		},
	}
	lexer, err := lexer.InitLexer()
	require.NoError(t, err)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := NewClassifier(lexer)
			s, _ := c.Classify(test.inputString)
			require.Equal(t, test.expectedResult, s)
		})
	}
}
