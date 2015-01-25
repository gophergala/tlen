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
	"bunk": NewTestLocation("family bunk", "just another bunkroom"),
}

var actions = map[string]Action{
	// game actions
	"talk_to_father": TalkToFatherAction{},
}

func init() {
	locations["bunk"].LinkAction(actions["talk_to_father"])
}

var game *Game

func start() {
	game = &Game{}
	game.SetLocation(locations["bunk"])
	game.SetCurrentStage(1)
	game.Start()
}

func main() {
	defer android.PanicHandler()

	android.OnStart(start)
	android.Enter()
}
