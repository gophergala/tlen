package main

import (
	"log"
	"strconv"

	"github.com/seletskiy/go-android-rpc/android"
	"github.com/seletskiy/go-android-rpc/android/sdk"
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

type LocationButtonHandler struct {
	location *Location
}

func (handler LocationButtonHandler) OnClick() {
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

		android.OnClick(button, LocationButtonHandler{
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
