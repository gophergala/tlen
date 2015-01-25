package main

import (
	"log"
	"math"

	"github.com/seletskiy/go-android-rpc/android"
	"github.com/seletskiy/go-android-rpc/android/sdk"
	"github.com/zazab/zhash"
)

type TalkToFatherAction struct {
	minX          float64
	minY          float64
	minZ          float64
	maxX          float64
	maxY          float64
	maxZ          float64
	tick          int64
	descView      sdk.TextView
	wakeUpSuccess bool
}

func (action TalkToFatherAction) GetButtonTitle() string {
	return "Talk to father"
}

func (action TalkToFatherAction) GetLayoutName() string {
	return "main_layout"
}

type ScreenText struct {
	Description string
	Button      map[int]string
}

func (action TalkToFatherAction) Run() {
	desc := game.CreateView("android.widget.TextView").(sdk.TextView)
	desc.SetText1s("YOU MUST MOVE YOUR PHONE FOR SAFE YOUR FATHER")
	desc.SetTextSize(50.0)
	action.descView = desc
	game.AttachView(desc.View)

	sensors := zhash.HashFromMap(android.GetSensorsList())
	accelerometerId, err := sensors.GetString(
		"sensors", "TYPE_ACCELEROMETER",
	)
	if err != nil {
		panic(err)
	}

	android.SubscribeToSensorValues(
		accelerometerId,
		TalkToFatherAccelerometerHandler{&action},
	)
}

func (action TalkToFatherAction) Talk() {
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
				locations["bunk"].UnlinkAction(action)
				game.SetLocation(locations["bunk"])
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

	talkScenario := &Scenario{
		Title:       topTitle,
		Description: topDescription,
		Answers:     answerOne,
	}

	talkScenario.Draw()
}

type TalkToFatherAccelerometerHandler struct {
	action *TalkToFatherAction
}

func (handler TalkToFatherAccelerometerHandler) OnChange(values []float64) {
	handler.action.OnChangeAccelerometerData(values)
}

func (handler TalkToFatherAccelerometerHandler) OnAccuracyChange() {
	//not usable
	log.Printf("%#v\n", "accuracy change!!")
}

func (action *TalkToFatherAction) OnChangeAccelerometerData(values []float64) {
	if action.wakeUpSuccess {
		return
	}

	x := values[0]
	y := values[1]
	z := values[2]

	action.tick++

	if action.tick > 5 {
		action.maxX = 0
		action.maxY = 0
		action.maxZ = 0
		action.minX = 0
		action.maxY = 0
		action.minZ = 0

		action.tick = 0
	}

	if x < action.minX {
		action.minX = x
	}
	if y < action.minY {
		action.minY = y
	}
	if z < action.minZ {
		action.minZ = z
	}

	if x > action.maxX {
		action.maxX = x
	}
	if y > action.maxY {
		action.maxY = y
	}
	if z > action.maxZ {
		action.maxZ = z
	}

	if math.Abs(action.maxX)+math.Abs(action.minX) > 20 ||
		math.Abs(action.maxY)+math.Abs(action.minY) > 20 ||
		math.Abs(action.maxZ)+math.Abs(action.maxZ) > 20 {
		//log.Printf("%#v\n", "YOU WIN")
		action.wakeUpSuccess = true
		action.Talk()
	}

	//log.Printf("%s %s %s | %s %s %s\n",
	//    action.minX, action.minY, action.minZ,
	//    action.maxX, action.maxY, action.maxZ)
}
