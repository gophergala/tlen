package main

type Action interface {
	GetHeader() string
	GetDescription() string
	Init()
	Run()
}
