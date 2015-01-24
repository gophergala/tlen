package main

import (
	"github.com/seletskiy/go-android-rpc/android"
	"github.com/seletskiy/go-android-rpc/android/sdk"
)

type Scenarios []*Scenario

type Scenario struct {
	Title       string
	Opened      func() bool
	Description string
	PreDraw     func(scenario *Scenario) bool
	Answers     Scenarios
}

type ScenarioButtonHandler struct {
	scenario *Scenario
	button   *sdk.Button
}

func (handler ScenarioButtonHandler) OnClick() {
	if handler.scenario.PreDraw == nil || handler.scenario.PreDraw(handler.scenario) {
		handler.scenario.Draw()
	}
}

func (scenarios *Scenarios) Draw() {
	for _, scenario := range *scenarios {
		if scenario.Opened != nil && !scenario.Opened() {
			continue
		}

		button := game.CreateView("android.widget.Button").(sdk.Button)
		button.SetText1s(scenario.Title)
		android.OnClick(button, ScenarioButtonHandler{
			scenario,
			&button,
		})

		game.AttachView(button.View)
	}
}

func (scenario *Scenario) Draw() {
	game.ClearViews()

	desc := game.CreateView("android.widget.TextView").(sdk.TextView)
	desc.SetText1s(scenario.Description)
	game.AttachView(desc.View)

	scenario.Answers.Draw()
}
