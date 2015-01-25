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
			Title: `End`,
			//Description: `Father takes your away from Imaginarium and locks the door on magnet lock.`,
			PreDraw: func(scenario *Scenario) bool {
				scenario.Description = "Father takes your away from Imaginarium and locks the door on <b>magnet lock</b>."
				game.SetCurrentStage(1)
				scenario.Draw()
				return false
			},
		},
	}

	answerTen := Scenarios{
		&Scenario{
			Title: `"Are we <b>home</b>?"`,
			Description: `<p>"Not yet. But it looks like we are so close".</p>
<p>Grin touches your father's lips.</p>`,
			Answers: answerEleven,
		},
	}

	answerNine := Scenarios{
		&Scenario{
			Title:       `"I feel pain!"`,
			Description: `<p>"You'd better see your mom. She's our lovely doctor."</p>`,
			Answers:     answerTen,
		},
	}

	answerEight := Scenarios{
		&Scenario{
			Title:       `I've fought with a monster!`,
			Description: `<p>Your father puts you on your feet.</p>`,
			Answers:     answerNine,
		},
	}

	answerSeven := Scenarios{
		&Scenario{
			Title: "Cry",
			Description: `<p>Your father is looking at you inquiringly.</p>
<p>"It's a just bad dream.", he says.</p>`,
			Answers: answerEight,
		},
	}

	answerSix := Scenarios{
		&Scenario{
			Title: "Anger",
			Description: `<p>You feel dark presence goes away.</p>
<p>You suddenly became aware of your body.</p>
<p>Rumbling around you. Low and pleasant rumbling. G-r-r-r-r. Like a tiger. Like a... an engines. Ginormous photon engines.</p>
<p>You're waking up.</p>
<p>You're seing you father. You're feeling bright <font color="red">pain</font>.</p>`,
			Answers: answerSeven,
		},
	}

	answerFive := Scenarios{
		&Scenario{
			Title:       "Panic",
			Description: `<p>You feel <font color="red">pain</font>.</p>`,
			Answers:     answerSix,
		},
	}

	answerFour := Scenarios{
		&Scenario{
			Title:       "Anger",
			Description: "<p>You feel dark presence goes away.</p>",
			Answers:     answerFive,
		},
	}

	answersThreeDescription := `<p>"Alive? Energy? Energy."</p>
<p>You start to feel yourself more clearly.</p>
<p>Start to feel something around you. And you realize the something is alive.</p>
<p>You don't need another eternity to understand that the something is deadly strange, deadly dark and deadly dangerous. You feel it doesn't even know who are you, but even more you know it want you not to be.</p>`

	answersThreeScenarios := Scenarios{
		&Scenario{
			Title:       "Run away",
			Description: "<p>It is pitch black. There is nowhere to run.</p>",
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
			Description: `<p>"You? Things? Cold. Solid. Energy. Dark."</p>`,
			Answers:     answersThree,
		},
	}

	answerOne := Scenarios{
		&Scenario{
			Title:       "Try to move",
			Description: `<p>You feel different. Time goes different.</p><p>You feel your mind wandering.</p>`,
			Answers:     answerTwo,
		},
	}

	monsterScenario := &Scenario{
		Title: "Monster scenario",
		Description: `<p><b>It is pitch black.</b></p>
<p>You try to make a move to escape, but barely feel your body.</p>
<p>Just powerful clammy darkness and strange feeling of existence and non-existence at the same time.</p>`,
		Answers: answerOne,
	}

	monsterScenario.Draw()
}
