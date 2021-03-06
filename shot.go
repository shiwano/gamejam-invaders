package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type shot struct {
	isDestroyed bool
	rect        sdl.Rect
	velocity    sdl.Point
}

func (s *shot) Type() gameObjectType {
	return gameObjectTypeShot
}

func (s *shot) IsDestroyed() bool {
	return s.isDestroyed
}

func (s *shot) Rects() []sdl.Rect {
	return []sdl.Rect{s.rect}
}

func (s *shot) Update() {
	s.rect.X += s.velocity.X
	s.rect.Y += s.velocity.Y
}

func (s *shot) Intersects(g gameObject) {
	for _, r := range g.Rects() {
		if _, collided := s.rect.Intersect(&r); collided {
			s.isDestroyed = true
		}
	}
}

func (s *shot) move(position *sdl.Point) {
	s.rect.X = position.X
	s.rect.Y = position.Y
}
