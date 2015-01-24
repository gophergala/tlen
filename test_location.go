package main

type TestLocation struct {
	BaseLocation
	Header string
	Description string
}

func NewTestLocation(header, desc string) *TestLocation {
	return &TestLocation{Header: header, Description: desc}
}

func (location TestLocation) GetHeader() string {
	return location.Header
}

func (location TestLocation) GetDescription() string {
	return location.Description
}
