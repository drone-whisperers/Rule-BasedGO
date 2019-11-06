package structures

import "strconv"

//Taxi is a representation of a taxi action
type Taxi struct {
	HoldingPoint *HoldingPoint
	RunWay       *RunWay
	HoldPoint    string
}

//NewTaxiAction returns a taxi struct
func NewTaxiAction() *Taxi {
	return &Taxi{}
}

//AddHoldingPoint adds holdingpoint to taxi
func (t *Taxi) AddHoldingPoint(h *HoldingPoint) {
	t.HoldingPoint = h
}

//AddRunWay adds runway to taxi
func (t *Taxi) AddRunWay(r *RunWay) {
	t.RunWay = r
}

//AddHoldPoint adds holdpoint to taxi
func (t *Taxi) AddHoldPoint(h string) {
	t.HoldPoint = h
}

//HoldingPoint is a representation of a holding point
type HoldingPoint struct {
	Location string
}

//NewHoldingPoint returns a HoldingPoint struct
func NewHoldingPoint() *HoldingPoint {
	return &HoldingPoint{}
}

//AddHoldingPointLocation adds a location to a holdingpoint
func (h *HoldingPoint) AddHoldingPointLocation(s string) {
	h.Location = s
}

//RunWay is a representation of a runway
type RunWay struct {
	Number int
}

//NewRunWay is a RunwayStruct
func NewRunWay() *RunWay {
	return &RunWay{}
}

//AddRunWayNumber adds number to runway
func (r *RunWay) AddRunWayNumber(i string) {
	n, err := strconv.Atoi(i)
	if err != nil {
		n = 0
	}
	r.Number = n
}
