package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type shot struct {
	isDestroyed bool
	rect        *sdl.Rect
	velocity    *sdl.Point
}

func (s *shot) IsDestroyed() bool {
	return s.isDestroyed
}

func (s *shot) Rect() *sdl.Rect {
	return s.rect
}

func (s *shot) Update() {
	s.rect.X += s.velocity.X
	s.rect.Y += s.velocity.Y
}

func (s *shot) move(position *sdl.Point) {
	s.rect.X = position.X
	s.rect.Y = position.Y
}
