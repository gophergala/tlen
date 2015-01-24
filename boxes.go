package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/seletskiy/go-android-rpc/android/sdk"
)

var boxesScenario = SubAction{
	Title:       "",
	Description: "Where is a ball?",
}

func InitPlayBoxes(action *Action, button sdk.Button) {
	rand.Seed(time.Now().Unix())
	win := rand.Intn(2) + 1

	for i := 1; i <= 3; i++ {
		subaction := SubAction{
			Title:       strconv.Itoa(i),
			Answers : SubActions{
				&SubAction{
					Title: "Exit",
					PreDraw: func(subaction *SubAction) bool {
						origin := Locations["home"]
						origin.Draw()
						return false
					},
				},
			},
		}

		if i != win {
			subaction.Description = "Wrong."
		} else {
			subaction.Description = "YOU WIN!!1"
		}

		boxesScenario.Answers = append(boxesScenario.Answers, &subaction)
	}

	boxesScenario.Draw()
}
