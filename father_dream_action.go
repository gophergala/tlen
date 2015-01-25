package main

import (
	"log"

	"github.com/seletskiy/go-android-rpc/android"
	"github.com/seletskiy/go-android-rpc/android/sdk"
	"github.com/zazab/zhash"
)

type FatherDreamAction struct {
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
}

func (action *FatherDreamAction) Run() {
	desc := game.CreateView("android.widget.TextView").(sdk.TextView)
	desc.SetText1s("YOU MUST MOVE YOUR PHONE FOR SAFE YOUR FATHER")
	desc.SetTextSize(50.0)
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

func (action *FatherDreamAction) OnChangeAccelerometerData(values []float64) {
	log.Printf("%#v\n", values)
}
