package main

type TalkToFatherTwoAction struct{}

func (action TalkToFatherTwoAction) GetButtonTitle() string {
	return "Talk to father"
}

func (action TalkToFatherTwoAction) GetLayoutName() string {
	return "main_layout"
}

func (action TalkToFatherTwoAction) Run() {
	topTitle := "Todo"
	topDescription := `Todo`

	text := map[int]ScreenText{
		0: ScreenText{
			Button: map[int]string{
				0: `Talk`,
				1: `Stop talking`,
			},
		},
	}

	answersOne := Scenarios{
		&Scenario{
			Title:       text[0].Button[0],
			Description: text[0].Description,
			PreDraw: func(scenario *Scenario) bool {
				// haptic
				return false
			},
		},
		&Scenario{
			Title:       text[0].Button[1],
			Description: text[0].Description,
			PreDraw: func(scenario *Scenario) bool {
				locations["med1"].UnlinkAction(action)
				game.SetLocation(locations["med1"])
				game.SwitchLocation()
				return false
			},
		},
	}

	scenario := &Scenario{
		Title:       topTitle,
		Description: topDescription,
		Answers:     answersOne,
	}

	scenario.Draw()
}
