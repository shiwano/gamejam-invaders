package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type gameObject interface {
	IsDestroyed() bool
	Rect() *sdl.Rect
	Update()
}
