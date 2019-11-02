package lexer

import (
	"strings"

	lex "github.com/timtadh/lexmachine"
	"github.com/timtadh/lexmachine/machines"
)

// Keywords represents keyword tokens
var Keywords []string

// Connectors represents keyword tokens
var Connectors []string

// HoldPoints represents keyword tokens
var HoldPoints []string

// Condition represents keyword tokens
var Condition []string

// Tokens are all tokens
var Tokens []string

// TokenIds is a map from the token names to their int ids
var TokenIds map[string]int

func initTokens() {
	Keywords = []string{
		"RUNWAY",
		"TAXI",
		"HOLDING POINT",
		"CONTACT",
		"CROSS",
		"SQUAWK",
	}
	Tokens = []string{
		"COMMENT",
		"ID",
		"NUMBER",
		"DRONE",
		"TOWER",
		"LOCATION",
	}

	Connectors = []string{
		"to",
		"and",
		"from",
		"of",
		"at",
	}

	HoldPoints = []string{
		"hold",
		"hold position",
		"hold short of",
	}

	Condition = []string{
		"before",
		"after",
	}

	Tokens = append(Tokens, Keywords...)
	Tokens = append(Tokens, "CONNECTOR")
	Tokens = append(Tokens, "HOLDPOINT")
	Tokens = append(Tokens, "CONDITION")
	Tokens = append(Tokens, "DIRECTION")
	TokenIds = make(map[string]int)
	for i, tok := range Tokens {
		TokenIds[tok] = i
	}
}

// Creates the lexer object and compiles the NFA.
func initLexer(lexerType string) (*lex.Lexer, error) {
	lexer := lex.NewLexer()
	initTokens()

	for _, con := range Connectors {
		lexer.Add([]byte(con), token("CONNECTOR"))
	}
	for _, hold := range HoldPoints {
		lexer.Add([]byte(hold), token("HOLDPOINT"))
	}
	for _, cond := range Condition {
		lexer.Add([]byte(cond), token("CONDITION"))
	}

	for _, name := range Keywords {
		lexer.Add([]byte(strings.ToLower(name)), token(name))
	}

	lexer.Add([]byte("metro ground"), token("TOWER"))
	lexer.Add([]byte("big jet 345"), token("DRONE"))

	lexer.Add([]byte(`\w\s\d+`), token("LOCATION"))
	lexer.Add([]byte(`\w\s\d+?(\s\w\s)`), token("LOCATION"))

	lexer.Add([]byte(`\d?(\d)\s\w*o'clock\w*`), token("DIRECTION"))

	lexer.Add([]byte(`//[^\n]*\n?`), token("COMMENT"))
	lexer.Add([]byte(`/\*([^*]|\r|\n|(\*+([^*/]|\r|\n)))*\*+/`), token("COMMENT"))
	lexer.Add([]byte(`([a-z]|[A-Z])([a-z]|[A-Z]|[0-9]|_)*`), token("LOCATION"))
	lexer.Add([]byte(`"([^\\"]|(\\.))*"`), token("LOCATION"))
	lexer.Add([]byte(`"([^\\"]|(\\.))*"`), token("LOCATION"))
	lexer.Add([]byte("( |\t|\n|\r)+"), skip)

	lexer.Add([]byte(`-?\d+(,\d+)*(\.\d+(e\d+)?)?`), token("NUMBER"))
	err := lexer.Compile()
	if err != nil {
		return nil, err
	}
	return lexer, nil
}

// a lex.Action function which skips the match.
func skip(*lex.Scanner, *machines.Match) (interface{}, error) {
	return nil, nil
}

// a lex.Action function with constructs a Token of the given token type by
// the token type's name.
func token(name string) lex.Action {
	return func(s *lex.Scanner, m *machines.Match) (interface{}, error) {
		return s.Token(TokenIds[name], string(m.Bytes), m), nil
	}
}
