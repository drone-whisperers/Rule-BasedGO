package utils

import (
	"strings"
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
		{
			name:        "Cross Action with double action",
			inputString: "Big Jet 345 cross runway 18 at A 1 taxi to holding point C runway 27",
			expectedResult: []interface{}{
				&structures.Cross{
					RunWay: &structures.RunWay{
						Number: 18,
					},
					Location: "a 1",
				},
				&structures.Taxi{
					HoldingPoint: &structures.HoldingPoint{
						Location: "c",
					},
					RunWay: &structures.RunWay{
						Number: 27,
					},
				},
			},
		},
		{
			name:        "Start up Action with Contact request",
			inputString: "Big Jet 345 start up approved contact Metro Ground 118.75 for taxi instructions",
			expectedResult: []interface{}{
				&structures.StartUp{
					Condition: &structures.Condition{
						RequirementsMet: true,
					},
				},
				&structures.Contact{
					Frequency: 118.75,
					Target:    "metro ground",
					Request:   []string{"taxi instructions"},
				},
			},
		},
		{
			name:        "PRE POST ACTIONS",
			inputString: "Big Jet 345 after landing Airbus 321 cross Runway 19 at C 2 after",
			expectedResult: []interface{}{
				&structures.Condition{
					PreCondition: &structures.Landing{
						Plane: "airbus 321",
					},
					PostCondition: &structures.Cross{
						RunWay: &structures.RunWay{
							Number: 19,
						},
						Location: "c 2",
					},
				},
			},
		},
		{
			name:        "LineUp ACTIONS",
			inputString: "Big Jet 345 Metro ground line up runway 27",
			expectedResult: []interface{}{
				&structures.LineUp{
					RunWay: &structures.RunWay{
						Number: 27,
					},
				},
			},
		},
		{
			name:        "RunWay at Top level",
			inputString: "Big Jet 345 runway 27 cleared for take off",
			expectedResult: []interface{}{
				&structures.RunWay{
					Number: 27,
					Condition: &structures.Condition{
						RequirementsMet: true,
					},
					TakeOff: true,
				},
			},
		},
		{
			name:        "HOLD At top level",
			inputString: "Big Jet 345 Metro Tower hold at C 1",
			expectedResult: []interface{}{
				&structures.HoldingPoint{
					Location: "c 1",
				},
			},
		},
		{
			name:        "PRE POST ACTIONS with departure and climb",
			inputString: "Big Jet 345 hold position after departure climb to altitude 6000 feet",
			expectedResult: []interface{}{
				&structures.HoldingPoint{
					CurrentPosition: true,
				},
				&structures.Condition{
					PreCondition: &structures.Departure{},
					PostCondition: &structures.Climb{
						Altitude: &structures.Altitude{
							Number: 6000,
							Unit:   "feet",
						},
					},
				},
			},
		},
		{
			name:        "PRE POST ACTIONS with behind",
			inputString: "Big Jet 345 behind landing Boeing 757 line up runway 27 behind",
			expectedResult: []interface{}{
				&structures.Condition{
					PreCondition: &structures.Landing{
						Plane: "boeing 757",
					},
					PostCondition: &structures.LineUp{
						RunWay: &structures.RunWay{
							Number: 27,
						},
					},
				},
			},
		},
		{
			name:        "Hold Current Position With canceled Action",
			inputString: "Big Jet 345 hold position Cancel take off I say again cancel take off due to vehicle on the runway",
			expectedResult: []interface{}{
				&structures.HoldingPoint{
					CurrentPosition: true,
				},
				&structures.Condition{
					Canceled:     true,
					PreCondition: &structures.TakeOff{},
				},
				&structures.Condition{
					Canceled:     true,
					PreCondition: &structures.TakeOff{},
				},
			},
		},
		{
			name:        "Immediate Action",
			inputString: "Big Jet 345 stop immediately Big Jet 345 stop immediately",
			expectedResult: []interface{}{
				&structures.Stop{
					Condition: &structures.Condition{
						ImmediateAction: true,
					},
				},
				&structures.Stop{
					Condition: &structures.Condition{
						ImmediateAction: true,
					},
				},
			},
		},
		{
			name:        "Fly Action with climb action with flightlevel",
			inputString: "Big Jet 345 fly heading 260 degrees climb to FL 100 no speed restrictions",
			expectedResult: []interface{}{
				&structures.Fly{
					Heading: &structures.Heading{
						Number: 260,
						Unit:   "degrees",
					},
				},
				&structures.Climb{
					Altitude: &structures.Altitude{
						FlightLevel: &structures.FlightLevel{
							Level:          100,
							AdditionalInfo: []string{"no speed restrictions"},
						},
					},
				},
			},
		},
		{
			name:        "fly with climb action",
			inputString: "Big Jet 345 fly direct BONNY climb to FL 360",
			expectedResult: []interface{}{
				&structures.Fly{
					Locations: []string{"bonny"},
				},
				&structures.Climb{
					Altitude: &structures.Altitude{
						FlightLevel: &structures.FlightLevel{
							Level: 360,
						},
					},
				},
			},
		},
		{
			name:        "Fly with single location",
			inputString: "Big Jet 345 Northern Control fly direct CLYDE",
			expectedResult: []interface{}{
				&structures.Fly{
					Locations: []string{"clyde"},
				},
			},
		},
		{
			name:        "turn action with avoidance",
			inputString: "Big Jet 345 turn left immediately heading 270 to avoid traffic at 2 o'clock 5 miles crossing right to left 500 feet below",
			expectedResult: []interface{}{
				&structures.Turn{
					Direction: "left",
					Heading: &structures.Heading{
						Number: 270,
					},
					Condition: &structures.Condition{
						ImmediateAction: true,
					},
				},
				&structures.Avoid{
					Object: &structures.Traffic{
						Direction: &structures.Heading{
							Number: 2,
							Unit:   "o'clock",
						},
						Distance: &structures.Distance{
							Number: 5,
							Unit:   "miles",
						},
						Action: &structures.Crossing{
							From:   "right",
							To:     "left",
							Unit:   "feet",
							Number: 500,
							Info:   []string{"below"},
						},
					},
				},
			},
		},
		{
			name:        "turn action with avoidance",
			inputString: "Big Jet 345 turn right immediately heading 30 degrees to avoid traffic at 2 o'clock 5 miles crossing right to left 500 feet below",
			expectedResult: []interface{}{
				&structures.Turn{
					Direction: "right",
					Heading: &structures.Heading{
						Number: 30,
						Unit:   "degrees",
					},
					Condition: &structures.Condition{
						ImmediateAction: true,
					},
				},
				&structures.Avoid{
					Object: &structures.Traffic{
						Direction: &structures.Heading{
							Number: 2,
							Unit:   "o'clock",
						},
						Distance: &structures.Distance{
							Number: 5,
							Unit:   "miles",
						},
						Action: &structures.Crossing{
							From:   "right",
							To:     "left",
							Unit:   "feet",
							Number: 500,
							Info:   []string{"below"},
						},
					},
				},
			},
		},
		{
			name:        "climb action with traffic",
			inputString: "Big Jet 345 climb immediately to FL 160 traffic at 12 o'clock 3 miles opposite direction same level",
			expectedResult: []interface{}{
				&structures.Climb{
					Altitude: &structures.Altitude{
						FlightLevel: &structures.FlightLevel{
							Level: 160,
						},
					},
					Condition: &structures.Condition{
						ImmediateAction: true,
					},
				},
				&structures.Traffic{
					Direction: &structures.Heading{
						Number: 12,
						Unit:   "o'clock",
					},
					Distance: &structures.Distance{
						Number: 3,
						Unit:   "miles",
					},
					Info: []string{"opposite", "direction", "same", "level"},
				},
			},
		},
		{
			name:        "decend with traffic",
			inputString: "Big Jet 345 descend immediately to FL 160 traffic at 12 o'clock 3 miles opposite direction same level",
			expectedResult: []interface{}{
				&structures.Descend{
					Altitude: &structures.Altitude{
						FlightLevel: &structures.FlightLevel{
							Level: 160,
						},
					},
					Condition: &structures.Condition{
						ImmediateAction: true,
					},
				},
				&structures.Traffic{
					Direction: &structures.Heading{
						Number: 12,
						Unit:   "o'clock",
					},
					Distance: &structures.Distance{
						Number: 3,
						Unit:   "miles",
					},
					Info: []string{"opposite", "direction", "same", "level"},
				},
			},
		},
		{
			name:        "QNH Example",
			inputString: "Big Jet 345 Metro Approach new information Q new QNH 998",
			expectedResult: []interface{}{
				&structures.QNH{
					Pressure: 998,
				},
			},
		},
		{
			name:        "Leave action with decend",
			inputString: "Big Jet 345 leave MAYFIELD heading 120 descend to 6000 feet QNH 998 speed 210 knots",
			expectedResult: []interface{}{
				&structures.Departure{
					Location: []string{"mayfield"},
					Heading: &structures.Heading{
						Number: 120,
					},
				},
				&structures.Descend{
					Altitude: &structures.Altitude{
						Number: 6000,
						Unit:   "feet",
					},
					QNH: &structures.QNH{
						Pressure: 998,
					},
					Speed: &structures.Speed{
						Number: 210,
						Unit:   "knots",
					},
				},
			},
		},
		{
			name:        "Turn Action with ILS",
			inputString: "Big Jet 345 turn right heading 180 speed 180 knots vectoring ILS runway 27 Right",
			expectedResult: []interface{}{
				&structures.Turn{
					Direction: "right",
					Heading: &structures.Heading{
						Number: 180,
					},
					Speed: &structures.Speed{
						Number: 180,
						Unit:   "knots",
					},
				},
				&structures.ILS{
					RunWay: &structures.RunWay{
						Number: 27,
					},
					Direction: "right",
				},
			},
		},
		{
			name:        "Turn With descend Action with ILS",
			inputString: "Big Jet 345 turn right heading 240 descend to 3000 feet report established localiser runway 27 Right",
			expectedResult: []interface{}{
				&structures.Turn{
					Direction: "right",
					Heading: &structures.Heading{
						Number: 240,
					},
				},
				&structures.Descend{
					Altitude: &structures.Altitude{
						Number: 3000,
						Unit:   "feet",
					},
				},
				&structures.ILS{
					RunWay: &structures.RunWay{
						Number: 27,
					},
					Direction: "right",
				},
			},
		},
		{
			name:        "Cleared For ILS action",
			inputString: "Big Jet 345 cleared ILS approach runway 27 Right",
			expectedResult: []interface{}{
				&structures.Condition{
					RequirementsMet: true,
				},
				&structures.ILS{
					RunWay: &structures.RunWay{
						Number: 27,
					},
					Direction: "right",
				},
			},
		},
		{
			name:        "TURN With Cleared ILS ACTION With Maintain Action",
			inputString: "Big Jet 345 turn right heading 240 degrees cleared ILS approach runway 27 Right maintain 3000 feet until glide path interception",
			expectedResult: []interface{}{
				&structures.Turn{
					Direction: "right",
					Heading: &structures.Heading{
						Number: 240,
						Unit:   "degrees",
					},
				},
				&structures.Condition{
					RequirementsMet: true,
				},
				&structures.ILS{
					RunWay: &structures.RunWay{
						Number: 27,
					},
					Direction: "right",
				},
				&structures.Maintain{
					Altitude: &structures.Altitude{
						Number: 3000,
						Unit:   "feet",
					},
					Condition: &structures.Condition{
						PreCondition: &structures.GlidePathInterception{},
					},
				},
			},
		},
		{
			name:        "Continue Action Example",
			inputString: "Big Jet 345 continue approach",
			expectedResult: []interface{}{
				&structures.Condition{
					RequirementsMet: true,
					PreCondition:    &structures.Approach{},
				},
			},
		},
		{
			name:        "Cleared to land, land action with wind",
			inputString: "Big Jet 345 cleared to land runway 27 Right wind 270 degrees 10 knots",
			expectedResult: []interface{}{
				&structures.Condition{
					RequirementsMet: true,
				},
				&structures.Land{
					RunWay: &structures.RunWay{
						Number: 27,
					},
					Direction: "right",
					Wind: &structures.Wind{
						Heading: &structures.Heading{
							Number: 270,
							Unit:   "degrees",
						},
						Speed: &structures.Speed{
							Number: 10,
							Unit:   "knots",
						},
					},
				},
			},
		},
		{
			name:        "go around action",
			inputString: "Big Jet 345 go around",
			expectedResult: []interface{}{
				&structures.GoAround{},
			},
		},
		{
			name:        "Turn With RadarVectors",
			inputString: "Big Jet 345 Roger the MAYDAY turn left heading 90 radar vectors ILS runway 27",
			expectedResult: []interface{}{
				&structures.Turn{
					Direction: "left",
					Heading: &structures.Heading{
						Number: 90,
					},
				},
				&structures.RadarVectors{
					ILS: &structures.ILS{
						RunWay: &structures.RunWay{
							Number: 27,
						},
					},
				},
			},
		},
		{
			name:        "Turn With RadarVectors and decend actions",
			inputString: "Big Jet 345 roger turn right heading 140 for radar vectoring runway 9 descend to 3000 feet QNH 995 report established",
			expectedResult: []interface{}{
				&structures.Turn{
					Direction: "right",
					Heading: &structures.Heading{
						Number: 140,
					},
				},
				&structures.RadarVectors{
					RunWay: &structures.RunWay{
						Number: 9,
					},
				},
				&structures.Descend{
					Altitude: &structures.Altitude{
						Number: 3000,
						Unit:   "feet",
					},
					QNH: &structures.QNH{
						Pressure: 995,
					},
				},
			},
		},
		{
			name:        "Cleared for departure with squawk",
			inputString: "Big Jet 345 Metro ground Cleared to Smallville T 1 A departure Squawk 3456 slot time 1905",
			expectedResult: []interface{}{
				&structures.Condition{
					RequirementsMet: true,
					PreCondition: &structures.Departure{
						Location: []string{"smallville", "t 1 a"},
						Squawk: &structures.Squawk{
							Number: 3456,
							SlotTime: &structures.SlotTime{
								Number: 1905,
							},
						},
					},
				},
			},
		},
		{
			name:        "Cleared for departure with squawk",
			inputString: "Big Jet 345 hold position amendment to clearance T 3 F departure climb to 6000 feet",
			expectedResult: []interface{}{
				&structures.HoldingPoint{
					CurrentPosition: true,
				},
				&structures.Condition{
					Amendment:       true,
					RequirementsMet: true,
					PreCondition: &structures.Departure{
						Location: []string{"t 3 f"},
					},
				},
				&structures.Climb{
					Altitude: &structures.Altitude{
						Number: 6000,
						Unit:   "feet",
					},
				},
			},
		},
	}
	lexer, err := lexer.InitLexer()
	require.NoError(t, err)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := NewClassifier(lexer)
			s, _ := c.Classify(strings.ToLower(test.inputString))
			// s2, _ := json.MarshalIndent(s, "", "\t")
			// fmt.Println(string(s2))
			require.Equal(t, test.expectedResult, s)
		})
	}
}
