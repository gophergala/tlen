package main

import (
	"github.com/seletskiy/go-android-rpc/android"

	// required for linking
	_ "github.com/seletskiy/go-android-rpc/android/modules"
	"github.com/seletskiy/go-android-rpc/android/sdk"
)

const (
	ViewInvisible = 4
	ViewVisible   = 0
	ViewGone      = 8
)

type State struct {
	Location Location
	Action   Action
}

var HeaderTextView sdk.TextView
var DescTextView sdk.TextView

var locations = map[string]Location{
	"game_room": &GameRoomLocation{},
	"kitchen":   &KitchenLocation{},
}

var actions = map[string]Action{
	"boxes": &BoxesAction{},
}

func init() {
	locations["game_room"].LinkLocation(locations["kitchen"])
	locations["game_room"].LinkAction(actions["boxes"])
}

////var actions = map[string]*Action{
////    "game_boxes": &GameBoxesAction{},
////}

//var actions = map[string]*Action{
//    "game_boxes": &Action{
//        Header:      "Dream",
//        Description: "for testing",
//        Callback: func(action *Action, button sdk.Button) {
//            InitPlayBoxes(action, button)
//        },
//    },
//}

func start() {
	//locationButtons = []sdk.Button{}
	//actionButtons = []sdk.Button{}
	//HeaderTextView = android.GetViewById(
	//    "main_layout", "header_text").(sdk.TextView)
	//DescTextView = android.GetViewById(
	//    "main_layout", "desc_text").(sdk.TextView)

	//origin := Locations["game_room"]
	//origin.Draw()
}

func main() {
	android.OnStart(start)
	android.Enter()
}
