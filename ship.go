package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type ship struct {
	rect *sdl.Rect
}

func (s *ship) move(position *sdl.Point) {
	s.rect.X = position.X
	s.rect.Y = position.Y
}

func (s *ship) fire() {
	fmt.Println("Fire!!!!!!!!!!")
}
