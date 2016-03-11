package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	eventChan := make(chan sdl.Event)
	go func() {
		for {
			if event := sdl.PollEvent(); event != nil {
				eventChan <- event
			}
		}
	}()

loop:
	for {
		select {
		case event := <-eventChan:
			switch t := event.(type) {
			case *sdl.QuitEvent:
				break loop
			case *sdl.MouseMotionEvent:
				fmt.Printf("[%d ms] MouseMotion\ttype:%d\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y, t.XRel, t.YRel)
			case *sdl.MouseButtonEvent:
				fmt.Printf("[%d ms] MouseButton\ttype:%d\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n",
					t.Timestamp, t.Type, t.Which, t.X, t.Y, t.Button, t.State)
			}
		}
	}

	sdl.Delay(1000)
	sdl.Quit()
	os.Exit(0)
}
