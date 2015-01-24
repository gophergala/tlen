package main

import (
	"log"
	"strconv"

	"github.com/seletskiy/go-android-rpc/android"
	"github.com/seletskiy/go-android-rpc/android/sdk"
)

type BaseActions []*SubAction

type SubAction struct {
	Title       string
	Opened      func() interface{}
	Description string
	PreDraw     func(subaction *SubAction) bool
	Answers     SubActions
}

var uids []string

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
		if subaction.Opened != nil && !subaction.Opened().(bool) {
			continue
		}
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
