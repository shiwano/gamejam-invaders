package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type fleet struct {
	ships           []*ship
	horizontalSpeed int32
}

func newFleet(positionY int32) *fleet {
	var ships []*ship
	for i := 0; i < 10; i++ {
		positionX := int32(i * 50)
		s := &ship{
			rect:         sdl.Rect{X: positionX, Y: positionY, W: 30, H: 30},
			shotVelocity: sdl.Point{X: 0, Y: 10},
		}
		ships = append(ships, s)
	}
	return &fleet{
		ships:           ships,
		horizontalSpeed: 1,
	}
}

func (f *fleet) Type() gameObjectType {
	return gameObjectTypeFleet
}

func (f *fleet) IsDestroyed() bool {
	for _, s := range f.ships {
		if !s.isDestroyed {
			return false
		}
	}
	return true
}

func (f *fleet) Rects() []sdl.Rect {
	var rects []sdl.Rect
	for _, s := range f.ships {
		if !s.isDestroyed {
			rects = append(rects, s.rect)
		}
	}
	return rects
}

func (f *fleet) Update() {
	var fleetRect sdl.Rect
	for _, s := range f.ships {
		s.move(&sdl.Point{X: s.rect.X + f.horizontalSpeed, Y: s.rect.Y})
		fleetRect = fleetRect.Union(&s.rect)
	}

	if (f.horizontalSpeed > 0 && fleetRect.X+fleetRect.W >= windowWidth) ||
		(f.horizontalSpeed < 0 && fleetRect.X <= 0) {
		f.horizontalSpeed = -f.horizontalSpeed

		for _, s := range f.ships {
			s.move(&sdl.Point{X: s.rect.X, Y: s.rect.Y + 30})
		}
	}
}

func (f *fleet) Intersects(g gameObject) {
	if g.Type() != gameObjectTypeShot {
		return
	}

	for _, r := range g.Rects() {
		for _, s := range f.ships {
			if !s.isDestroyed {
				if _, collided := s.rect.Intersect(&r); collided {
					s.isDestroyed = true
				}
			}
		}
	}
}

func (f *fleet) fire() *shot {
	return nil
}
