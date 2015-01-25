package main

import "log"

type FamilyBunkAction struct{}

func (action FamilyBunkAction) GetButtonTitle() string {
	return "Stage 1"
}

func (action FamilyBunkAction) GetLayoutName() string {
	return "main_layout"
}

type ScreenText struct {
	Description string
	Button      map[int]string
}

func (action FamilyBunkAction) Run() {
	topTitle := "Family bunk"

	topDescription := `You're in your family bunk. You see 春. The door to Imaginarium is closed.`

	text := map[int]ScreenText{
		0: ScreenText{
			Button: map[int]string{
				0: `"Hey, dad! The monster was sooo scary!"`,
			},
		},
		1: ScreenText{
			Description: `
"Huh?"

Dad turns on TV. You hear how presenter of the TV program tells you about new planet, detected at the edge of the ship's radars range.`,
			Button: map[int]string{
				0: `"It tried to hurt me! But I was brave."`,
				1: `"I was so scared!"`,
			},
		},
		2: ScreenText{
			Description: `
Dad is watching TV.

"Go see your mother. She needs to check you after hibernation."

Presenter continues to talk about all the possibilities that awaits you and 3'000 settlers on new planet.`,
			Button: map[int]string{
				0: `Ignore TV and dad`,
			},
		},
		3: ScreenText{
			Button: map[int]string{
				0: `Caress 春`,
				1: `Disregard 春`,
			},
		},
		4: ScreenText{
			Button: map[int]string{
				0: `Go outside`,
			},
		},
	}

	answerFive := Scenarios{
		&Scenario{
			Title: text[4].Button[0],
			PreDraw: func(scenario *Scenario) bool {
				log.Printf("!!! %#v\n", "lobby one")
				//replace recursion with next location
				game.SetLocation(locations["family_bunk"])
				game.SwitchLocation()
				return false
			},
		},
	}

	answersFour := Scenarios{
		&Scenario{
			Title: text[3].Button[0],
			PreDraw: func(scenario *Scenario) bool {
				log.Printf("!!! %#v\n", "put murr and huptic touch here")
				return true
			},
			Answers: answerFive,
		},

		&Scenario{
			Title: text[3].Button[1],
			PreDraw: func(scenario *Scenario) bool {
				log.Printf("!!! %#v\n", "bad choice for you")
				return true
			},
			Answers: answerFive,
		},
	}

	answerThree := Scenarios{
		&Scenario{
			Title:   text[2].Button[0],
			Answers: answersFour,
		},
	}

	answersTwo := Scenarios{
		&Scenario{
			Title:       text[1].Button[0],
			Description: text[2].Description,
			Answers:     answerThree,
		},
		&Scenario{
			Title:       text[1].Button[1],
			Description: text[2].Description,
			Answers:     answerThree,
		},
	}

	answerOne := Scenarios{
		&Scenario{
			Title:       text[0].Button[0],
			Description: text[1].Description,
			Answers:     answersTwo,
		},
	}

	stageOneScenario := &Scenario{
		Title:       topTitle,
		Description: topDescription,
		Answers:     answerOne,
	}

	stageOneScenario.Draw()
}
