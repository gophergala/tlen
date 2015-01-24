package main

type Location interface {
	GetHeader() string
	GetDescription() string
	LinkLocation(*Location)
	LinkAction(*Action)
}

type BaseLocation struct {
	LinkedLocations []*Location
	LinkedActions   []*Action
}

func (location GameRoomLocation) LinkLocation(nextLocation *Location) {
	location.LinkedLocations = append(location.LinkedLocations, nextLocation)
}

func (location GameRoomLocation) LinkAction(nextAction *Action) {
	location.LinkedActions = append(location.LinkedActions, nextAction)
}
