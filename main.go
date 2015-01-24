package main

import (
	"github.com/seletskiy/go-android-rpc/android"

	// required for linking
	_ "github.com/seletskiy/go-android-rpc/android/modules"
	"github.com/seletskiy/go-android-rpc/android/sdk"
)

type Location struct {
	Name        string
	Description string
	Locations   []*Location
}

var LocationShop = Location{
	Name:        "Shop",
	Description: "Welcome to food shop. I have been shoped! Go to home.",
}

var LocationKitchen = Location{
	Name:        "Kitchen",
	Description: "Kitchen. Your kholodilnik is empty",
	Locations: []*Location{
		&LocationShop,
	},
}

var LocationOrigin = Location{
	Name:        "Origin",
	Description: "You can go to kitchen or outside",
	Locations: []*Location{
		&LocationKitchen,
		&LocationShop,
	},
}

type NextButtonHandler struct {
	location *Location
}

func (handler NextButtonHandler) OnClick() {
	handler.location.Draw()
}

func (location *Location) Draw() {
	buttons := []sdk.Button{
		android.GetViewById(
			"main_layout", "choose_button_1").(sdk.Button),
		android.GetViewById(
			"main_layout", "choose_button_2").(sdk.Button),
		android.GetViewById(
			"main_layout", "choose_button_3").(sdk.Button),
	}

	for _, button := range buttons {
		button.SetText1s("[disabled]")
		button.SetEnabled(false)
	}

	for index, loc := range location.Locations {
		buttons[index].SetText1s(loc.Name)
		buttons[index].SetEnabled(true)

		android.OnClick(buttons[index], NextButtonHandler{
			loc,
		})
	}

	nameTextView := android.GetViewById(
		"main_layout", "name_text").(sdk.TextView)
	nameTextView.SetText1s(location.Name)

	descTextView := android.GetViewById(
		"main_layout", "desc_text").(sdk.TextView)
	descTextView.SetText1s(location.Description)
}

func start() {
	LocationShop.Locations = []*Location{
		&LocationOrigin,
	}

	LocationOrigin.Draw()
}

func main() {
	android.OnStart(start)
	android.Enter()
}
