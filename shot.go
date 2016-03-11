package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type shot struct {
	isDestroyed bool
	rect        *sdl.Rect
	velocity    *sdl.Point
}

func (s *shot) move(position *sdl.Point) {
	s.rect.X = position.X
	s.rect.Y = position.Y
}

func (s *shot) update() {
	s.rect.X += s.velocity.X
	s.rect.Y += s.velocity.Y
}
