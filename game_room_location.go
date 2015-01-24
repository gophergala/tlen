package main

type GameRoomLocation struct {
	BaseLocation
}

func (location GameRoomLocation) GetHeader() {
	return "Super-kool game room"
}

func (location GameRoomLocation) GetDescription() string {
	return "Various game stuff"
}
