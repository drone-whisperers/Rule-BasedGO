package lexer

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	lex "github.com/timtadh/lexmachine"
)

type testLexer struct {
	name       string
	scanString string
}

func TestLexer(t *testing.T) {
	tests := []testLexer{
		{
			name:       "test should work",
			scanString: "big jet 345 metro ground taxi to holding point a 1 hold short of runway 18",
		},
		{
			name:       "test should work",
			scanString: "big jet 345 contact metro ground 119.2",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 Metro Ground taxi to holding point C runway 27",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 Metro Ground taxi to holding point C runway 27",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 Metro Ground taxi to holding point A 1 runway 18",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 cross runway 18 at A 1 taxi to holding point C runway 27",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 Metro ground Cleared to Smallville T 1 A departure Squawk 3456 slot time 1905",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 start up approved contact Metro Ground 118.750 for taxi instructions",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 Metro Ground after the red and white Antonov with the purple fin taxi to holding point runway 08",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 after landing Airbus 321 cross Runway 9 at C 2 after",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 Metro ground line up runway 27",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 runway 27 cleared for take off",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 Metro Tower hold at C 1",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 hold position amendment to clearance T 3 F departure climb to 6000 feet",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 hold position after departure climb to altitude 6000 feet",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 behind landing Boeing 757 line up runway 27 behind",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 hold position Cancel take off I say again cancel take off due to vehicle on the runway",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 stop immediately Big Jet 345 stop immediately",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 Metro Radar radar contact",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 fly heading 260 degrees climb to FL 100 no speed restrictions",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 fly direct BONNY climb to FL 360",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 Northern Control fly direct CLYDE",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 turn left immediately heading 270 to avoid traffic at 2 o'clock 5 miles crossing right to left 500 feet below",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 turn right immediately heading 30 degrees to avoid traffic at 2 o'clock 5 miles crossing right to left 500 feet below",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 climb immediately to FL 160 traffic at 12 o'clock 3 miles opposite direction same level",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 descend immediately to FL 160 traffic at 12 o'clock 3 miles opposite direction same level",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 Metro Approach now information Q new QNH 998",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 leave MAYFIELD heading 120 descend to 6000 feet QNH 998 speed 210 knots",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 turn right heading 180 speed 180 knots vectoring ILS runway 27 Right",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 turn right heading 240 descend to 3000 feet report established localiser runway 27 Right",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 cleared ILS approach runway 27 Right",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 turn right heading 240 degrees cleared ILS approach runway 27 Right maintain 3000 feet until glide path interception",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 continue approach",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 cleared to land runway 27 Right wind 270 degrees 10 knots",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 go around",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 Roger the MAYDAY turn left heading 090 radar vectors ILS runway 27",
		},
		{
			name:       "test should work",
			scanString: "Big Jet 345 roger turn right heading 140 for radar vectoring runway 09 descend to 3000 feet QNH 995 report established",
		},
	}
	lexer, err := InitLexer()
	require.NoError(t, err)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s, err := lexer.Scanner([]byte(strings.ToLower(test.scanString)))
			require.NoError(t, err)
			fmt.Printf("INPUT: %s\n", strings.ToLower(test.scanString))
			fmt.Println("Type            | Lexeme               | Position")
			fmt.Println("----------------+----------------------+------------")
			for tok, err, eof := s.Next(); !eof; tok, err, eof = s.Next() {
				if err != nil {
					log.Fatal(err)
				}
				token := tok.(*lex.Token)
				fmt.Printf("%-15v | %-20v | %v:%v-%v:%v\n", Tokens[token.Type], string(token.Lexeme), token.StartLine, token.StartColumn, token.EndLine, token.EndColumn)
			}
			fmt.Println("")
		})
	}
}
