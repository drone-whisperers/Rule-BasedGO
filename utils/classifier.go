package utils

import (
	lexStruct "github.com/Rule-BasedGO/lexer"
	"github.com/Rule-BasedGO/structures"
	lex "github.com/timtadh/lexmachine"
)

//Classifier used to classify a sentence into objects
type Classifier struct {
	lexer         *lex.Lexer
	currentToken  *lex.Token
	previousToken *lex.Token
	head          *lex.Token
	scannedTokens map[int]*lex.Token
	scanIndex     int
	scanner       *lex.Scanner
	eof           bool
	actionSlice   []interface{}
}

//NewClassifier returns a classifier interface
func NewClassifier(lexer *lex.Lexer) *Classifier {
	return &Classifier{
		lexer:         lexer,
		scannedTokens: map[int]*lex.Token{},
		scanIndex:     -1,
	}
}

//Classify classifies a sentence into objects
func (c *Classifier) Classify(s string) ([]interface{}, error) {
	err := c.NewScanner(s)
	if err != nil {
		return nil, err
	}
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil, err
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "ACTION":

		case "TAXI":
			c.classifyTaxi()
		case "CONTACT":
			c.classifyContact()
		default:
			continue

		}

	}
	return c.actionSlice, nil
}

//NewScanner starts a new scan of a sentence
func (c *Classifier) NewScanner(s string) error {
	scanner, err := c.lexer.Scanner([]byte(s))
	if err != nil {
		return err
	}
	c.scanner = scanner
	return nil
}

func (c *Classifier) nextToken() error {
	if c.head == c.currentToken || len(c.scannedTokens) == 0 {
		tok, err, eof := c.scanner.Next()
		if err != nil {
			return err
		}
		if !eof {
			token := tok.(*lex.Token)
			c.currentToken = token
			c.head = token
			c.scanIndex++
			if c.scanIndex > 0 {
				c.previousToken = c.scannedTokens[c.scanIndex-1]
			}
			c.scannedTokens[c.scanIndex] = token
		}
		c.eof = eof
		return nil
	}
	c.scanIndex++
	c.currentToken = c.scannedTokens[c.scanIndex]
	if c.scanIndex > 0 {
		c.previousToken = c.scannedTokens[c.scanIndex-1]
	}
	return nil
}

func (c *Classifier) backToken() error {
	token := c.scannedTokens[c.scanIndex-1]
	c.currentToken = token
	if c.scanIndex > 0 {
		c.previousToken = c.scannedTokens[c.scanIndex-2]
	}
	c.scanIndex--
	return nil
}

func (c *Classifier) classifyTaxi() {
	obj := structures.NewTaxiAction()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "HOLDING POINT":
			obj.AddHoldingPoint(c.classifyHoldingPoint())
		case "HOLDPOINT":
			obj.AddHoldPoint(string(c.currentToken.Lexeme))
		case "RUNWAY":
			obj.AddRunWay(c.classifyRunWay())
		case "CONNECTOR":
			continue
		default:
			c.actionSlice = append(c.actionSlice, obj)
			return
		}

	}
	c.actionSlice = append(c.actionSlice, obj)
}

func (c *Classifier) classifyContact() {
	obj := structures.NewContactAction()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "TOWER":
			obj.AddTarget(string(c.currentToken.Lexeme))
		case "NUMBER":
			obj.AddFrequency(string(c.currentToken.Lexeme))
		case "CONNECTOR":
			continue
		default:
			c.actionSlice = append(c.actionSlice, obj)
			return
		}

	}
	c.actionSlice = append(c.actionSlice, obj)
}

func (c *Classifier) classifyHoldingPoint() *structures.HoldingPoint {
	obj := structures.NewHoldingPoint()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "LOCATION":
			obj.AddHoldingPointLocation(string(c.currentToken.Lexeme))
		default:
			{
				c.backToken()
				return obj
			}
		}
	}
	return obj
}

func (c *Classifier) classifyRunWay() *structures.RunWay {
	obj := structures.NewRunWay()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "NUMBER":
			obj.AddRunWayNumber(string(c.currentToken.Lexeme))
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}
