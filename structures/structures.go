package structures

import (
	"strconv"
)

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

//Contact is a representation of a taxi action
type Contact struct {
	Target    string
	Frequency float64
	Request   interface{}
}

//NewContactAction returns a taxi struct
func NewContactAction() *Contact {
	return &Contact{}
}

//AddFrequency adds frequency to contact
func (c *Contact) AddFrequency(f string) {
	n, _ := strconv.ParseFloat(f, 64)
	c.Frequency = n
}

//AddTarget adds target to contact
func (c *Contact) AddTarget(t string) {
	c.Target = t
}

//AddRequest adds target to contact
func (c *Contact) AddRequest(i interface{}) {
	c.Request = i
}

//HoldingPoint is a representation of a holding point
type HoldingPoint struct {
	Location        string
	CurrentPosition bool
}

//NewHoldingPoint returns a HoldingPoint struct
func NewHoldingPoint() *HoldingPoint {
	return &HoldingPoint{}
}

//AddHoldingPointLocation adds a location to a holdingpoint
func (h *HoldingPoint) AddHoldingPointLocation(s string) {
	h.Location = s
}

//HoldAtCurrentPosition adds a location to a holdingpoint
func (h *HoldingPoint) HoldAtCurrentPosition() {
	h.CurrentPosition = true
}

//RunWay is a representation of a runway
type RunWay struct {
	Number    float64
	Condition *Condition
	TakeOff   bool
}

//NewRunWay is a RunwayStruct
func NewRunWay() *RunWay {
	return &RunWay{}
}

//AddRunWayNumber adds number to runway
func (r *RunWay) AddRunWayNumber(i string) {
	n, _ := strconv.ParseFloat(i, 64)
	r.Number = n
}

//IsTakeOff indcates runway is a takeoff
func (r *RunWay) IsTakeOff() {
	r.TakeOff = true
}

//AddCondition adds a condition to A Runway
func (r *RunWay) AddCondition(c *Condition) {
	r.Condition = c
}

//Cross is a representation of a Cross actions
type Cross struct {
	RunWay   *RunWay
	Location string
}

//NewCrossAction is a Cross struct
func NewCrossAction() *Cross {
	return &Cross{}
}

//AddRunWay adds number to runway
func (c *Cross) AddRunWay(r *RunWay) {
	c.RunWay = r
}

//AddLocation adds number to runway
func (c *Cross) AddLocation(s string) {
	c.Location = s
}

//StartUp is a representation of a StartUp actions
type StartUp struct {
	Condition *Condition
}

//NewStartUpAction is a StartUp struct
func NewStartUpAction() *StartUp {
	return &StartUp{}
}

//AddCondition Adds a condition to a start up
func (s *StartUp) AddCondition(c *Condition) {
	s.Condition = c
}

//Condition is Condition struct
type Condition struct {
	PreCondition    interface{}
	PostCondition   interface{}
	RequirementsMet bool
	Canceled        bool
	ImmediateAction bool
	Amendment       bool
}

//NewCondition is a StartUp struct
func NewCondition() *Condition {
	return &Condition{}
}

//ToggleRequirementsMet Adds a condition to a start up
func (c *Condition) ToggleRequirementsMet() {
	c.RequirementsMet = !c.RequirementsMet
}

//IsCanceled Cancels requirements
func (c *Condition) IsCanceled() {
	c.Canceled = true
}

//IsAmendment Amends an action
func (c *Condition) IsAmendment() {
	c.Amendment = true
}

//IsImmediateAction Cancels requirements
func (c *Condition) IsImmediateAction() {
	c.ImmediateAction = true
}

//AddPreCondition Adds a condition to a start up
func (c *Condition) AddPreCondition(a interface{}) {
	c.PreCondition = a
}

//AddPostCondition Adds a condition to a start up
func (c *Condition) AddPostCondition(a interface{}) {
	c.PostCondition = a
}

//Landing is a representation of a Landing actions
type Landing struct {
	Plane string
}

//NewLandingAction is a Landing struct
func NewLandingAction() *Landing {
	return &Landing{}
}

//AddPlane Adds a Plane to a start up
func (l *Landing) AddPlane(s string) {
	l.Plane = s
}

//LineUp is a representation of a Landing actions
type LineUp struct {
	RunWay *RunWay
}

//NewLineUpAction is a Landing struct
func NewLineUpAction() *LineUp {
	return &LineUp{}
}

//AddRunWay Adds a Runway to a Line up
func (l *LineUp) AddRunWay(r *RunWay) {
	l.RunWay = r
}

//Departure is a representation of a Departure actions
type Departure struct {
	Location []string
	Heading  *Heading
	Squawk   *Squawk
}

//NewDepartureAction is a Landing struct
func NewDepartureAction() *Departure {
	return &Departure{}
}

//AddLocation Adds a location to a departure
func (d *Departure) AddLocation(l string) {
	d.Location = append(d.Location, l)
}

//AddHeading Adds a location to a departure
func (d *Departure) AddHeading(h *Heading) {
	d.Heading = h
}

//AddSquawk Adds a Squawk to a departure
func (d *Departure) AddSquawk(s *Squawk) {
	d.Squawk = s
}

//Climb is a representation of a Departure actions
type Climb struct {
	Altitude  *Altitude
	Condition *Condition
}

//NewClimbAction is a Landing struct
func NewClimbAction() *Climb {
	return &Climb{}
}

//AddAltitude Adds a Runway to a Line up
func (c *Climb) AddAltitude(a *Altitude) {
	c.Altitude = a
}

//AddCondition Adds a Runway to a Line up
func (c *Climb) AddCondition(con *Condition) {
	c.Condition = con
}

//Descend is a representation of a Departure actions
type Descend struct {
	Altitude  *Altitude
	QNH       *QNH
	Speed     *Speed
	Condition *Condition
}

//NewDescendAction is a Landing struct
func NewDescendAction() *Descend {
	return &Descend{}
}

//AddAltitude Adds a Runway to a Line up
func (d *Descend) AddAltitude(a *Altitude) {
	d.Altitude = a
}

//AddCondition Adds a Runway to a Line up
func (d *Descend) AddCondition(con *Condition) {
	d.Condition = con
}

//AddQNH Adds a speed to a decend
func (d *Descend) AddQNH(q *QNH) {
	d.QNH = q
}

//AddSpeed Adds a Speed to a decend
func (d *Descend) AddSpeed(s *Speed) {
	d.Speed = s
}

//Altitude is a representation of a Altitude
type Altitude struct {
	Number      float64
	Unit        string
	FlightLevel *FlightLevel
}

//NewAltitude is a Landing struct
func NewAltitude() *Altitude {
	return &Altitude{}
}

//AddNumber Adds a Number to a Altitude
func (a *Altitude) AddNumber(i string) {
	n, _ := strconv.ParseFloat(i, 64)
	a.Number = n
}

//AddUnit Adds a Unit to a Altitude
func (a *Altitude) AddUnit(s string) {
	a.Unit = s
}

//AddFlightLevel Adds a FlightLevel to a Altitude
func (a *Altitude) AddFlightLevel(f *FlightLevel) {
	a.FlightLevel = f
}

//TakeOff is a representation of a TakeOff
type TakeOff struct {
}

//NewTakeOff is a TakeOff struct
func NewTakeOff() *TakeOff {
	return &TakeOff{}
}

//Stop is a representation of a Stop action
type Stop struct {
	Condition *Condition
}

//NewStopAction is a Stop struct
func NewStopAction() *Stop {
	return &Stop{}
}

//AddCondition Adds a Condition to a Stop action
func (s *Stop) AddCondition(c *Condition) {
	s.Condition = c
}

//Fly is a representation of a Stop action
type Fly struct {
	Heading   *Heading
	Locations []string
}

//NewFlyAction is a fly struct
func NewFlyAction() *Fly {
	return &Fly{}
}

//AddHeading Adds a Hading to a fly action
func (f *Fly) AddHeading(h *Heading) {
	f.Heading = h
}

//AddLocation adds a location to a fly
func (f *Fly) AddLocation(s string) {
	f.Locations = append(f.Locations, s)
}

//Heading is a representation of a Heading
type Heading struct {
	Number float64
	Unit   string
}

//NewHeading is a Heading struct
func NewHeading() *Heading {
	return &Heading{}
}

//AddNumber Adds a Number to a Heading
func (h *Heading) AddNumber(i string) {
	n, _ := strconv.ParseFloat(i, 64)
	h.Number = n
}

//AddUnit Adds a Number to a Heading
func (h *Heading) AddUnit(i string) {
	h.Unit = i
}

//IsEmpty Finds if heading is empty
func (h *Heading) IsEmpty() bool {
	if h.Unit == "" && h.Number == 0 {
		return true
	}
	return false
}

//FlightLevel is a representation of a FlightLevel
type FlightLevel struct {
	Level          float64
	AdditionalInfo []string
}

//NewFlightLevel is a FlightLevel struct
func NewFlightLevel() *FlightLevel {
	return &FlightLevel{}
}

//AddLevel Adds a Number to a flight Level
func (f *FlightLevel) AddLevel(i string) {
	n, _ := strconv.ParseFloat(i, 64)
	f.Level = n
}

//AddAdditionalInfo Adds additional info to a flight Level
func (f *FlightLevel) AddAdditionalInfo(i string) {
	f.AdditionalInfo = append(f.AdditionalInfo, i)
}

//Turn is a representation of a Turn
type Turn struct {
	Direction string
	Condition *Condition
	Heading   *Heading
	Speed     *Speed
}

//NewTurnAction is a Turn struct
func NewTurnAction() *Turn {
	return &Turn{}
}

//AddCondition Adds a Condition to a turn
func (t *Turn) AddCondition(c *Condition) {
	t.Condition = c
}

//AddHeading Adds a Heading to a turn
func (t *Turn) AddHeading(h *Heading) {
	t.Heading = h
}

//AddDirection Adds a direction to a turn
func (t *Turn) AddDirection(s string) {
	t.Direction = s
}

//AddSpeed Adds a Speed to a turn
func (t *Turn) AddSpeed(s *Speed) {
	t.Speed = s
}

//Avoid is a representation of a Turn
type Avoid struct {
	Object interface{}
}

//NewAvoidAction is a Avoid struct
func NewAvoidAction() *Avoid {
	return &Avoid{}
}

//AddObject Adds a Object to a avoid
func (a *Avoid) AddObject(i interface{}) {
	a.Object = i
}

//Traffic is a representation of traffic
type Traffic struct {
	Direction *Heading
	Distance  *Distance
	Action    interface{}
	Info      []string
}

//NewTraffic is a Traffic struct
func NewTraffic() *Traffic {
	return &Traffic{}
}

//AddHeading Adds a Heading to a Avoid
func (t *Traffic) AddHeading(h *Heading) {
	t.Direction = h
}

//AddDistance Adds a distance to a Avoid
func (t *Traffic) AddDistance(d *Distance) {
	t.Distance = d
}

//AddAction Adds a Action to a Avoid
func (t *Traffic) AddAction(i interface{}) {
	t.Action = i
}

//AddInfo Adds a Action to a Avoid
func (t *Traffic) AddInfo(s string) {
	t.Info = append(t.Info, s)
}

//Distance is a representation of a Distance
type Distance struct {
	Number float64
	Unit   string
}

//NewDistance is a distance struct
func NewDistance() *Distance {
	return &Distance{}
}

//AddNumber Adds a Number to a avoid
func (d *Distance) AddNumber(i string) {
	n, _ := strconv.ParseFloat(i, 64)
	d.Number = n
}

//AddUnit Adds a Heading to a Avoid
func (d *Distance) AddUnit(s string) {
	d.Unit = s
}

//Crossing is a representation of a Crossing
type Crossing struct {
	From   string
	To     string
	Unit   string
	Number float64
	Info   []string
}

//NewCrossing is a Crossing struct
func NewCrossing() *Crossing {
	return &Crossing{}
}

//AddNumber Adds a Number to a Crossing
func (c *Crossing) AddNumber(i string) {
	n, _ := strconv.ParseFloat(i, 64)
	c.Number = n
}

//AddUnit Adds a Unit to a Crossing
func (c *Crossing) AddUnit(i string) {
	c.Unit = i
}

//AddFrom Adds a From to a Crossing
func (c *Crossing) AddFrom(i string) {
	c.From = i
}

//AddTo Adds a to to a Crossing
func (c *Crossing) AddTo(i string) {
	c.To = i
}

//AddInfo Adds a info to a Crossing
func (c *Crossing) AddInfo(i string) {
	c.Info = append(c.Info, i)
}

//QNH is a representation of a Distance
type QNH struct {
	Pressure float64
}

//NewQNH is a QNH struct
func NewQNH() *QNH {
	return &QNH{}
}

//AddPressure Adds a Pressure to a QNH
func (d *QNH) AddPressure(i string) {
	n, _ := strconv.ParseFloat(i, 64)
	d.Pressure = n
}

//Speed is a representation of a speed
type Speed struct {
	Number float64
	Unit   string
}

//NewSpeed is a Speed struct
func NewSpeed() *Speed {
	return &Speed{}
}

//AddNumber Adds a Number to a Speed
func (s *Speed) AddNumber(i string) {
	n, _ := strconv.ParseFloat(i, 64)
	s.Number = n
}

//AddUnit Adds a Unit to a Speed
func (s *Speed) AddUnit(i string) {
	s.Unit = i
}

//ILS is a representation of a ILS
type ILS struct {
	RunWay    *RunWay
	Direction string
}

//NewILS is a ILS struct
func NewILS() *ILS {
	return &ILS{}
}

//AddRunWay Adds a RunWay to a ILS
func (s *ILS) AddRunWay(r *RunWay) {
	s.RunWay = r
}

//AddDirection Adds a Direction to a ILS
func (s *ILS) AddDirection(i string) {
	s.Direction = i
}

//Maintain is a representation of a Maintain
type Maintain struct {
	Altitude  *Altitude
	Condition *Condition
}

//NewMaintain is a Maintain struct
func NewMaintain() *Maintain {
	return &Maintain{}
}

//AddAltitude Adds a Altitude to a Maintain
func (m *Maintain) AddAltitude(a *Altitude) {
	m.Altitude = a
}

//AddCondition Adds a Condition to a Maintain
func (m *Maintain) AddCondition(c *Condition) {
	m.Condition = c
}

//GlidePathInterception is a glide path interception struct
type GlidePathInterception struct {
}

//NewGlidePathInterception is a Maintain struct
func NewGlidePathInterception() *GlidePathInterception {
	return &GlidePathInterception{}
}

//Approach is a Approach struct
type Approach struct {
}

//NewApproach is a Maintain struct
func NewApproach() *Approach {
	return &Approach{}
}

//Land is a Land struct
type Land struct {
	RunWay    *RunWay
	Wind      *Wind
	Direction string
}

//NewLand is a Maintain struct
func NewLand() *Land {
	return &Land{}
}

//AddRunWay Adds a Runway to a Land
func (l *Land) AddRunWay(r *RunWay) {
	l.RunWay = r
}

//AddWind Adds a Wind to a Land
func (l *Land) AddWind(w *Wind) {
	l.Wind = w
}

//AddDirection Adds a Direction to a Land
func (l *Land) AddDirection(s string) {
	l.Direction = s
}

//Wind is a Wind struct
type Wind struct {
	Heading *Heading
	Speed   *Speed
}

//NewWind is a Maintain struct
func NewWind() *Wind {
	return &Wind{}
}

//AddHeading Adds a Heading to a Wind
func (l *Wind) AddHeading(h *Heading) {
	l.Heading = h
}

//AddSpeed Adds a Speed to a Wind
func (l *Wind) AddSpeed(s *Speed) {
	l.Speed = s
}

//GoAround is a GoAround struct
type GoAround struct {
	Heading *Heading
	Speed   *Speed
}

//NewGoAround is a GoAround struct
func NewGoAround() *GoAround {
	return &GoAround{}
}

//RadarVectors is a RadarVectors struct
type RadarVectors struct {
	ILS    *ILS
	RunWay *RunWay
}

//NewRadarVectors is a RadarVectors struct
func NewRadarVectors() *RadarVectors {
	return &RadarVectors{}
}

//AddILS Adds a ILS to a RadarVectors
func (r *RadarVectors) AddILS(i *ILS) {
	r.ILS = i
}

//AddRunWay Adds a Runway to a RadarVectors
func (r *RadarVectors) AddRunWay(run *RunWay) {
	r.RunWay = run
}

//Squawk is a Squawk struct
type Squawk struct {
	Number   float64
	SlotTime *SlotTime
}

//NewSquawk is a Squawk struct
func NewSquawk() *Squawk {
	return &Squawk{}
}

//AddNumber Adds a Number to a Squawk
func (r *Squawk) AddNumber(i string) {
	n, _ := strconv.ParseFloat(i, 64)
	r.Number = n
}

//AddSlotTime Adds a Slot Time to a Squawk
func (r *Squawk) AddSlotTime(s *SlotTime) {
	r.SlotTime = s
}

//SlotTime is a SlotTime struct
type SlotTime struct {
	Number float64
}

//NewSlotTime is a SlotTime struct
func NewSlotTime() *SlotTime {
	return &SlotTime{}
}

//AddNumber Adds a Number to a SlotTime
func (r *SlotTime) AddNumber(i string) {
	n, _ := strconv.ParseFloat(i, 64)
	r.Number = n
}

//Amendment is a Amendment struct
type Amendment struct {
	Action interface{}
}

//NewAmendment is a Amendment struct
func NewAmendment() *Amendment {
	return &Amendment{}
}

//AddAction Adds a Action to a Amendment
func (r *Amendment) AddAction(i interface{}) {
	r.Action = i
}
