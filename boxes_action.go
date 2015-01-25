package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"
)

type BoxesAction struct {
}

func (action BoxesAction) GetLayoutName() string {
	return "boxes_layout"
}

func (action BoxesAction) GetButtonTitle() string {
	return "Game of Boxes"
}

var boxesScenario = Scenario{
	Title:       "",
	Description: "Where is a ball?",
}

func (action BoxesAction) Run() {
	boxesScenario.Answers = Scenarios{}
	//game.ClearViews()

	rand.Seed(time.Now().Unix())
	win := rand.Intn(2) + 1

	for i := 1; i <= 3; i++ {
		scenario := Scenario{
			Title: strconv.Itoa(i),
			PreDraw: func(scenario *Scenario) bool {
				game.ClearViews()
				return true
			},
			Answers: Scenarios{
				&Scenario{
					Title: "Exit",
					PreDraw: func(scenario *Scenario) bool {
						log.Printf("boxes_action.go:42 %#v", 1)
						game.ClearViews()
						log.Printf("boxes_action.go:44 %#v", 2)
						game.RestoreMainLayout()
						log.Printf("boxes_action.go:46 %#v", 3)
						game.SwitchLocation()
						//game.SwitchLayout()
						//game.SwitchLocation()
						return false
					},
				},
			},
		}

		if i != win {
			scenario.Description = "Wrong."
		} else {
			scenario.Description = "YOU WIN!!1"
		}

		boxesScenario.Answers = append(boxesScenario.Answers, &scenario)
	}

	boxesScenario.Draw()
}
