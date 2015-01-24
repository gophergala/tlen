package main

type KitchenLocation struct {
	BaseLocation
}

func (location KitchenLocation) GetHeader() {
	return "Kitchen"
}

func (location KitchenLocation) GetDescription() string {
	return "Some fragments on the table."
}
