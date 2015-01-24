package main

import (
	"log"
	"strconv"

	"github.com/seletskiy/go-android-rpc/android"

	// required for linking
	_ "github.com/seletskiy/go-android-rpc/android/modules"
	"github.com/seletskiy/go-android-rpc/android/sdk"
)

const (
	viewInvisible = 4
	viewVisible   = 0
	viewGone      = 8
)

type Location struct {
	Header      string
	Description string
	Locations   []string
	Actions     []string
}

type Action struct {
	Header      string
	Description string
	Callback    func()
}

var actions = map[string]*Action{
	"hideButtons": &Action{
		Header:      "Hide all location buttons",
		Description: "for testing",
		Callback: func() {
			for index, _ := range locationButtons {
				if locationButtons[index].IsShown()["result"] == "true" {
					locationButtons[index].SetVisibility(viewGone)
				}
			}
		},
	},
}

var locations = map[string]*Location{
	"shop": &Location{
		Header:      "Shop",
		Description: "Welcome to food shop. I have been shoped! Go to home.",
		Locations: []string{
			"outside",
		},
	},
	"kitchen": &Location{
		Header:      "Kitchen",
		Description: "Kitchen. Your kholodilnik is empty",
		Locations: []string{
			"outside",
			"home",
		},
		Actions: []string{
			"hideButtons",
		},
	},
	"home": &Location{
		Header:      "Home",
		Description: "You can go to kitchen or outside",
		Locations: []string{
			"kitchen",
			"outside",
		},
	},
	"outside": &Location{
		Header:      "Outside",
		Description: "Hello world!",
		Locations: []string{
			"home",
			"shop",
		},
	},
}

type NextButtonHandler struct {
	location *Location
}

func (handler NextButtonHandler) OnClick() {
	handler.location.Draw()
}

type ActionButtonHandler struct {
	action *Action
}

func (handler ActionButtonHandler) OnClick() {
	handler.action.Callback()
}

var locationButtons []sdk.Button
var actionButtons []sdk.Button

func (location *Location) Draw() {
	log.Printf("%#v\n", location)

	indexActions := 0
	for _, actionName := range location.Actions {
		action := actions[actionName]

		var button sdk.Button
		var isNew bool

		if indexActions >= len(actionButtons) {
			isNew = true
			button = android.CreateView(
				strconv.Itoa(200+indexActions), "android.widget.Button").(sdk.Button)
		} else {
			button = actionButtons[indexActions]

			if button.IsShown()["result"] == "false" {
				button.SetVisibility(viewVisible)
			}
		}

		button.SetText1s(action.Header)

		android.OnClick(button, ActionButtonHandler{
			action,
		})

		if isNew {
			android.AttachView(button, "2130903040")
			actionButtons = append(actionButtons, button)
		}

		indexActions++

	}

	if indexActions < len(actionButtons) {
		for i, _ := range actionButtons[indexActions:] {
			actionButtons[indexActions+i].SetVisibility(viewGone)
		}
	}

	indexLocations := 0
	for _, locationName := range location.Locations {
		loc := locations[locationName]

		var button sdk.Button
		var isNew bool

		if indexLocations >= len(locationButtons) {
			isNew = true
			button = android.CreateView(
				strconv.Itoa(100+indexLocations), "android.widget.Button").(sdk.Button)
		} else {
			button = locationButtons[indexLocations]

			if button.IsShown()["result"] == "false" {
				button.SetVisibility(viewVisible)
			}
		}

		button.SetText1s(loc.Header)

		android.OnClick(button, NextButtonHandler{
			loc,
		})

		if isNew {
			android.AttachView(button, "2130903040")
			locationButtons = append(locationButtons, button)
		}

		indexLocations++
	}

	if indexLocations < len(locationButtons) {
		for i, _ := range locationButtons[indexLocations:] {
			locationButtons[indexLocations+i].SetVisibility(viewGone)
		}
	}

	headerTextView := android.GetViewById(
		"main_layout", "header_text").(sdk.TextView)
	headerTextView.SetText1s(location.Header)

	descTextView := android.GetViewById(
		"main_layout", "desc_text").(sdk.TextView)
	descTextView.SetText1s(location.Description)
}

func start() {
	origin := locations["home"]
	origin.Draw()
}

func main() {
	android.OnStart(start)
	android.Enter()
}
