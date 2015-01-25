package main

import (
	"log"

	"github.com/seletskiy/go-android-rpc/android"
	"github.com/seletskiy/go-android-rpc/android/sdk"
)

type CaressCatSubgame struct {
	Subgame
}

func (subgame CaressCatSubgame) GetButtonTitle() string {
	return "Play with cat"
}

func (subgame CaressCatSubgame) GetLayoutName() string {
	return "main_layout"
}

type CaressCatLocation struct {
	BaseLocation

	game *CaressCatSubgame
}

func (location CaressCatLocation) GetDescription() string {
	return "<p>You see " + Cat + ".</p>"
}

func (location CaressCatLocation) GetHeader() string {
	return "<p>Your family bunk.</p>"
}

func (location CaressCatLocation) Enter(state *State) {
	log.Printf("caress_cat_subgame.go:29 %#v", location.game)
	caressButton := location.game.CreateView(
		"android.widget.Button",
	).(sdk.Button)

	android.SetTextFromHtml(caressButton, "Caress "+Cat)

	android.OnClick(caressButton, CaressCatOnClickHandler{
		caressButton,
		state,
	})

	location.game.AttachView(caressButton.View)
}

func (subgame CaressCatSubgame) Enter(state *State) {
	subgame.Subgame.Enter(state)

	defer android.PanicHandler()

	main := &CaressCatLocation{game: &subgame}

	main.Link(globalLocations["home"])

	subgame.Link(globalLocations["home"])
	log.Printf("caress_cat_subgame.go:45 %#v", subgame)
	subgame.SetLocation(main)

	subgame.Start()
}

type CaressCatOnClickHandler struct {
	button sdk.Button
	state  *State
}

func (handler CaressCatOnClickHandler) OnClick() {
	handler.button.PerformHapticFeedback(0)
	handler.state.Progress++
	handler.state.Cat++
	handler.state.MoveCounter++
}
