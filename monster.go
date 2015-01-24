package main

import (
	"strconv"

	"github.com/seletskiy/go-android-rpc/android"
	"github.com/seletskiy/go-android-rpc/android/sdk"
)

var scenario = SubAction{
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
	scenario.Draw()
}
