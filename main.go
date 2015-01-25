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

const Cat = `<font color="green">春</font>`

var globalLocations = map[string]Location{
	"bunk":  NewTestLocation("Go to bunk", "Your place"),
	"lobby": NewTestLocation("Go to lobby", "You enter the lobby"),
	"din":   NewTestLocation("Go to dinnary", "Dinnay"),
	"med":   NewTestLocation("Go to medical", "Medical cabinet"),
	"cap":   NewTestLocation("Go to captain room", "Captain's room"),

	"caress_cat":            &CaressCatSubgame{},
	"wakeup_father_subgame": &WakeUpFatherSubgame{},
}

func init() {
	globalLocations["bunk"].Link(globalLocations["caress_cat"])
	globalLocations["bunk"].Link(globalLocations["wakeup_father_subgame"])
	globalLocations["bunk"].Link(globalLocations["lobby"])

	//globalLocations["lobby"].Link(globalLocations["woman1"])
	globalLocations["lobby"].Link(globalLocations["bunk"])
	globalLocations["lobby"].Link(globalLocations["din"])
	globalLocations["lobby"].Link(globalLocations["med"])
	globalLocations["lobby"].Link(globalLocations["cap"])

	//globalLocations["med"].Link(actions["mother1"])
	//globalLocations["med"].Link(actions["captain1"])
	globalLocations["med"].Link(globalLocations["lobby"])

	//globalLocations["din"].Link(actions["cook1"])
	globalLocations["din"].Link(globalLocations["lobby"])

	globalLocations["cap"].Link(globalLocations["lobby"])

	//locations["bunk"].Link(actions["father2"])

	//locations["lobby"].Link(actions["woman2"])

	//locations["med"].Link(actions["pick_lock"])

	//locations["din"].Link(actions["mother2"])
	//locations["din"].Link(actions["cook2"])
	//if !game.IsCaptainAbused()
	//locations["bunk"].Link(actions["father3"])
	//}

	//if game.ScalpelIsStolen {
	//locations["lobby"].Link(actions["woman3"])
	//}

	//locations["med"].Link(actions["mother3"])
	//locations["din"].Link(actions["cook3"])
	//locations["cap"].Link(actions["explore_captain_room"])
}

func start() {
	game := NewGame(&State{})
	game.SetLocation(globalLocations["bunk"])
	game.SetLayoutName("main_layout")
	game.Start()
}

func main() {
	defer android.PanicHandler()

	android.OnStart(start)
	android.Enter()
}
