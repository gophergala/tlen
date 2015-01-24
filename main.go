package main

import "fmt"

type Scene struct {
	Plot []string
	Next *Scene
}

func (scene *Scene) Run() {
	for {
		text := scene.GetText()

		if text == "" {
			break
		}

		fmt.Println(text)
	}

	if scene.Next != nil {
		scene.Next.Run()
	}
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
	SceneOne := &Scene{
		Plot: []string{
			"welcome to tlen",
			"enjoy the ride",
		},
	}

	SceneTwo := &Scene{
		Plot: []string{
			"30 years ago...",
		},
	}

	SceneOne.Next = SceneTwo

	SceneOne.Run()

	fmt.Println("\ngame over")
}
