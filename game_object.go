package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type gameObjectType int

const (
	gameObjectTypeFleet gameObjectType = iota + 1
	gameObjectTypeShot
)

type gameObject interface {
	Type() gameObjectType
	IsDestroyed() bool
	Rects() []sdl.Rect
	Update()
}
