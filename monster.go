package main

import (
	"log"
	"strconv"

	"github.com/seletskiy/go-android-rpc/android"
	"github.com/seletskiy/go-android-rpc/android/sdk"
)

type SubActions []*SubAction

type SubAction struct {
	Title       string
	Description string
	PreDraw     func(subaction *SubAction) bool
	Answers     SubActions
}

var uids []string

var subactions = SubAction{
	Title:       "",
	Description: "Danger! First step. What do you feel?",
	Answers: SubActions{
		&SubAction{
			Title:       "1",
			Description: "1111",
			Answers: SubActions{
				&SubAction{
					Title:       "1_1",
					Description: "1_11111",
					PreDraw: func(subaction *SubAction) bool {
						sss := android.CreateView("888", "android.widget.Button").(sdk.Button)
						sss.SetText1s("hello from dream")
						android.AttachView(sss.View, strconv.Itoa(MainLayoutId))
						return false
					},
				},
				&SubAction{
					Title:       "1_2",
					Description: "1_2222",
					PreDraw: func(subaction *SubAction) bool {
						sss := android.CreateView("999", "android.widget.Button").(sdk.Button)
						sss.SetText1s("yet button")
						android.AttachView(sss.View, strconv.Itoa(MainLayoutId))
						return true
					},
				},
			},
		},
		&SubAction{
			Title:       "2",
			Description: "222222",
			Answers: SubActions{
				&SubAction{
					Title: "Game over pls",
					PreDraw: func(subaction *SubAction) bool {
						sss := android.CreateView("888", "android.widget.Button").(sdk.Button)
						sss.SetText1s("Game over :((((((((((((")
						android.AttachView(sss.View, strconv.Itoa(MainLayoutId))
						return false
					},
				},
			},
		},
	},
}

func InitDreamWithMonster(action *Action, button sdk.Button) {
	subactions.Draw()
}

type SubActionButtonHandler struct {
	subaction *SubAction
	button    *sdk.Button
}

func (handler SubActionButtonHandler) OnClick() {
	if handler.subaction.PreDraw == nil || handler.subaction.PreDraw(handler.subaction) {
		handler.subaction.Draw()
	}
}

func (subactions *SubActions) Draw() {
	for _, subaction := range *subactions {
		button := createView("android.widget.Button").(sdk.Button)
		button.SetText1s(subaction.Title)
		android.OnClick(button, SubActionButtonHandler{
			subaction,
			&button,
		})
		attachView(button.View)
	}
}

func (subaction *SubAction) Draw() {
	hideAll()

	desc := createView("android.widget.TextView").(sdk.TextView)
	desc.SetText1s(subaction.Description)
	attachView(desc.View)

	subaction.Answers.Draw()
}

func hideAll() {
	for _, uid := range uids {
		view := android.GetViewById("main_layout", uid)
		if textView, ok := view.(sdk.TextView); ok {
			log.Printf("%#v\n", textView)
			textView.SetVisibility(ViewGone)
		} else if button, ok := view.(sdk.Button); ok {
			log.Printf("%#v\n", button)
			button.SetVisibility(ViewGone)
		} else {
			log.Printf("%#v\n", "vse govno!")
		}
	}

	uids = []string{}
}

var viewUid = 100

func createView(what string) interface{} {
	viewUid++

	uid := strconv.Itoa(viewUid)
	uids = append(uids, uid)

	return android.CreateView(uid, what)
}

func attachView(view sdk.View) {
	android.AttachView(view, strconv.Itoa(MainLayoutId))
}
