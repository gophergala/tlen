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

var globalLocations = map[string]Location{
	"home":    NewTestLocation("home", "home desc"),
	"kitchen": NewTestLocation("kitchen", "kitchen desc"),
	"outside": NewTestLocation("outside", "outside desc"),
	"shop":    NewTestLocation("shop", "shop desc"),
	"monster": &MonsterSubgame{},
}

func init() {
	globalLocations["home"].Link(globalLocations["kitchen"])
	globalLocations["home"].Link(globalLocations["outside"])

	globalLocations["kitchen"].Link(globalLocations["home"])

	globalLocations["outside"].Link(globalLocations["shop"])
	globalLocations["outside"].Link(globalLocations["home"])
	globalLocations["outside"].Link(globalLocations["kitchen"])

	globalLocations["shop"].Link(globalLocations["outside"])

	globalLocations["home"].Link(globalLocations["monster"])
}

func start() {
	game := NewGame(&State{})
	game.SetLocation(globalLocations["kitchen"])
	game.SetLayoutName("main_layout")
	game.Start()
}

func main() {
	defer android.PanicHandler()

	android.OnStart(start)
	android.Enter()
}
