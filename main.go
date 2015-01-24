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
	"game_boxes": &Location{
		Header:      "Boxes",
		Description: "Stupid game",
		Actions: []string{
			"play",
		},
	},
}

var actions = map[string]*Action{
	"monster": &Action{
		Header:      "Dream",
		Description: "for testing",
		Callback: func(action *Action, button sdk.Button) {
			InitDreamWithMonster()
		},
	},
}

func start() {
	HeaderTextView = android.GetViewById(
		"main_layout", "header_text").(sdk.TextView)
	DescTextView = android.GetViewById(
		"main_layout", "desc_text").(sdk.TextView)

	android.CallControlMusicPlayback("start", android.GetResourceById("raw/file"))

	origin := Locations["imaginarium"]
	origin.Draw()
}

func main() {
	android.OnStart(start)
	android.Enter()
}
