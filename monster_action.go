package main

type MonsterAction struct{}

func (action MonsterAction) GetButtonTitle() string {
	return "Monster Title"
}

func (action MonsterAction) GetLayoutName() string {
	return "main_layout"
}

var rootScenario = Scenario{}

func (action MonsterAction) Run() {

	answerEleven := Scenarios{
		&Scenario{
			Title:       `End`,
			Description: `Father takes your away from Imaginarium and locks the door on magnet lock.`,
			PreDraw: func(scenario *Scenario) bool {
				scenario.Description = "Father takes your away from Imaginarium and locks the door on magnet lock."
				scenario.Draw()
				return false
			},
		},
	}

	answerTen := Scenarios{
		&Scenario{
			Title:       `"Are we home?"`,
			Description: `"Not yet. But it looks like we are so close". Grin touches your father's lips.`,
			Answers:     answerEleven,
		},
	}

	answerNine := Scenarios{
		&Scenario{
			Title:       `"I feel pain!"`,
			Description: `"You'd better see your mom. She's our lovely doctor."`,
			Answers:     answerTen,
		},
	}

	answerEight := Scenarios{
		&Scenario{
			Title:       `I've fought with a monster!`,
			Description: `Your father puts you on your feet.`,
			Answers:     answerNine,
		},
	}

	answerSeven := Scenarios{
		&Scenario{
			Title: "Cry",
			Description: `Your father is looking at you inquiringly.
"It's a just bad dream.", he says.`,
			Answers: answerEight,
		},
	}

	answerSix := Scenarios{
		&Scenario{
			Title: "Anger",
			Description: `You feel dark presence goes away.

You suddenly became aware of your body. Rumbling around you. Low and pleasant rumbling. G-r-r-r-r. Like a tiger. Like a... an engines. Ginormous photon engines.
You're waking up.

You're seing you father. You're feeling bright pain.`,
			Answers: answerSeven,
		},
	}

	answerFive := Scenarios{
		&Scenario{
			Title:       "Panic",
			Description: "You feel pain.",
			Answers:     answerSix,
		},
	}

	answerFour := Scenarios{
		&Scenario{
			Title:       "Anger",
			Description: "You feel dark presence goes away.",
			Answers:     answerFive,
		},
	}

	answersThreeDescription := `"Alive? Energy? Energy."

You start to feel yourself more clearly. Start to feel something around you. And you realize the something is alive.

You don't need another eternity to understand that the something is deadly strange, deadly dark and deadly dangerous. You feel it doesn't even know who are you, but even more you know it want you not to be.`

	answersThreeScenarios := Scenarios{
		&Scenario{
			Title:       "Run away",
			Description: "It is pitch black. There is nowhere to run.",
			Answers:     answerFour,
		},
	}

	answersThree := Scenarios{
		&Scenario{
			Title:       "Play.",
			Description: answersThreeDescription,
			Answers:     answersThreeScenarios,
		},
		&Scenario{
			Title:       "Toys.",
			Description: answersThreeDescription,
			Answers:     answersThreeScenarios,
		},
		&Scenario{
			Title:       "Mother.",
			Description: answersThreeDescription,
			Answers:     answersThreeScenarios,
		},
		&Scenario{
			Title:       "Father.",
			Description: answersThreeDescription,
			Answers:     answersThreeScenarios,
		},
	}

	answerTwo := Scenarios{
		&Scenario{
			Title:       `"Who are you?"`,
			Description: `"You? Things? Cold. Solid. Energy. Dark."`,
			Answers:     answersThree,
		},
	}

	answerOne := Scenarios{
		&Scenario{
			Title:       "Try to move",
			Description: `You feel different. Time goes different. You feel your mind wandering.`,
			Answers:     answerTwo,
		},
	}

	monsterScenario := &Scenario{
		Title:       "Monster scenario",
		Description: `It is pitch black. You try to make a move to escape, but barely feel your body. Just powerful clammy darkness and strange feeling of existence and non-existence at the same time.`,
		Answers:     answerOne,
	}

	monsterScenario.Draw()
}
