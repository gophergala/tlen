package main

import (
	"github.com/seletskiy/go-android-rpc/android"

	// required for linking
	_ "github.com/seletskiy/go-android-rpc/android/modules"
	"github.com/seletskiy/go-android-rpc/android/sdk"
)

const MainLayoutId = 0x7f030000

const (
	ViewInvisible = 4
	ViewVisible   = 0
	ViewGone      = 8
)

var HeaderTextView sdk.TextView
var DescTextView sdk.TextView

var Locations = map[string]*Location{
	"game_room": &Location{
		Header: "Boxes",
		Description: "Stupid game",
		Actions: []string{
			"game_boxes",
		},
	},
}

var actions = map[string]*Action{
	"game_boxes": &Action{
		Header:      "Dream",
		Description: "for testing",
		Callback: func(action *Action, button sdk.Button) {
			InitPlayBoxes(action, button)
		},
	},
}

func start() {
	HeaderTextView = android.GetViewById(
		"main_layout", "header_text").(sdk.TextView)
	DescTextView = android.GetViewById(
		"main_layout", "desc_text").(sdk.TextView)

	origin := Locations["game_room"]
	origin.Draw()
}

func main() {
	android.OnStart(start)
	android.Enter()
}
