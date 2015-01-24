package main

import "fmt"

type Scene struct {
	Plot []string
	Next *Scene
}

func (scene *Scene) GetText() string {
	if len(scene.Plot) == 0 {
		return ""
	}
	text := scene.Plot[0]
	scene.Plot = scene.Plot[1:]
	return text
}

func main() {
	Scene := &Scene{
		Plot: []string{
			"welcome to tlen",
			"enjoy the ride",
		},
	}

	for {
		text := Scene.GetText()

		if text == "" {
			break
		}

		fmt.Println(text)
	}

	fmt.Println("\ngame over")
}
