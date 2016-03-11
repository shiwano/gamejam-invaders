package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type ship struct {
	isDestroyed bool
	rect        *sdl.Rect
}

func (s *ship) move(position *sdl.Point) {
	s.rect.X = position.X
	s.rect.Y = position.Y
}

func (s *ship) fire() *shot {
	fmt.Println("Fire!!!!!!!!!!")
	return &shot{
		rect:     &sdl.Rect{X: s.rect.X, Y: s.rect.Y, W: 10, H: 10},
		velocity: &sdl.Point{X: 0, Y: 1},
	}
}
