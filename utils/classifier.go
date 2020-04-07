package utils

import (
	// "errors"
	"reflect"

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
func (c *Classifier) restClassifier() {
	c.actionSlice = []interface{}{}
	c.scannedTokens = map[int]*lex.Token{}
	c.scanIndex = -1
}

//Classify classifies a sentence into objects
func (c *Classifier) Classify(s string) ([]interface{}, error) {
	c.restClassifier()
	err := c.NewScanner(s)
	if err != nil {
		return nil, err
	}
	c.nextToken()
	if lexStruct.Tokens[c.currentToken.Type] != "DRONE" {
		// return nil, errors.New("Not current Drone")
	}
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil, err
		}
		var obj interface{}
		obj = c.classifyTopLevel()
		if obj != nil && obj != reflect.Zero(reflect.TypeOf(obj)).Interface() {
			c.actionSlice = append(c.actionSlice, obj)
		}
	}
	actions := c.actionSlice
	c.actionSlice = []interface{}{}
	return actions, nil
}

//ClassifyTopLevel classifies a at the top into objects
func (c *Classifier) classifyTopLevel() interface{} {
	c.backToken()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		var obj interface{}
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "ACTION":
			obj = c.classifyAction()
			return obj
		case "TAXI":
			obj = c.classifyTaxi()
			return obj
		case "CONTACT":
			obj = c.classifyContact()
			return obj
		case "CROSS":
			obj = c.classifyCross()
			return obj
		case "CONDITION":
			obj = c.classifyCondition()
			return obj
		case "RUNWAY":
			obj = c.classifyRunWay()
			return obj
		case "HOLDPOINT":
			obj = c.classifyHoldingPoint()
			return obj
		case "TRAFFIC":
			obj = c.classifyTraffic()
			return obj
		case "QNH":
			obj = c.classifyQNH()
			return obj
		case "ILS", "LOCALISER":
			obj = c.classifyILS()
			return obj
		case "APPROACH":
			obj = c.classifyApproach()
			return obj
		case "RADAR VECTORS", "RADAR VECTORING":
			obj = c.classifyRadarVectors()
			return obj
		default:
			continue
		}
	}
	return nil
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

func (c *Classifier) classifyAction() interface{} {
	var obj interface{}
	c.backToken()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch string(c.currentToken.Lexeme) {
		case "start up":
			obj = c.classifyStartUp()
			return obj
		case "landing":
			obj = c.classifyLanding()
			return obj
		case "line up":
			obj = c.classifyLineUp()
			return obj
		case "departure", "leave":
			obj = c.classifyDeparture()
			return obj
		case "climb":
			obj = c.classifyClimb()
			return obj
		case "take off":
			obj = c.classifyTakeOff()
			return obj
		case "stop":
			obj = c.classifyStop()
			return obj
		case "fly":
			obj = c.classifyFlyAction()
			return obj
		case "turn":
			obj = c.classifyTurnAction()
			return obj
		case "avoid":
			obj = c.classifyAvoidAction()
			return obj
		case "crossing":
			obj = c.classifyCrossingAction()
			return obj
		case "descend":
			obj = c.classifyDescendAction()
			return obj
		case "maintain":
			obj = c.classifyMaintain()
			return obj
		case "glide path interception":
			obj = c.classifyGlidePathInterception()
			return obj
		case "land":
			obj = c.classifyLandAction()
			return obj
		case "go around":
			obj = c.classifyGoAround()
			return obj
		default:
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyTaxi() *structures.Taxi {
	obj := structures.NewTaxiAction()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "HOLDING POINT":
			obj.AddHoldingPoint(c.classifyHoldingPoint())
		case "HOLDPOINT":
			obj.AddHoldPoint(string(c.currentToken.Lexeme))
		case "RUNWAY":
			obj.AddRunWay(c.classifyRunWay())
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}

	}
	return obj
}

func (c *Classifier) classifyContact() *structures.Contact {
	obj := structures.NewContactAction()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "TOWER":
			obj.AddTarget(string(c.currentToken.Lexeme))
		case "NUMBER":
			obj.AddFrequency(string(c.currentToken.Lexeme))
		case "INFO":
			obj.AddRequest(string(c.currentToken.Lexeme))
		case "ACTION", "TAXI":
			obj.AddRequest(c.classifyTopLevel())
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}

	}
	return obj
}

func (c *Classifier) classifyCross() *structures.Cross {
	obj := structures.NewCrossAction()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "RUNWAY":
			obj.AddRunWay(c.classifyRunWay())
		case "LOCATION":
			obj.AddLocation(string(c.currentToken.Lexeme))
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyHoldingPoint() *structures.HoldingPoint {
	obj := structures.NewHoldingPoint()
	if string(c.currentToken.Lexeme) == "hold position" {
		obj.HoldAtCurrentPosition()
	}
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "LOCATION":
			obj.AddHoldingPointLocation(string(c.currentToken.Lexeme))
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			{
				c.backToken()
				return obj
			}
		}
	}
	return obj
}

func (c *Classifier) classifyStartUp() *structures.StartUp {
	obj := structures.NewStartUpAction()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "INFO":
			continue
		case "CONDITION":
			obj.AddCondition(c.classifyCondition())
		default:
			{
				c.backToken()
				return obj
			}
		}
	}
	return obj
}

func (c *Classifier) classifyCondition() *structures.Condition {
	obj := structures.NewCondition()
	c.backToken()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		if lexStruct.Tokens[c.currentToken.Type] == "CONNECTOR" {
			continue
		}
		switch string(c.currentToken.Lexeme) {
		case "immediately":
			obj.IsImmediateAction()
		case "amendment":
			obj.IsAmendment()
		case "cancel":
			obj.IsCanceled()
			err := c.nextToken()
			if err != nil {
				return nil
			}
			obj.AddPreCondition(c.classifyTopLevel())
		case "continue":
			obj.ToggleRequirementsMet()
			err := c.nextToken()
			if err != nil {
				return nil
			}
			obj.AddPreCondition(c.classifyTopLevel())
		case "cleared", "clearance":
			obj.ToggleRequirementsMet()
			c.nextToken()
			if lexStruct.Tokens[c.currentToken.Type] == "CONNECTOR" {
				c.nextToken()
			}
			if lexStruct.Tokens[c.currentToken.Type] == "LOCATION" {
				c.backToken()
				obj.AddPreCondition(c.classifyDeparture())
			} else {
				c.backToken()
			}
		case "approved":
			obj.ToggleRequirementsMet()

		case "after", "behind", "until":
			if _ = c.nextToken(); c.eof || err != nil {
				return nil
			}
			obj.AddPreCondition(c.classifyAction())
			err = c.nextToken()
			if err == nil && !c.eof {
				if lexStruct.IsActionOrKeyword(c.currentToken) {
					obj.AddPostCondition(c.classifyTopLevel())
				}
			}
			return obj
		case "INFO":
			continue
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
	notEmpty := false
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "NUMBER":
			obj.AddRunWayNumber(string(c.currentToken.Lexeme))
			notEmpty = true
		case "CONDITION":
			obj.AddCondition(c.classifyCondition())
			notEmpty = true
		case "ACTION":
			if string(c.currentToken.Lexeme) == "take off" {
				notEmpty = true
				obj.IsTakeOff()
			} else {
				c.backToken()
				return obj
			}
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	if !notEmpty {
		return nil
	}
	return obj
}

func (c *Classifier) classifyLanding() *structures.Landing {
	obj := structures.NewLandingAction()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "PLANE":
			obj.AddPlane(string(c.currentToken.Lexeme))
		case "INFO":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyLineUp() *structures.LineUp {
	obj := structures.NewLineUpAction()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "RUNWAY":
			obj.AddRunWay(c.classifyRunWay())
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyDeparture() *structures.Departure {
	obj := structures.NewDepartureAction()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		if string(c.currentToken.Lexeme) == "departure" {
			continue
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "LOCATION":
			obj.AddLocation(string(c.currentToken.Lexeme))
		case "HEADING":
			obj.AddHeading(c.classifyHeading())
		case "SQUAWK":
			obj.AddSquawk(c.classifySquawk())
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyClimb() *structures.Climb {
	obj := structures.NewClimbAction()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "ALTITUDE":
			obj.AddAltitude(c.classifyAltitude())
		case "FL", "FLIGHT LEVEL":
			c.backToken()
			obj.AddAltitude(c.classifyAltitude())
		case "NUMBER":
			c.backToken()
			obj.AddAltitude(c.classifyAltitude())
		case "CONDITION":
			obj.AddCondition(c.classifyCondition())
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyDescendAction() *structures.Descend {
	obj := structures.NewDescendAction()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "ALTITUDE":
			obj.AddAltitude(c.classifyAltitude())
		case "FL", "FLIGHT LEVEL":
			c.backToken()
			obj.AddAltitude(c.classifyAltitude())
		case "NUMBER":
			c.backToken()
			obj.AddAltitude(c.classifyAltitude())
		case "CONDITION":
			obj.AddCondition(c.classifyCondition())
		case "QNH":
			obj.AddQNH(c.classifyQNH())
		case "SPEED":
			obj.AddSpeed(c.classifySpeed())
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyAltitude() *structures.Altitude {
	obj := structures.NewAltitude()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "UNIT":
			obj.AddUnit(string(c.currentToken.Lexeme))
		case "NUMBER":
			obj.AddNumber(string(c.currentToken.Lexeme))
		case "FL", "FLIGHT LEVEL":
			obj.AddFlightLevel(c.classifyFlightLevel())
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyTakeOff() *structures.TakeOff {
	obj := structures.NewTakeOff()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyStop() *structures.Stop {
	obj := structures.NewStopAction()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "INFO":
			continue
		case "CONDITION":
			obj.AddCondition(c.classifyCondition())
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyFlyAction() *structures.Fly {
	obj := structures.NewFlyAction()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "HEADING":
			obj.AddHeading(c.classifyHeading())
		case "LOCATION":
			obj.AddLocation(string(c.currentToken.Lexeme))
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyHeading() *structures.Heading {
	obj := structures.NewHeading()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "NUMBER":
			obj.AddNumber(string(c.currentToken.Lexeme))
		case "UNIT":
			obj.AddUnit(string(c.currentToken.Lexeme))
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyFlightLevel() *structures.FlightLevel {
	obj := structures.NewFlightLevel()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "NUMBER":
			obj.AddLevel(string(c.currentToken.Lexeme))
		case "NO SPEED RESTRICTIONS":
			obj.AddAdditionalInfo(string(c.currentToken.Lexeme))
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyTurnAction() *structures.Turn {
	obj := structures.NewTurnAction()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "DIRECTION":
			obj.AddDirection(string(c.currentToken.Lexeme))
		case "HEADING":
			obj.AddHeading(c.classifyHeading())
		case "CONDITION":
			if string(c.currentToken.Lexeme) == "cleared" {
				c.backToken()
				return obj
			}
			obj.AddCondition(c.classifyCondition())
		case "SPEED":
			obj.AddSpeed(c.classifySpeed())
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyAvoidAction() *structures.Avoid {
	obj := structures.NewAvoidAction()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "TRAFFIC":
			obj.AddObject(c.classifyTraffic())
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyCrossingAction() *structures.Crossing {
	obj := structures.NewCrossing()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "DIRECTION":
			if obj.From == "" {
				obj.AddFrom(string(c.currentToken.Lexeme))
			} else {
				obj.AddTo(string(c.currentToken.Lexeme))
			}
		case "UNIT":
			obj.AddUnit(string(c.currentToken.Lexeme))
		case "NUMBER":
			obj.AddNumber(string(c.currentToken.Lexeme))
		case "INFO":
			obj.AddInfo(string(c.currentToken.Lexeme))
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyDistance() *structures.Distance {
	obj := structures.NewDistance()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "UNIT":
			obj.AddUnit(string(c.currentToken.Lexeme))
		case "NUMBER":
			obj.AddNumber(string(c.currentToken.Lexeme))
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyTraffic() *structures.Traffic {
	obj := structures.NewTraffic()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "ACTION":
			obj.AddAction(c.classifyAction())
		case "NUMBER":
			c.nextToken()
			if string(c.currentToken.Lexeme) == "o'clock" {
				c.backToken()
				obj.AddHeading(structures.NewHeading())
				obj.Direction.AddNumber(string(c.currentToken.Lexeme))
				c.nextToken()
				obj.Direction.AddUnit(string(c.currentToken.Lexeme))
			} else {
				c.backToken()
				obj.AddDistance(structures.NewDistance())
				obj.Distance.AddNumber(string(c.currentToken.Lexeme))
				c.nextToken()
				obj.Distance.AddUnit(string(c.currentToken.Lexeme))
			}
		case "INFO":
			obj.AddInfo(string(c.currentToken.Lexeme))
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyQNH() *structures.QNH {
	obj := structures.NewQNH()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "NUMBER":
			obj.AddPressure(string(c.currentToken.Lexeme))
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifySpeed() *structures.Speed {
	obj := structures.NewSpeed()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "NUMBER":
			obj.AddNumber(string(c.currentToken.Lexeme))
		case "UNIT":
			obj.AddUnit(string(c.currentToken.Lexeme))
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyILS() *structures.ILS {
	obj := structures.NewILS()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "RUNWAY":
			obj.AddRunWay(c.classifyRunWay())
		case "DIRECTION":
			obj.AddDirection(string(c.currentToken.Lexeme))
		case "INFO":
			continue
		case "APPROACH":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyGlidePathInterception() *structures.GlidePathInterception {
	obj := structures.NewGlidePathInterception()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyMaintain() *structures.Maintain {
	obj := structures.NewMaintain()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "Altitude":
			obj.AddAltitude(c.classifyAltitude())
		case "NUMBER":
			c.backToken()
			obj.AddAltitude(c.classifyAltitude())
		case "CONDITION":
			obj.AddCondition(c.classifyCondition())
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyApproach() *structures.Approach {
	obj := structures.NewApproach()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyLandAction() *structures.Land {
	obj := structures.NewLand()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "RUNWAY":
			obj.AddRunWay(c.classifyRunWay())
		case "DIRECTION":
			obj.AddDirection(string(c.currentToken.Lexeme))
		case "WIND":
			obj.AddWind(c.classifyWind())
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyWind() *structures.Wind {
	obj := structures.NewWind()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "HEADING":
			obj.AddHeading(c.classifyHeading())
		case "NUMBER":
			c.nextToken()
			if string(c.currentToken.Lexeme) == "degrees" {
				c.backToken()
				obj.AddHeading(structures.NewHeading())
				obj.Heading.AddNumber(string(c.currentToken.Lexeme))
				c.nextToken()
				obj.Heading.AddUnit(string(c.currentToken.Lexeme))
			} else if string(c.currentToken.Lexeme) == "knots" {
				c.backToken()
				obj.AddSpeed(structures.NewSpeed())
				obj.Speed.AddNumber(string(c.currentToken.Lexeme))
				c.nextToken()
				obj.Speed.AddUnit(string(c.currentToken.Lexeme))
			}
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyGoAround() *structures.GoAround {
	obj := structures.NewGoAround()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifyRadarVectors() *structures.RadarVectors {
	obj := structures.NewRadarVectors()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "ILS":
			obj.AddILS(c.classifyILS())
		case "RUNWAY":
			obj.AddRunWay(c.classifyRunWay())
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifySquawk() *structures.Squawk {
	obj := structures.NewSquawk()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "NUMBER":
			obj.AddNumber(string(c.currentToken.Lexeme))
		case "SLOT TIME":
			obj.AddSlotTime(c.classifySlotTime())
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}

func (c *Classifier) classifySlotTime() *structures.SlotTime {
	obj := structures.NewSlotTime()
	for err := c.nextToken(); !c.eof; err = c.nextToken() {
		if err != nil {
			return nil
		}
		switch lexStruct.Tokens[c.currentToken.Type] {
		case "NUMBER":
			obj.AddNumber(string(c.currentToken.Lexeme))
		case "INFO":
			continue
		case "CONNECTOR":
			continue
		default:
			c.backToken()
			return obj
		}
	}
	return obj
}
