package components

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type ComponentName string
type Component interface {
	WhenCreated()
	WhenDeleted()
}

// Vector components
type Vec2 struct {
	X, Y float32
}

func (v *Vec2) WhenCreated() {
	v.X = 0
	v.Y = 0
}
func (v *Vec2) WhenDeleted() {}

type Vec3 struct {
	X, Y, Z float32
}

func (v *Vec3) WhenCreated() {
	v.X = 0
	v.Y = 0
	v.Z = 0
}
func (v *Vec3) WhenDeleted() {}

// Rendering-related components
type Drawable struct {
	Renderer      *sdl.Renderer
	TexturePath   string
	Texture       *sdl.Texture
	ScalingFactor float32
}

func (d *Drawable) WhenCreated() {
	img, err := sdl.LoadBMP(d.TexturePath)
	if err != nil {
		panic(fmt.Errorf("loading image %s: %v", d.TexturePath, err))
	}
	defer img.Free()

	d.Texture, err = d.Renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("creating texture from image %s: %v", d.TexturePath, err))
	}
}
func (d *Drawable) WhenDeleted() {
	d.Texture.Destroy()
}
