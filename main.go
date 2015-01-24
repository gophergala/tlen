package main

import (
	"fmt"

	"github.com/seletskiy/go-android-rpc/android"

	// required for linking
	_ "github.com/seletskiy/go-android-rpc/android/modules"
	"github.com/seletskiy/go-android-rpc/android/sdk"
)

type Scene struct {
	Plot []string
	Next *Scene
}

var SceneOne = &Scene{
	Plot: []string{
		"welcome to tlen",
		"enjoy the ride",
	},
	Next: SceneTwo,
}

var SceneTwo = &Scene{
	Plot: []string{
		"30 years ago...",
	},
}

func (scene *Scene) Run() {
	for {
		text := scene.GetText()

		if text == "" {
			break
		}

		fmt.Println(text)
	}

	if scene.Next != nil {
		scene.Next.Run()
	}
}

func (scene *Scene) GetText() string {
	if len(scene.Plot) == 0 {
		return ""
	}
	text := scene.Plot[0]
	scene.Plot = scene.Plot[1:]
	return text
}

type NextTextButtonHandler struct {
	button   sdk.Button
	textView sdk.TextView
}

func (handler NextTextButtonHandler) OnClick() {
	handler.button.PerformHapticFeedback(0)
	handler.textView.SetText1s(SceneOne.GetText())
}

func start() {
	textView := android.GetViewById("main_layout", "plot_text").(sdk.TextView)
	nextTextButton := android.GetViewById("main_layout", "next_text").(sdk.Button)
	android.OnClick(nextTextButton, NextTextButtonHandler{nextTextButton, textView})
}

func main() {
	android.OnStart(start)

	android.Enter()
}
