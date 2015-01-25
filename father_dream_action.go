package main

import (
	"log"
	"math"

	"github.com/seletskiy/go-android-rpc/android"
	"github.com/seletskiy/go-android-rpc/android/sdk"
	"github.com/zazab/zhash"
)

type FatherDreamAction struct {
	minX     float64
	minY     float64
	minZ     float64
	maxX     float64
	maxY     float64
	maxZ     float64
	tick     int64
	descView sdk.TextView
}

func (action FatherDreamAction) GetLayoutName() string {
	return "dream_layout"
}

func (action FatherDreamAction) GetButtonTitle() string {
	return "Game of FatherDream"
}

type FatherDreamActionAccelerometerHandler struct {
	action *FatherDreamAction
}

func (handler FatherDreamActionAccelerometerHandler) OnChange(values []float64) {
	handler.action.OnChangeAccelerometerData(values)
}

func (handler FatherDreamActionAccelerometerHandler) OnAccuracyChange() {
	//not usable
	log.Printf("%#v\n", "accuracy change!!")
}

func (action *FatherDreamAction) Run() {
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
		FatherDreamActionAccelerometerHandler{action},
	)

}

func (action *FatherDreamAction) afterWakeUp() {
	log.Printf("%#v\n", action.descView)
	action.descView.SetText1s("WINNER")
}

func (action *FatherDreamAction) OnChangeAccelerometerData(values []float64) {
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

	if math.Abs(action.maxX)+math.Abs(action.minX) > 30 ||
		math.Abs(action.maxY)+math.Abs(action.minY) > 30 ||
		math.Abs(action.maxZ)+math.Abs(action.maxZ) > 30 {
		log.Printf("%#v\n", "YOU WIN")
		action.afterWakeUp()
	}

	log.Printf("%s %s %s | %s %s %s\n",
		action.minX, action.minY, action.minZ,
		action.maxX, action.maxY, action.maxZ)
}
