package main

import (
	"log"
	"strconv"

	"github.com/seletskiy/go-android-rpc/android"
	"github.com/seletskiy/go-android-rpc/android/sdk"
)

type HagridAction struct {
	Relations map[int][]int
	Buttons   map[int]sdk.Switch
	State     map[int]bool
}

func (action HagridAction) GetLayoutName() string {
	return "hagrid_layout"
}

func (action HagridAction) GetButtonTitle() string {
	return "hagrid layout testing"
}

func (action HagridAction) Run() {
	// buttons:
	//     1 2 3
	//     4 5 6
	//     7 8 9
	action.Relations = make(map[int][]int)
	action.Buttons = make(map[int]sdk.Switch)
	action.State = make(map[int]bool)

	action.State = map[int]bool{
		1: false, 2: false, 3: false,
		4: false, 5: false, 6: false,
		7: false, 8: false, 9: false,
	}

	log.Printf("%#v", action)

	action.Relations = map[int][]int{
		1: []int{
			2, 3, 4, 7,
		},
		2: []int{
			1, 3, 5, 6,
		},
		3: []int{
			1, 2, 6, 9,
		},
		4: []int{
			1, 7, 5, 6,
		},
		5: []int{
			2, 8, 4, 6,
		},
		6: []int{
			4, 5, 3, 9,
		},
		7: []int{
			1, 4, 8, 9,
		},
		8: []int{
			2, 5, 7, 9,
		},
		9: []int{
			3, 6, 7, 8,
		},
	}

	for i := 1; i <= 9; i++ {
		button := android.GetViewById(
			"toggle_button_" + strconv.Itoa(i)).(sdk.Switch)

		android.OnClick(button, HagridOnClickHandler{
			i,
			&action,
		})

		action.Buttons[i] = button
	}

	log.Printf("%#v", action.Relations)
}

func (action HagridAction) HangridOnClick(index int) {
	action.State[index] = !action.State[index]
	value := action.State[index]

	for _, otherIndex := range action.Relations[index] {
		action.Buttons[otherIndex].SetChecked(value)
		action.State[otherIndex] = value
	}

	log.Printf("%#v", action.State)
}

type HagridOnClickHandler struct {
	index  int
	action *HagridAction
}

func (handler HagridOnClickHandler) OnClick() {
	handler.action.HangridOnClick(handler.index)
}
