package main

import (
	"github.com/seletskiy/go-android-rpc/android"

	// required for linking
	_ "github.com/seletskiy/go-android-rpc/android/modules"
)

const (
	ViewInvisible = 4
	ViewVisible   = 0
	ViewGone      = 8
)

var locations = map[string]Location{
	"home":    NewTestLocation("home", "home desc"),
	"kitchen": NewTestLocation("kitchen", "kitchen desc"),
	"outside": NewTestLocation("outside", "outside desc"),
	"shop":    NewTestLocation("shop", "shop desc"),
}

var actions = map[string]Action{
	"boxes": &BoxesAction{},
}

func init() {
	locations["home"].LinkLocation(locations["kitchen"])
	locations["home"].LinkLocation(locations["outside"])

	locations["kitchen"].LinkLocation(locations["home"])

	locations["outside"].LinkLocation(locations["shop"])
	locations["outside"].LinkLocation(locations["home"])
	locations["outside"].LinkLocation(locations["kitchen"])

	locations["shop"].LinkLocation(locations["outside"])

	// DANGEROUS
	locations["kitchen"].LinkAction(actions["boxes"])
}

var game *Game

func start() {
	//locationButtons = []sdk.Button{}
	//actionButtons = []sdk.Button{}

	//origin := Locations["game_room"]
	//origin.Draw()
	game = &Game{}
	game.SetLocation(locations["home"])
	game.Start()
}

func main() {
	defer android.PanicHandler()

	android.OnStart(start)
	android.Enter()
}
