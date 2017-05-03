package main

import (
	"fmt"
	"runtime"

	sf "github.com/zyedidia/sfml/v2.3/sfml"
)

const (
	screenWidth  = 1200
	screenHeight = 800
)

func main() {
	runtime.LockOSThread()

	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Rectangle", sf.StyleDefault, nil)
	window.SetVerticalSyncEnabled(true)
	window.SetFramerateLimit(60)

	window.Clear(sf.ColorBlack)

	fmt.Print("Hi")

}
