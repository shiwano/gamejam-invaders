package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"runtime"
	"time"
)

const (
	windowWidth  = 600
	windowHeight = 800
)

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := gameLoop(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func gameLoop() error {
	sdl.Init(sdl.INIT_EVERYTHING)
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		windowWidth, windowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return err
	}
	defer renderer.Destroy()

	ticker := time.Tick(time.Second / 60)
	myShip := &ship{
		isPlayer:     true,
		rect:         sdl.Rect{X: 100, Y: windowHeight - 50, W: 50, H: 50},
		shotVelocity: sdl.Point{X: 0, Y: -10},
	}

	var gameObjects []gameObject
	for i := 0; i < 4; i++ {
		positionY := int32(i * 60)
		f := newFleet(positionY)
		gameObjects = append(gameObjects, f)
	}

loop:
	for {
		select {
		default:
		case <-ticker:
			renderer.SetDrawColor(0, 0, 0, 255)
			renderer.Clear()
			renderer.SetDrawColor(255, 255, 255, 255)
			renderer.FillRect(&myShip.rect)

			for _, g := range gameObjects {
				g.Update()
				renderer.FillRects(g.Rects())
			}
			renderer.Present()

			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch t := event.(type) {
				case *sdl.QuitEvent:
					break loop
				case *sdl.MouseMotionEvent:
					myShip.move(&sdl.Point{X: t.X - 25, Y: myShip.rect.Y})
				case *sdl.MouseButtonEvent:
					if t.State == 0 {
						shot := myShip.fire()
						gameObjects = append(gameObjects, shot)
					}
				}
			}
		}
	}
	return nil
}
