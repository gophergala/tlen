package main

import (
	"log"
	"strconv"

	"github.com/seletskiy/go-android-rpc/android"
	"github.com/seletskiy/go-android-rpc/android/sdk"
)

type State struct {
	Location Location
	//Action   *Action
	LayoutId   string
	LayoutName string
}

type Game struct {
	state      State
	headerView sdk.TextView
	descView   sdk.TextView

	viewId      int
	viewObjects []interface{}

	LocationOnClickHandler GameLocationOnClickHandler
	ActionOnClickHandler   GameActionOnClickHandler
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
	android.ChangeLayout(game.state.LayoutName)
}

func (game *Game) Start() {
	game.SetLayoutName("main_layout")

	log.Printf("%#v\n", game)
	log.Printf("%#v\n", game.state)

	game.headerView = android.GetViewById(
		"main_layout", "header_text").(sdk.TextView)

	game.descView = android.GetViewById(
		"main_layout", "desc_text").(sdk.TextView)

	game.SwitchLocation()
}

func (game *Game) ClearViews() {
	for _, view := range game.viewObjects {
		switch view.(type) {
		case sdk.TextView:
			view.(sdk.TextView).SetVisibility(ViewGone)
		case sdk.Button:
			view.(sdk.Button).SetVisibility(ViewGone)
		default:
			log.Printf("%#v\n", "huevi dela")
			log.Printf("%#v\n", view)
		}
	}

	game.viewObjects = make([]interface{}, 0)
}

func (game *Game) SwitchLocation() {
	game.ClearViews()

	location := game.state.Location

	game.headerView.SetText1s(location.GetHeader())
	game.descView.SetText1s(location.GetDescription())

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

	action.Run()
}

func (game *Game) LocationOnClick(button sdk.Button, location Location) {
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
	game.viewObjects = append(game.viewObjects, created)
	return created
}

func (game *Game) AttachView(view sdk.View) {
	android.AttachView(view, game.state.LayoutId)
}
