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
	"dream": NewTestLocation("Enter dream", "dream"),

	"bunk1":            NewTestLocation("Go to bunk", "just another bunkroom"),
	"transition_bunk1": NewTestLocation("Go to bunk", "just another bunkroom"),
	"bunk2":            NewTestLocation("Go to bunk2", "just another bunkroom"),

	"lobby1": NewTestLocation("Go to lobby", "just another bunkroom"),
	"din1":   NewTestLocation("Go to dinary room", "just another bunkroom"),
	"med1":   NewTestLocation("Go to medical room", "just another bunkroom"),
	"cap1":   NewTestLocation("Go to captain room", "just another bunkroom"),
}

var actions = map[string]Action{
	// game actions
	"monster_facing":           MonsterFacingAction{},
	"talk_to_father1":          TalkToFatherOneAction{},
	"play_with_cat1":           PlayWithCatOneAction{},
	"talk_to_woman1":           TalkToWomanOneAction{},
	"talk_to_mother1":          TalkToMotherOneAction{},
	"talk_to_captain1":         TalkToCaptainOneAction{},
	"talk_to_cook1":            TalkToCookOneAction{},
	"approve_transition_bunk1": ApproveTransitionBunkOneAction{},
}

func init() {
	initStageOne()
}

func initStageOne() {
	// implemented, linked to bunk1
	//locations["dream"].LinkAction(actions["monster_facing"])

	locations["bunk1"].LinkAction(actions["talk_to_father1"])
	locations["bunk1"].LinkAction(actions["play_with_cat1"])
	locations["bunk1"].LinkLocation(locations["lobby1"])

	locations["transition_bunk1"].LinkAction(actions["approve_transition_bunk1"])

	locations["lobby1"].LinkAction(actions["talk_to_woman1"])
	locations["lobby1"].LinkLocation(locations["bunk1"])
	locations["lobby1"].LinkLocation(locations["din1"])
	locations["lobby1"].LinkLocation(locations["med1"])
	locations["lobby1"].LinkLocation(locations["cap1"])

	locations["med1"].LinkAction(actions["talk_to_mother1"])
	locations["med1"].LinkAction(actions["talk_to_captain1"])
	locations["med1"].LinkLocation(locations["lobby1"])

	locations["din1"].LinkAction(actions["talk_to_cook1"])
	locations["din1"].LinkLocation(locations["lobby1"])

	locations["cap1"].LinkLocation(locations["lobby1"])

	// stage 2
	locations["bunk2"].LinkLocation(locations["bunk1"]) // should be lobby2
}

func initStageTwo() {
	//    locations["bunk2"].LinkAction(actions["talk_to_father2"])

	//    locations["lobby2"].LinkAction(actions["talk_to_woman2"])
	//    locations["lobby2"].LinkLocation(locations["bunk2"])
	//    locations["lobby2"].LinkLocation(locations["din2"])
	//    locations["lobby2"].LinkLocation(locations["med2"])
	//    locations["lobby2"].LinkLocation(locations["cap2"])

	//    locations["med2"].LinkLocation(locations["lobby2"])
	//    locations["med2"].LinkAction(actions["pick_lock_2"])

	//    locations["din2"].LinkAction(actions["talk_to_mother2"])
	//    locations["din2"].LinkAction(actions["talk_to_cook2"])
	//    locations["din2"].LinkLocation(locations["lobby2"])

	//    locations["cap2"].LinkLocation(locations["lobby2"])
}

func initStageThree() {
	//    if !game.CaptainIsAbused() {
	//        // captain killed father
	//        locations["bunk3"].LinkAction(actions["talk_to_father3"])
	//    }

	//    locations["bunk3"].LinkLocation(locations["lobby3"])

	//    if game.ScalpelIsStolen {
	//        locations["lobby3"].LinkAction(actions["talk_to_woman3"])
	//    }

	//    locations["lobby3"].LinkLocation(locations["bunk3"])
	//    locations["lobby3"].LinkLocation(locations["din3"])
	//    locations["lobby3"].LinkLocation(locations["med3"])
	//    locations["lobby3"].LinkLocation(locations["cap3"])

	//    locations["med3"].LinkLocation(locations["lobby3"])
	//    locations["med3"].LinkAction(actions["talk_to_mother3"])

	//    locations["din3"].LinkAction(actions["talk_to_cook3"])
	//    locations["din3"].LinkLocation(locations["lobby3"])

	//    locations["cap3"].LinkLocation(locations["lobby3"])

	//    locations["cap3"].LinkAction(actions["explore_captain_room3"])
}

var game *Game

func start() {
	game = &Game{}
	game.SetLocation(locations["bunk1"])
	game.SetCurrentStage(1)
	game.Start()
}

func main() {
	defer android.PanicHandler()

	android.OnStart(start)
	android.Enter()
}
