package tetris

import (
	"github.com/gopxl/pixel/pixelgl"
	"tetris/mesh"
)

const (
	blockSize = 15 // 参照屏幕坐标的大小， 单位像素
	speed     = 1  // 参照网格坐标大小，实际屏幕内移动像素为 blockSize * speed
)

type Tetris struct {
	*Current
	win       *pixelgl.Window
	score     int
	level     int
	paused    bool
	winWidth  float64
	winHeight float64
	Archive   *mesh.Mesh
}

func NewTetris(win *pixelgl.Window, winWidth int, winHeight int) *Tetris {
	archive := mesh.NewMesh(win, blockSize)
	return &Tetris{
		win:       win,
		winWidth:  float64(winWidth),
		winHeight: float64(winHeight),
		Archive:   archive,
		Current: &Current{
			win:     win,
			archive: archive,
			speed:   speed,
		},
	}
}
