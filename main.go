package main

import (
	"log"

	"github.com/Zalaxci/space-gopher/pkg/components"
	"github.com/Zalaxci/space-gopher/pkg/entities"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	gameWidth      = 600
	gameHeight     = 800
	airResistConst = 0.01
)

func windowWasClosed(event *sdl.Event) bool {
	for {
		*event = sdl.PollEvent()
		switch (*event).(type) {
		case *sdl.QuitEvent:
			return true
		case nil:
			return false
		}
	}
}
func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatal("Error initializing SDL2: ", err)
	}

	window, err := sdl.CreateWindow(
		"Hello Go!",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		gameWidth,
		gameHeight,
		sdl.WINDOW_OPENGL,
	)
	if err != nil {
		log.Fatal("Error creating game window: ", err)
	}
	defer window.Destroy()
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatal("Error creating game renderer: ", err)
	}
	defer renderer.Destroy()
	renderer.SetDrawColor(25, 5, 35, 255)

	player := entities.CreatePlayer(renderer)
	playerDrawable := (*player.Components["Drawable"][0]).(*components.Drawable)

	var event sdl.Event
	for !windowWasClosed(&event) {
		renderer.Clear()
		renderer.Copy(
			playerDrawable.Texture,
			&sdl.Rect{
				X: 0,
				Y: 0,
				W: 256,
				H: 256,
			},
			&sdl.Rect{
				X: gameWidth / 2.0,
				Y: gameHeight / 2.0,
				W: 96,
				H: 96,
			},
		)
		renderer.Present()
		sdl.Delay(5)
	}
}
