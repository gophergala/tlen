package main

import "log"

type BoxesAction struct {
}

func (action BoxesAction) GetButtonTitle() string {
	return "Game of Boxes"
}

func (action BoxesAction) Run() {
	log.Printf("%#v", "123")
}

func (action BoxesAction) GetLayoutName() string {
	return "boxes_layout"
}
