package main

import (
	"log"
	"strconv"

	"github.com/seletskiy/go-android-rpc/android"
	"github.com/seletskiy/go-android-rpc/android/sdk"
)

type State struct {
	ViewId      int
	Location    Location
	LayoutId    string
	LayoutName  string
	MoveCounter int
}

type Game struct {
	state      *State
	headerView sdk.TextView
	descView   sdk.TextView

	//viewId      int
	viewObjects map[string][]sdk.View

	LocationOnClickHandler GameLocationOnClickHandler
}

func NewGame(state *State) *Game {
	return &Game{
		state: state,
	}
}

func (game *Game) IncrementMoveCounter() {
	game.state.MoveCounter += 1
}

func (game *Game) SetLocation(location Location) {
	game.state.Location = location
}

func (game *Game) SetLayoutName(layoutName string) {
	layoutResponse := android.GetLayoutById(layoutName)
	log.Printf("game.go:44 %#v", layoutResponse)
	log.Printf("game.go:46 %#v", game.state)
	game.state.LayoutId = layoutResponse["layout_id"].(string)
	game.state.LayoutName = layoutName
}

func (game *Game) RestoreMainLayout() {
	game.SetLayoutName("main_layout")
	game.SwitchLayout()
}

func (game *Game) SwitchLayout() {
	android.ChangeLayout(game.state.LayoutName)
}

func (game *Game) Start() {
	game.headerView = android.GetViewById("header_text").(sdk.TextView)
	game.descView = android.GetViewById("desc_text").(sdk.TextView)
	game.viewObjects = make(map[string][]sdk.View)

	game.SwitchLocation()
}

func (game *Game) ClearViews() {
	log.Printf("game.go:70 %#v", game.viewObjects)

	for _, view := range game.viewObjects[game.state.LayoutName] {
		log.Printf("game.go:71 %#v", view)
		android.RemoveView(view, game.state.LayoutId)
	}

	game.viewObjects[game.state.LayoutName] = []sdk.View{}
}

func (game *Game) SwitchLocation() {
	game.ClearViews()

	location := game.state.Location

	game.headerView.SetText1s(location.GetHeader())
	android.SetTextFromHtml(game.descView, location.GetDescription())

	game.SetLocationArterfactsVisibility(true)

	linkedLocations := location.GetLinkedLocations()
	for _, linkedLocation := range linkedLocations {
		log.Printf("game.go:90 %#v", linkedLocation)
		button := game.CreateView("android.widget.Button").(sdk.Button)
		//button.SetText1s(linkedLocation.GetButtonTitle())
		android.SetTextFromHtml(button, linkedLocation.GetButtonTitle())

		android.OnClick(button,
			GameLocationOnClickHandler{
				game, button, linkedLocation,
			},
		)

		game.AttachView(button.View)
	}

	log.Printf("game.go:103 %#v", game.viewObjects)
	location.Enter(game.state)
	log.Printf("game.go:105 %#v", game.viewObjects)
}

func (game *Game) SetLocationArterfactsVisibility(should bool) {
	headerVisible, _ := game.headerView.IsShown()
	descVisible, _ := game.descView.IsShown()

	if should {
		if !headerVisible {
			game.headerView.SetVisibility(ViewVisible)
		}
		if !descVisible {
			game.descView.SetVisibility(ViewVisible)
		}

	} else {
		if headerVisible {
			game.headerView.SetVisibility(ViewGone)
		}

		if descVisible {
			game.descView.SetVisibility(ViewGone)
		}
	}
}

func (game *Game) LocationOnClick(button sdk.Button, location Location) {
	game.IncrementMoveCounter()
	game.SetLocation(location)
	game.SwitchLocation()
}

func (game *Game) CreateView(viewName string) interface{} {
	game.state.ViewId++
	id := strconv.Itoa(game.state.ViewId)
	created := android.CreateView(id, viewName)
	return created
}

func (game *Game) AttachView(view sdk.View) {
	game.viewObjects[game.state.LayoutName] = append(
		game.viewObjects[game.state.LayoutName], view)

	android.AttachView(view, game.state.LayoutId)
}

type GameLocationOnClickHandler struct {
	game     *Game
	button   sdk.Button
	location Location
}

func (handler GameLocationOnClickHandler) OnClick() {
	handler.game.LocationOnClick(handler.button, handler.location)
}
