package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"time"
)

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
		800, 600, sdl.WINDOW_SHOWN)
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

	eventChan := make(chan sdl.Event, 1024)
	go func() {
		for {
			if event := sdl.PollEvent(); event != nil {
				eventChan <- event
			}
		}
	}()

	ticker := time.Tick(time.Second / 60)
	myShip := &ship{
		rect: &sdl.Rect{X: 100, Y: 100, W: 100, H: 100},
	}

loop:
	for {
		select {
		default:
		case <-ticker:
			renderer.SetDrawColor(0, 0, 0, 255)
			renderer.Clear()
			renderer.SetDrawColor(255, 255, 255, 255)
			renderer.FillRect(myShip.rect)
			renderer.Present()
		case event := <-eventChan:
			switch t := event.(type) {
			case *sdl.QuitEvent:
				break loop
			case *sdl.MouseMotionEvent:
				myShip.move(&sdl.Point{X: t.X, Y: myShip.rect.Y})
			case *sdl.MouseButtonEvent:
				fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
			}
		}
	}
	return nil
}
