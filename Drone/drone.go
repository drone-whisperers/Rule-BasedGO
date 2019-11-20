package drone

import (
	"fmt"
	"sync"

	"github.com/Rule-BasedGO/lexer"
	"github.com/Rule-BasedGO/utils"
	"github.com/kr/pretty"
	lex "github.com/timtadh/lexmachine"
)

//Drone represents a drone
type Drone struct {
	name         string
	currentLexer *lex.Lexer
	classifier   *utils.Classifier
	mux          sync.Mutex
	Tasks        []interface{}
}

//CreateNewDrone Returns a new drone
func CreateNewDrone(name string) *Drone {
	lexer, err := lexer.InitLexer(name)
	if err != nil {
		return nil
	}
	classifier := utils.NewClassifier(lexer)

	return &Drone{
		name:         name,
		currentLexer: lexer,
		classifier:   classifier,
	}
}

func (d *Drone) updateLex() error {
	lexer, err := lexer.InitLexer(d.name)
	if err != nil {
		return err
	}
	d.mux.Lock()
	d.currentLexer = lexer
	d.mux.Unlock()
	d.updateClassifier(lexer)
	return nil
}

func (d *Drone) updateClassifier(lexer *lex.Lexer) {
	d.mux.Lock()
	d.classifier = utils.NewClassifier(lexer)
	d.mux.Unlock()
}

//ClassifyStatements opens steam to allow for classification of statements
func (d *Drone) ClassifyStatements(statementChan chan string, stopchan chan struct{}) {
	for {
		select {
		case statement := <-statementChan:
			task, err := d.classifier.Classify(statement)
			if err == nil {
				fmt.Printf("DRONE: %s\nReceived command: %s\n", d.name, statement)
				fmt.Println("Classified Statement:")
				fmt.Printf("%# v\n", pretty.Formatter(task))
				d.Tasks = append(d.Tasks, task)
				fmt.Println("DRONES CURRENT QUEUE:")
				fmt.Printf("%# v\n", pretty.Formatter(d.Tasks))
			}
		case <-stopchan:
			return
		}
	}
}
