package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type ship struct {
	isDestroyed  bool
	rect         *sdl.Rect
	shotVelocity *sdl.Point
}

func (s *ship) IsDestroyed() bool {
	return s.isDestroyed
}

func (s *ship) Rect() *sdl.Rect {
	return s.rect
}

func (s *ship) Update() {
}

func (s *ship) move(position *sdl.Point) {
	s.rect.X = position.X
	s.rect.Y = position.Y
}

func (s *ship) fire() *shot {
	return &shot{
		rect: &sdl.Rect{
			X: s.rect.X + s.rect.W/2,
			Y: s.rect.Y,
			W: 10,
			H: 10,
		},
		velocity: s.shotVelocity,
	}
}
