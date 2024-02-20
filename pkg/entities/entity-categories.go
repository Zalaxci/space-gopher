package entities

import (
	"github.com/Zalaxci/space-gopher/pkg/components"
	"github.com/veandco/go-sdl2/sdl"
)

func CreateEnemies(
	renderer *sdl.Renderer,
) *EntityCategory {
	return createEntityCategory(
		"enemy",
		map[components.ComponentName]components.Component{
			"Size": &components.Vec2{},
			"Pos":  &components.Vec2{},
			"Drawable": &components.Drawable{
				Renderer:      renderer,
				TexturePath:   "assets/gopher.bmp",
				ScalingFactor: float32(3) / float32(8),
			},
		},
		1,
		func(_ components.ComponentName, _ uint16, _ *components.Component) {},
	)
}
func CreatePlayer(
	renderer *sdl.Renderer,
) *EntityCategory {
	return createEntityCategory(
		"player",
		map[components.ComponentName]components.Component{
			"Size":         &components.Vec2{},
			"Pos":          &components.Vec2{},
			"Velocity":     &components.Vec2{},
			"Accelaration": &components.Vec2{},
			"Drawable": &components.Drawable{
				Renderer:      renderer,
				TexturePath:   "assets/gopher.bmp",
				ScalingFactor: float32(3) / float32(8),
			},
		},
		1,
		func(_ components.ComponentName, _ uint16, _ *components.Component) {},
	)
}
