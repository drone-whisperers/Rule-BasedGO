package lexer

import (
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type testLexer struct {
	name       string
	scanString string
}

func BenchmarkLexer(b *testing.B) {
	tests := []testLexer{
		{
			scanString: "big jet 345 metro ground taxi to holding point a 1 hold short of runway 18",
		},
		{
			scanString: "big jet 345 contact metro ground 119.2",
		},
		{
			scanString: "Big Jet 345 Metro Ground taxi to holding point C runway 27",
		},
		{
			scanString: "Big Jet 345 Metro Ground taxi to holding point C runway 27",
		},
		{
			scanString: "Big Jet 345 Metro Ground taxi to holding point A 1 runway 18",
		},
		{
			scanString: "Big Jet 345 cross runway 18 at A 1 taxi to holding point C runway 27",
		},
		{
			scanString: "Big Jet 345 Metro ground Cleared to Smallville T 1 A departure Squawk 3456 slot time 1905",
		},
		{
			scanString: "Big Jet 345 start up approved contact Metro Ground 118.750 for taxi instructions",
		},
		{
			scanString: "Big Jet 345 Metro Ground after the red and white Antonov with the purple fin taxi to holding point runway 08",
		},
		{
			scanString: "Big Jet 345 after landing Airbus 321 cross Runway 9 at C 2 after",
		},
		{
			scanString: "Big Jet 345 Metro ground line up runway 27",
		},
		{
			scanString: "Big Jet 345 runway 27 cleared for take off",
		},
		{
			scanString: "Big Jet 345 Metro Tower hold at C 1",
		},
		{
			scanString: "Big Jet 345 hold position amendment to clearance T 3 F departure climb to 6000 feet",
		},
		{
			scanString: "Big Jet 345 hold position after departure climb to altitude 6000 feet",
		},
		{
			scanString: "Big Jet 345 behind landing Boeing 757 line up runway 27 behind",
		},
		{
			scanString: "Big Jet 345 hold position Cancel take off I say again cancel take off due to vehicle on the runway",
		},
		{
			scanString: "Big Jet 345 stop immediately Big Jet 345 stop immediately",
		},
		{
			scanString: "Big Jet 345 Metro Radar radar contact",
		},
		{
			scanString: "Big Jet 345 fly heading 260 degrees climb to FL 100 no speed restrictions",
		},
		{
			scanString: "Big Jet 345 fly direct BONNY climb to FL 360",
		},
		{
			scanString: "Big Jet 345 Northern Control fly direct CLYDE",
		},
		{
			scanString: "Big Jet 345 turn left immediately heading 270 to avoid traffic at 2 o'clock 5 miles crossing right to left 500 feet below",
		},
		{
			scanString: "Big Jet 345 turn right immediately heading 30 degrees to avoid traffic at 2 o'clock 5 miles crossing right to left 500 feet below",
		},
		{
			scanString: "Big Jet 345 climb immediately to FL 160 traffic at 12 o'clock 3 miles opposite direction same level",
		},
		{
			scanString: "Big Jet 345 descend immediately to FL 160 traffic at 12 o'clock 3 miles opposite direction same level",
		},
		{
			scanString: "Big Jet 345 Metro Approach now information Q new QNH 998",
		},
		{
			scanString: "Big Jet 345 leave MAYFIELD heading 120 descend to 6000 feet QNH 998 speed 210 knots",
		},
		{
			scanString: "Big Jet 345 turn right heading 180 speed 180 knots vectoring ILS runway 27 Right",
		},
		{
			scanString: "Big Jet 345 turn right heading 240 descend to 3000 feet report established localiser runway 27 Right",
		},
		{
			scanString: "Big Jet 345 cleared ILS approach runway 27 Right",
		},
		{
			scanString: "Big Jet 345 turn right heading 240 degrees cleared ILS approach runway 27 Right maintain 3000 feet until glide path interception",
		},
		{
			scanString: "Big Jet 345 continue approach",
		},
		{
			scanString: "Big Jet 345 cleared to land runway 27 Right wind 270 degrees 10 knots",
		},
		{
			scanString: "Big Jet 345 go around",
		},
		{
			scanString: "Big Jet 345 Roger the MAYDAY turn left heading 090 radar vectors ILS runway 27",
		},
		{
			scanString: "Big Jet 345 roger turn right heading 140 for radar vectoring runway 09 descend to 3000 feet QNH 995 report established",
		},
	}
	lexer, err := InitLexer("big jet 345")
	require.NoError(b, err)
	for _, test := range tests {
		b.Run(test.scanString, func(b *testing.B) {
			s, err := lexer.Scanner([]byte(strings.ToLower(test.scanString)))
			require.NoError(b, err)
			// fmt.Printf("INPUT: %s\n", strings.ToLower(test.scanString))
			// fmt.Println("Type            | Lexeme               | Position")
			// fmt.Println("----------------+----------------------+------------")
			for _, err, eof := s.Next(); !eof; _, err, eof = s.Next() {
				if err != nil {
					log.Fatal(err)
				}
				// token := tok.(*lex.Token)
				// fmt.Printf("%-15v | %-20v | %v:%v-%v:%v\n", Tokens[token.Type], string(token.Lexeme), token.StartLine, token.StartColumn, token.EndLine, token.EndColumn)
			}
			// fmt.Println("")
		})
	}
}
