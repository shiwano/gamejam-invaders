package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type fleet struct {
	ships           []*ship
	rect            *sdl.Rect
	horizontalSpeed int32
}

func newFleet(positionY int32) *fleet {
	var ships []*ship
	for i := 0; i < 10; i++ {
		positionX := int32(i * 60)
		s := &ship{
			rect:         &sdl.Rect{X: positionX, Y: positionY, W: 50, H: 50},
			shotVelocity: &sdl.Point{X: 0, Y: 10},
		}
		ships = append(ships, s)
	}
	return &fleet{
		ships:           ships,
		horizontalSpeed: 1,
	}
}

func (f *fleet) IsDestroyed() bool {
	for _, s := range f.ships {
		if !s.isDestroyed {
			return false
		}
	}
	return true
}

func (f *fleet) Rect() *sdl.Rect {
	return f.rect
}

func (f *fleet) Update() {
	for _, s := range f.ships {
		s.move(&sdl.Point{X: s.rect.X + f.horizontalSpeed, Y: s.rect.Y})
	}
}

func (f *fleet) fire() *shot {
	return nil
}
