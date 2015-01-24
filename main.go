package main

import (
	"log"

	"github.com/seletskiy/go-android-rpc/android"

	// required for linking
	_ "github.com/seletskiy/go-android-rpc/android/modules"
	"github.com/seletskiy/go-android-rpc/android/sdk"
)

const (
	viewInvisible = 4
	viewVisible   = 0
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
		Header: "Outside",
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

func (location *Location) Draw() {
	log.Printf("%#v\n", location)
	buttons := []sdk.Button{
		android.GetViewById(
			"main_layout", "choose_button_1").(sdk.Button),
		android.GetViewById(
			"main_layout", "choose_button_2").(sdk.Button),
		android.GetViewById(
			"main_layout", "choose_button_3").(sdk.Button),
	}

	for _, button := range buttons {
		button.SetText1s("")
		button.SetVisibility(viewInvisible)
	}

	for index, locationName := range location.Locations {
		loc := locations[locationName]
			buttons[index].SetText1s(loc.Header)
		buttons[index].SetVisibility(viewVisible)

		android.OnClick(buttons[index], NextButtonHandler{
			loc,
		})
	}

	headerTextView := android.GetViewById(
		"main_layout", "header_text").(sdk.TextView)
	headerTextView.SetText1s(location.Header)

	descTextView := android.GetViewById(
		"main_layout", "desc_text").(sdk.TextView)
	descTextView.SetText1s(location.Description)
}

func start() {
	log.Printf("%#v", locations)
	origin := locations["home"]
	log.Printf("%#v", origin)
	origin.Draw()
}

func main() {
	android.OnStart(start)
	android.Enter()
}
