package main

type PickLockTwoAction struct{}

func (action PickLockTwoAction) GetButtonTitle() string {
	return "Pick lock"
}

func (action PickLockTwoAction) GetLayoutName() string {
	return "main_layout"
}

func (action PickLockTwoAction) Run() {
	topTitle := "Todo"
	topDescription := `Todo`

	text := map[int]ScreenText{
		0: ScreenText{
			Button: map[int]string{
				0: `Play`,
				1: `Put away`,
			},
		},
	}

	answersOne := Scenarios{
		&Scenario{
			Title:       text[0].Button[0],
			Description: text[0].Description,
			PreDraw: func(scenario *Scenario) bool {
				// haptic
				game.state.GameProgress += 1

				// should work, see main.go
				//if game.state.GameProgress == 3 {
				game.SetLocation(locations["lobby2"])
				game.SwitchLocation()
				//}

				return false
			},
		},
		&Scenario{
			Title:       text[0].Button[1],
			Description: text[0].Description,
			PreDraw: func(scenario *Scenario) bool {
				game.SetLocation(locations["bunk1"])
				game.SwitchLocation()
				return false
			},
		},
	}

	catScenario := &Scenario{
		Title:       topTitle,
		Description: topDescription,
		Answers:     answersOne,
	}

	catScenario.Draw()
}
