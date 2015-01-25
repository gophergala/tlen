package main

import "log"

type StageOneFamilyBunkOne struct{}

func (action StageOneFamilyBunkOne) GetButtonTitle() string {
	return "Stage 1"
}

func (action StageOneFamilyBunkOne) GetLayoutName() string {
	return "main_layout"
}

func (action StageOneFamilyBunkOne) Run() {

	answerFive := Scenarios{
		&Scenario{
			Title: `Go outside`,
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
			Title: `Caress 春`,
			PreDraw: func(scenario *Scenario) bool {
				log.Printf("!!! %#v\n", "put murr and huptic touch here")
				return true
			},
			Answers: answerFive,
		},
		&Scenario{
			Title: `Disregard 春`,
			PreDraw: func(scenario *Scenario) bool {
				log.Printf("!!! %#v\n", "bad choice for you")
				return true
			},
			Answers: answerFive,
		},
	}

	answerThree := Scenarios{
		&Scenario{
			Title:       `Ignore TV and dad`,
			Description: ``,
			Answers:     answersFour,
		},
	}

	answersTwoDescription := `Dad is watching TV.

"Go see your mother. She needs to check you after hibernation."
Presenter continues to talk about all the possibilities that awaits you and 3'000 settlers on new planet.`

	answersTwo := Scenarios{
		&Scenario{
			Title:       `"It tried to hurt me! But I was brave."`,
			Description: answersTwoDescription,
			Answers:     answerThree,
		},
		&Scenario{
			Title:       `"I was so scared!"`,
			Description: answersTwoDescription,
			Answers:     answerThree,
		},
	}

	answerOne := Scenarios{
		&Scenario{
			Title: `"Hey, dad! The monster was sooo scary!"`,
			Description: `"Huh?"

Dad turns on TV. You hear how presenter of the TV program tells you about new planet, detected at the edge of the ship's radars range.`,
			Answers: answersTwo,
		},
	}

	stageOneScenario := &Scenario{
		Title:       "Family bunk",
		Description: `You're in your family bunk. You see 春. The door to Imaginarium is closed.`,
		Answers:     answerOne,
	}

	stageOneScenario.Draw()
}
