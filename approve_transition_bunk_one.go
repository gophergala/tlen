package main

type ApproveTransitionBunkOneAction struct{}

func (action ApproveTransitionBunkOneAction) GetButtonTitle() string {
	return "Ok, got it"
}

func (action ApproveTransitionBunkOneAction) GetLayoutName() string {
	return "main_layout"
}

func (action ApproveTransitionBunkOneAction) Run() {
	game.SetLocation(locations["bunk2"])
	game.SwitchLocation()
}
