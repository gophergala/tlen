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

func initLocations() map[string]Location {
	locations := make(map[string]Location)
	locations["bunk"] = NewTestLocation("Go to bunk", "Your place")
	locations["lobby"] = NewTestLocation("Go to lobby", "You enter the lobby")
	locations["din"] = NewTestLocation("Go to dinnary", "Dinnay")
	locations["med"] = NewTestLocation("Go to medical", "Medical cabinet")
	locations["cap"] = NewTestLocation("Go to captain room", "Captain's room")

	locations["stub"] = &StubSubgame{}
	locations["woman1"] = &Woman1Subgame{}
	locations["mother1"] = &Mother1Subgame{}
	locations["captain1"] = &Captain1Subgame{}
	locations["cook1"] = &Cook1Subgame{}

	locations["monster"] = &MonsterSubgame{
		NextLocation: locations["bunk"],
	}

	locations["caress_cat"] = &CaressCatSubgame{
		OriginLocation: locations["bunk"],
	}

	locations["wakeup_father_subgame"] = &WakeUpFatherSubgame{
		NextLocation: &JumpLocation{
			Jump: locations["bunk"],
			BaseLocation: BaseLocation{
				ButtonTitle: "Wake up",
			},
		},
	}

	locations["bunk"].Link(locations["caress_cat"])
	locations["bunk"].Link(locations["wakeup_father_subgame"])
	locations["bunk"].Link(locations["lobby"])

	//locations["lobby"].Link(locations["woman1"])
	locations["lobby"].Link(locations["bunk"])
	locations["lobby"].Link(locations["din"])
	locations["lobby"].Link(locations["med"])
	locations["lobby"].Link(locations["cap"])

	//locations["med"].Link(actions["mother1"])
	//locations["med"].Link(actions["captain1"])
	locations["med"].Link(locations["lobby"])

	//locations["din"].Link(actions["cook1"])
	locations["din"].Link(locations["lobby"])

	locations["cap"].Link(locations["lobby"])

	//locations["bunk"].Link(locations["father2"])

	//locations["lobby"].Link(locations["woman2"])

	//locations["med"].Link(locations["pick_lock"])

	//locations["din"].Link(locations["mother2"])
	//locations["din"].Link(locations["cook2"])
	//if !game.IsCaptainAbused()
	//locations["bunk"].Link(locations["father3"])
	//}

	//if game.ScalpelIsStolen {
	//locations["lobby"].Link(locations["woman3"])
	//}

	//locations["med"].Link(locations["mother3"])
	//locations["din"].Link(locations["cook3"])
	//locations["cap"].Link(locations["explore_captain_room"])

	return locations
}

func start() {
	locations := initLocations()

	game := NewGame(&State{})
	game.SetLocation(locations["monster"])
	game.SetLayoutName("main_layout")
	game.Start()
}

func main() {
	defer android.PanicHandler()

	android.OnStart(start)
	android.Enter()
}
