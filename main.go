package main

import (
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"math/rand"
	"tetris/tetris"
	"time"
)

const (
	winWidth  = 390
	winHeight = 600
	blockSize = 15
	speed     = 15
)

func main() {
	pixelgl.Run(run)
}

func run() {
	rand.Seed(time.Now().UnixNano())

	cfg := pixelgl.WindowConfig{
		Title:  "Tetris",
		Bounds: pixel.R(0, 0, winWidth, winHeight),
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	game := tetris.NewTetris(win, winWidth, winHeight)
	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()
		win.Clear(colornames.Skyblue)

		game.Archive.Draw()
		if game.Current.Empty() {
			game.Current.Make()
			game.Current.Draw()
			win.Update()
			time.Sleep(time.Millisecond * 16)
			continue
		}
		game.Current.Down(0)

		// Handle user input
		if win.Pressed(pixelgl.KeyLeft) {
			// Move tetromino left
			win.Clear(colornames.Skyblue)

			game.Archive.Draw()
			game.Current.Left()
		}
		if win.Pressed(pixelgl.KeyRight) {
			// Move tetromino right
			win.Clear(colornames.Skyblue)

			game.Archive.Draw()
			game.Current.Right()
		}
		if win.Pressed(pixelgl.KeyDown) {
			// Move tetromino down
			win.Clear(colornames.Skyblue)

			game.Archive.Draw()
			game.Current.Down(int(dt))
		}

		if win.Pressed(pixelgl.KeyUp) {
			// Rotate tetromino
			win.Clear(colornames.Skyblue)

			game.Archive.Draw()
			game.Current.Rotate()
			game.Current.Draw()
		}

		ele := game.Archive.CanEliminates()
		if len(ele) > 0 {
			game.Archive.Eliminates(ele)
			game.Current.Draw()
		}

		win.Update()
		time.Sleep(time.Millisecond * 50)
	}
}
