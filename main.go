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
	ViewInvisible = 4
	ViewVisible   = 0
	ViewGone      = 8
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
	Callback    func(action *Action, button sdk.Button)
}

var actions = map[string]*Action{
	"monster": &Action{
		Header:      "Dream",
		Description: "for testing",
		Callback: func(action *Action, button sdk.Button) {
			InitDreamWithMonster(action, button)
		},
	},
	"playWithCat": &Action{
		Header:      "Play with cat",
		Description: "",
		Callback: func(action *Action, button sdk.Button) {
			android.OnClick(button, PlayWithCatHandler{button})
		},
	},
}

type PlayWithCatHandler struct {
	button sdk.Button
}

func (handler PlayWithCatHandler) OnClick() {
	handler.button.PerformHapticFeedback(0)
}

var locations = map[string]*Location{
	"bunkroom": &Location{
		Header:      "Bunkroom",
		Description: "You see cat in front of you.",
		Actions: []string{
			"playWithCat",
		},
	},
	"shop": &Location{
		Header:      "Shop",
		Description: "Welcome to food shop. I have been shoped! Go to home.",
		Locations: []string{
			"outside",
		},
		Actions: []string{
			"customer",
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
	"imaginarium": &Location{
		Header:      "Some Imaginarium",
		Description: "Whatever",
		Actions: []string{
			"monster",
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
	button sdk.Button
}

func (handler ActionButtonHandler) OnClick() {
	HeaderTextView.SetText1s(handler.action.Header)
	DescTextView.SetText1s(handler.action.Description)

	hideLocation()

	handler.action.Callback(handler.action, handler.button)
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
				button.SetVisibility(ViewVisible)
			}
		}

		button.SetText1s(action.Header)

		android.OnClick(button, ActionButtonHandler{
			action, button,
		})

		if isNew {
			android.AttachView(button, strconv.Itoa(MainLayoutId))
			actionButtons = append(actionButtons, button)
		}

		indexActions++

	}

	if indexActions < len(actionButtons) {
		for i, _ := range actionButtons[indexActions:] {
			actionButtons[indexActions+i].SetVisibility(ViewGone)
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
				button.SetVisibility(ViewVisible)
			}
		}

		button.SetText1s(loc.Header)

		android.OnClick(button, NextButtonHandler{
			loc,
		})

		if isNew {
			android.AttachView(button, strconv.Itoa(MainLayoutId))
			locationButtons = append(locationButtons, button)
		}

		indexLocations++
	}

	if indexLocations < len(locationButtons) {
		for i, _ := range locationButtons[indexLocations:] {
			locationButtons[indexLocations+i].SetVisibility(ViewGone)
		}
	}

	HeaderTextView.SetText1s(location.Header)
	DescTextView.SetText1s(location.Description)
}

func hideLocation() {
	for index, _ := range actionButtons {
		if actionButtons[index].IsShown()["result"] == "true" {
			actionButtons[index].SetVisibility(ViewGone)
		}
	}

	for index, _ := range locationButtons {
		if locationButtons[index].IsShown()["result"] == "true" {
			locationButtons[index].SetVisibility(ViewGone)
		}
	}
}

var HeaderTextView sdk.TextView
var DescTextView sdk.TextView

const MainLayoutId = 0x7f030000

func start() {
	HeaderTextView = android.GetViewById(
		"main_layout", "header_text").(sdk.TextView)
	DescTextView = android.GetViewById(
		"main_layout", "desc_text").(sdk.TextView)

	origin := locations["imaginarium"]
	origin.Draw()
}

func main() {
	android.OnStart(start)
	android.Enter()
}
