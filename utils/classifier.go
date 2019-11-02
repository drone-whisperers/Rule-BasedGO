package utils

import "strings"

const towerName = "metro Tower"
const droneName = "big jet 345"

var actionTokens = map[string]bool{"taxi": true}
var holdPoints = []string{"hold", "hold position", "hold short of"}

type actionToken struct {
	action string
	start  int
	end    int
}

type Taxi struct {
	holdingPoint string
	holdPostion  string
	runway       Runway
}
type Runway struct {
	number int
}

func classify(s string) map[string]string {
	sSlice := strings.Split(s, " ")
	var actionList []interface{}
	var tokenSlice map[string]interface{}
	for _, word := range sSlice {
		if word == towerName {
			tokenSlice[word] = "tower"
		}
		if word == droneName {
			tokenSlice[word] = "drone"
		}
		actionName := actionClassify(word)
		var action interface{}
		switch actionName {
		case "taxi":
			taxiAction := &Taxi{}
			taxiClassify(taxiAction)
			action = taxiAction

		case "":
			action = nil
		}
		if action != nil {
			actionList = append(actionList, action)
		}
	}
	return nil
}

func actionClassify(s string) string {
	if actionTokens[s] {
		return s
	}
	return ""
}

func taxiClassify(a *Taxi) {

}

func ConditionClassify() {

}
