package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type gameObject interface {
	IsDestroyed() bool
	Rects() []sdl.Rect
	Update()
}
