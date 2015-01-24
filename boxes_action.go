package main

import "log"

type BoxesAction struct {
}

func (location BoxesAction) GetHeader() {
	return "Game of Boxes"
}

func (location BoxesAction) GetDescription() string {
	return "Ta-ta ta-ta-ta-ta ta-ta"
}

func (location BoxesAction) Init() {

}

func (location BoxesAction) Run() {
	log.Printf("%#v", "123")
}
