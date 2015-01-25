package main

import (
	"log"
	"strconv"

	"github.com/seletskiy/go-android-rpc/android"
	"github.com/seletskiy/go-android-rpc/android/sdk"
)

type State struct {
	Location     Location
	Action       Action
	LayoutId     string
	LayoutName   string
	MoveCounter  int
	CurrentStage int
	GameProgress int
}

type Game struct {
	state      State
	headerView sdk.TextView
	descView   sdk.TextView

	viewId      int
	viewObjects map[string][]sdk.View

	LocationOnClickHandler GameLocationOnClickHandler
	ActionOnClickHandler   GameActionOnClickHandler
}

func (game *Game) IncrementMoveCounter() {
	game.state.MoveCounter += 1
	log.Printf("!!! current game MOVE is  %#v\n", game.state.MoveCounter)
}

func (game *Game) SetCurrentStage(stage int) {
	log.Printf("!!! current game STAGE is  %#v\n", stage)
	game.state.CurrentStage = stage
}

func (game *Game) SetLocation(location Location) {
	game.state.Location = location
}

func (game *Game) SetLayoutName(layoutName string) {
	layoutResponse := android.GetLayoutById(layoutName)
	game.state.LayoutId = layoutResponse["layout_id"].(string)
	game.state.LayoutName = layoutName
}

func (game *Game) RestoreMainLayout() {
	game.SetLayoutName("main_layout")
	game.SwitchLayout()
}

func (game *Game) SwitchLayout() {
	log.Printf("game.go:46 %#v", game.state.LayoutName)
	android.ChangeLayout(game.state.LayoutName)
}

func (game *Game) Start() {
	game.SetLayoutName("main_layout")

	game.headerView = android.GetViewById("header_text").(sdk.TextView)

	game.descView = android.GetViewById("desc_text").(sdk.TextView)

	game.viewObjects = make(map[string][]sdk.View)

	game.SwitchLocation()
}

func (game *Game) ClearViews() {
	for _, view := range game.viewObjects[game.state.LayoutName] {
		android.RemoveView(view, game.state.LayoutId)
	}

	game.viewObjects[game.state.LayoutName] = []sdk.View{}
}

func (game *Game) SwitchLocation() {
	game.ClearViews()

	location := game.state.Location

	game.headerView.SetText1s(location.GetHeader())
	//game.descView.SetText1s(location.GetDescription())
	android.SetTextFromHtml(game.descView, location.GetDescription())

	game.SetLocationArterfactsVisibility(true)

	linkedLocations := location.GetLinkedLocations()
	for _, linkedLocation := range linkedLocations {
		button := game.CreateView("android.widget.Button").(sdk.Button)
		button.SetText1s(linkedLocation.GetHeader())

		android.OnClick(button, GameLocationOnClickHandler{
			button, linkedLocation})

		game.AttachView(button.View)
	}

	linkedActions := location.GetLinkedActions()
	for _, linkedAction := range linkedActions {
		button := game.CreateView("android.widget.Button").(sdk.Button)
		button.SetText1s(linkedAction.GetButtonTitle())

		android.OnClick(button, GameActionOnClickHandler{button, linkedAction})

		game.AttachView(button.View)
	}
}

func (game *Game) RunAction(action Action) {
	layout := action.GetLayoutName()

	game.SetLayoutName(layout)
	game.SwitchLayout()

	game.SetLocationArterfactsVisibility(false)

	game.state.Action = action
	action.Run()
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

type GameLocationOnClickHandler struct {
	button   sdk.Button
	location Location
}

func (handler GameLocationOnClickHandler) OnClick() {
	game.LocationOnClick(handler.button, handler.location)
}

func (game *Game) ActionOnClick(button sdk.Button, action Action) {
	game.RunAction(action)
}

type GameActionOnClickHandler struct {
	button sdk.Button
	action Action
}

func (handler GameActionOnClickHandler) OnClick() {
	game.ActionOnClick(handler.button, handler.action)
}

func (game *Game) CreateView(viewName string) interface{} {
	game.viewId++
	id := strconv.Itoa(game.viewId)
	created := android.CreateView(id, viewName)
	return created
}

func (game *Game) AttachView(view sdk.View) {
	game.viewObjects[game.state.LayoutName] = append(
		game.viewObjects[game.state.LayoutName], view)

	android.AttachView(view, game.state.LayoutId)
}
