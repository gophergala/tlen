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

var locationButtons []sdk.Button

func (location *Location) Draw() {
	log.Printf("%#v\n", location)

	index := 0
	for _, locationName := range location.Locations {
		loc := locations[locationName]

		var button sdk.Button
		var isNew bool

		if index >= len(locationButtons) {
			isNew = true
			button = android.CreateView(
				strconv.Itoa(100+index), "android.widget.Button").(sdk.Button)
		} else {
			button = locationButtons[index]
			if !button.IsShown() {
				button.SetVisibility(viewVisible)
			}
		}

		button.SetText1s(loc.Header)

		android.OnClick(button, NextButtonHandler{
			loc,
		})

		if isNew {
			android.AttachView(button, "2130837504")
			locationButtons = append(locationButtons, button)
		}

		index++
	}

	if index < len(locationButtons) {
		for i, _ := range locationButtons[index:] {
			locationButtons[index+i].SetVisibility(viewGone)
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
