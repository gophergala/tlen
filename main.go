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
	locations["bunk"] = &BaseLocation{
		ButtonTitle: "Go to bunk",
		Header:      "<p>It's yours family bunk.</p>",
		Description: "<p>You see " + Cat + ".</p>" +
			"<p>You see your dad watching TV.</p>" +
			"<p>There are door to ship lobby.</p>",
	}

	locations["lobby"] = NewTestLocation("Go to lobby", "You enter the lobby")
	locations["din"] = NewTestLocation("Go to dinnary", "Dinnay")
	locations["med"] = NewTestLocation("Go to medical", "Medical cabinet")
	locations["cap"] = NewTestLocation("Go to captain room", "Captain's room")

	locations["stub"] = &StubSubgame{
		GlobalLocations: locations,
	}

	locations["woman1"] = &Woman1Subgame{
		GlobalLocations: locations,
	}

	locations["mother1"] = &Mother1Subgame{
		GlobalLocations: locations,
	}

	locations["captain1"] = &Captain1Subgame{
		GlobalLocations: locations,
	}

	locations["cook1"] = &Cook1Subgame{
		GlobalLocations: locations,
	}

	locations["monster"] = &MonsterSubgame{
		GlobalLocations: locations,
	}

	locations["caress_cat"] = &CaressCatSubgame{
		OriginLocation: locations["bunk"],
	}

	locations["locked"] = &LockedSubgame{
		GlobalLocations: locations,
	}

	locations["wakeup_father_subgame"] = &WakeUpFatherSubgame{
		NextLocation: locations["bunk"],
	}

	locations["bunk"].Link(locations["caress_cat"])
	locations["bunk"].Link(locations["wakeup_father_subgame"])
	locations["bunk"].Link(locations["lobby"])

	locations["lobby"].Link(
		&Stage1ProxyLocation{
			OldWay: locations["woman1"],
			NewWay: locations["caress_cat"],
		},
	)

	locations["lobby"].Link(locations["bunk"])
	locations["lobby"].Link(locations["din"])
	locations["lobby"].Link(locations["med"])
	locations["lobby"].Link(locations["cap"])
	locations["lobby"].Link(locations["locked"])

	locations["med"].Link(locations["mother1"])
	locations["med"].Link(locations["captain1"])
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
	game.SetLocation(locations["lobby"])
	game.SetLayoutName("main_layout")
	game.Start()
}

func main() {
	defer android.PanicHandler()

	android.OnStart(start)
	android.Enter()
}
