package main

type Location interface {
	GetHeader() string
	GetDescription() string
	LinkLocation(Location)
	LinkAction(Action)
	UnlinkAction(Action)
	GetLinkedLocations() []Location
	GetLinkedActions() []Action
}

type BaseLocation struct {
	LinkedLocations []Location
	LinkedActions   []Action
}

func (location *BaseLocation) LinkLocation(nextLocation Location) {
	location.LinkedLocations = append(location.LinkedLocations, nextLocation)
}

func (location *BaseLocation) GetLinkedLocations() []Location {
	return location.LinkedLocations
}

func (location *BaseLocation) LinkAction(nextAction Action) {
	location.LinkedActions = append(location.LinkedActions, nextAction)
}

func (location *BaseLocation) GetLinkedActions() []Action {
	return location.LinkedActions
}

func (location *BaseLocation) UnlinkAction(action Action) {
	newLinkedActions := []Action{}

	for _, linkedAction := range location.LinkedActions {
		if action.GetButtonTitle() == linkedAction.GetButtonTitle() &&
			action.GetLayoutName() == linkedAction.GetLayoutName() {
			continue
		}
		newLinkedActions = append(newLinkedActions, linkedAction)
	}

	location.LinkedActions = newLinkedActions
}
