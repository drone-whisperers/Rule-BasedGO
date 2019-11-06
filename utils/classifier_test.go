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
			name:        "test should work",
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
		// {
		// 	name:        "test should work",
		// 	inputString: "big jet 345, metro ground taxi to holding point c runway 27",
		// 	expectedResult: map[string]interface{}{
		// 		"big jet 345":  "drone",
		// 		"metro ground": "targer",
		// 		"taxi": Taxi{
		// 			holdingPoint: "c",
		// 			holdPostion:  "",
		// 			runway: Runway{
		// 				number: 27,
		// 			},
		// 		},
		// 	},
		// },
		// {
		// 	name:        "test should work",
		// 	inputString: "big jet 345 contact metro tower 119.2",
		// 	expectedResult: map[string]interface{}{
		// 		"big jet 345":  "drone",
		// 		"metro ground": "targer",
		// 		"taxi": Taxi{
		// 			holdingPoint: "c",
		// 			holdPostion:  "",
		// 			runway: Runway{
		// 				number: 27,
		// 			},
		// 		},
		// 	},
		// },
	}
	lexer, err := lexer.InitLexer()
	require.NoError(t, err)
	c := NewClassifier(lexer)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s, _ := c.Classify(test.inputString)
			require.Equal(t, test.expectedResult, s)
		})
	}
}
