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

// Actions represents keyword tokens
var Actions []string

// Units represents keyword tokens
var Units []string

// Confirmation represents keyword tokens
var Confirmation []string

// Requests represents keyword tokens
var Requests []string

// Locations represents keyword tokens
var Locations []string

// Directions represents keyword tokens
var Directions []string

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
		"SLOT TIME",
		"FL",
		"FLIGHT LEVEL",
		"RADAR VECTORS",
		"RADAR VECTORING",
		"WIND",
		"GLIDE PATH INTERCEPTION",
		"NO SPEED RESTRICTIONS",
		"RADAR CONTACT",
		"TRAFFIC",
		"REPORT ESTABLISHED",
		"LOCALISER",
		"ALTITUDE",
		"HEADING",
		"QNH",
		"SPEED",
		"ILS",
		"APPROACH",
	}

	Actions = []string{
		"departure",
		"landing",
		"line up",
		"start up",
		"climb",
		"stop",
		"turn",
		"descend",
		"ascend",
		"go around",
		"land",
		"maintain",
		"fly",
		"leave",
		"take off",
		"avoid",
		"crossing",
		"glide path interception",
	}

	Tokens = []string{
		"COMMENT",
		"ID",
		"NUMBER",
		"LOCATION",
		"DIRECTION",
		"CONDITION",
		"DRONE",
		"TOWER",
		"CONNECTOR",
		"HOLDPOINT",
		"ACTION",
		"PLANE",
		"UNIT",
		"CONFIRMATION",
		"INFO",
		"REQUEST",
	}

	Connectors = []string{
		"to",
		"and",
		"from",
		"of",
		"at",
		"for",
		"due",
		"the",
		"direct",
		"with",
	}

	HoldPoints = []string{
		"hold",
		"hold position",
		"hold short of",
	}

	Condition = []string{
		"before",
		"after",
		"cleared",
		"approved",
		"behind",
		"cancel",
		"immediately",
		"continue",
		"until",
		"clearance",
		"now",
		"amendment",
	}

	Units = []string{
		"feet",
		"degrees",
		"o'clock",
		"knots",
		"miles",
		"o'clock",
	}

	Confirmation = []string{
		"roger",
		"roger that",
		"roger the mayday",
		"i say again",
	}

	Locations = []string{
		"bonny",
		"clyde",
		"c 2",
		"a 1",
		"c",
		"c 1",
		"mayfield",
		"smallville",
		"t 1 a",
		"t 3 f",
	}

	Directions = []string{
		"left",
		"right",
	}

	Tokens = append(Tokens, Keywords...)

	TokenIds = make(map[string]int)
	for i, tok := range Tokens {
		TokenIds[tok] = i
	}
}

// InitLexer Creates the lexer object and compiles the NFA.
func InitLexer(name string) (*lex.Lexer, error) {
	lexer := lex.NewLexer()
	initTokens()
	//This assumes we know our name
	lexer.Add([]byte(name), token("DRONE"))

	//Connectors are a list of keyword connectors
	for _, con := range Connectors {
		lexer.Add([]byte(con), token("CONNECTOR"))
	}
	//holdPoints are a list of key Phrases
	for _, hold := range HoldPoints {
		lexer.Add([]byte(hold), token("HOLDPOINT"))
	}
	//Conditions Represent a conditional Statement
	for _, cond := range Condition {
		lexer.Add([]byte(cond), token("CONDITION"))
	}
	//KeyWords are a list of keyword to search for
	for _, name := range Keywords {
		lexer.Add([]byte(strings.ToLower(name)), token(name))
	}

	//Actions are a list of actions to search for
	for _, act := range Actions {
		lexer.Add([]byte(act), token("ACTION"))
	}

	//Actions are a list of actions to search for
	for _, unit := range Units {
		lexer.Add([]byte(unit), token("UNIT"))
	}

	//Actions are a list of actions to search for
	for _, conf := range Confirmation {
		lexer.Add([]byte(conf), token("CONFIRMATION"))
	}

	//Actions are a list of actions to search for
	for _, request := range Requests {
		lexer.Add([]byte(request), token("REQUEST"))
	}

	//Actions are a list of actions to search for
	for _, location := range Locations {
		lexer.Add([]byte(location), token("LOCATION"))
	}

	//Directions are a list of actions to search for
	for _, direction := range Directions {
		lexer.Add([]byte(direction), token("DIRECTION"))
	}

	//This assumes we know the tower name
	lexer.Add([]byte("metro ground"), token("TOWER"))
	lexer.Add([]byte("metro approach"), token("TOWER"))
	lexer.Add([]byte("metro tower"), token("TOWER"))
	lexer.Add([]byte("metro radar"), token("TOWER"))
	lexer.Add([]byte("northern control"), token("TOWER"))

	//Matches locations of format Letter NUMBERS letter
	// lexer.Add([]byte(`\w`), token("LOCATION"))
	// lexer.Add([]byte(`\w\s\d+`), token("LOCATION"))
	// lexer.Add([]byte(`\w\s\d+?(\s\w\s)`), token("LOCATION"))

	//matches numbers (optional decimals)
	lexer.Add([]byte(`-?\d+(,\d+)*(\.\d+(e\d+)?)?`), token("NUMBER"))
	lexer.Add([]byte(`[0-9]+:[0-9]+`), token("NUMBER"))

	//Matches planes
	lexer.Add([]byte(`\w*airbus\w*\s\d+`), token("PLANE"))
	lexer.Add([]byte(`\w*boeing\w*\s\d+`), token("PLANE"))
	lexer.Add([]byte(`antonov`), token("PLANE"))

	lexer.Add([]byte(`([a-z]|[A-Z])([a-z]|[A-Z]|[0-9]|_)*`), token("INFO"))

	//skip useless characters
	lexer.Add([]byte("( |\t|\n|\r)+"), skip)

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

// IsActionOrKeyword is helper function for pre and post conditions
func IsActionOrKeyword(token *lex.Token) bool {
	if Tokens[token.Type] == "ACTION" {
		return true
	}
	for _, k := range Keywords {
		if k == strings.ToUpper(string(token.Lexeme)) {
			return true
		}
	}
	return false
}
