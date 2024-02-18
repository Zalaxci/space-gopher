package components

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type ComponentName string
type Component interface {
	whenCreated(r *sdl.Renderer) error
	whenDeleted() error
}

// Physics/position-related components
type vec2 struct {
	X, Y float32
}

func (v *vec2) whenCreated(_ *sdl.Renderer) error {
	v.X = 0
	v.Y = 0
	return nil
}
func (v *vec2) whenDeleted() error {
	return nil
}

type Pos vec2
type Velocity vec2
type Accelaration vec2

// Rendering-related components
type Drawable struct {
	TexturePath   string
	ScalingFactor float32
	texture       *sdl.Texture
}

func (d *Drawable) whenCreated(r *sdl.Renderer) error {
	img, err := sdl.LoadBMP(d.TexturePath)
	if err != nil {
		panic(fmt.Errorf("loading image %s: %v", d.TexturePath, err))
	}
	defer img.Free()

	d.texture, err = r.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("creating texture from image %s: %v", d.TexturePath, err))
	}

	return nil
}
func (d *Drawable) whenDeleted() error {
	d.texture.Destroy()
	return nil
}
