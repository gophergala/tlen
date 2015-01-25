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
	"home":        NewTestLocation("home", "home desc"),
	"kitchen":     NewTestLocation("kitchen", "kitchen desc"),
	"outside":     NewTestLocation("outside", "outside desc"),
	"shop":        NewTestLocation("shop", "shop desc"),
	"family_bunk": NewTestLocation("family bunk", "just another bunkroom"),
}

var actions = map[string]Action{
	"boxes":                 &BoxesAction{},
	"monster":               &MonsterAction{},
	"father_dream":          &FatherDreamAction{},
	"stage_1_family_bunk_1": &StageOneFamilyBunkOne{},
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
	locations["kitchen"].LinkAction(actions["monster"])
	locations["kitchen"].LinkAction(actions["father_dream"])

	// Story started here
	locations["family_bunk"].LinkAction(actions["stage_1_family_bunk_1"])
}

var game *Game

func start() {
	//locationButtons = []sdk.Button{}
	//actionButtons = []sdk.Button{}

	//origin := Locations["game_room"]
	//origin.Draw()
	game = &Game{}
	game.SetLocation(locations["family_bunk"])
	game.SetCurrentStage(1)
	game.Start()
}

func main() {
	defer android.PanicHandler()

	android.OnStart(start)
	android.Enter()
}
