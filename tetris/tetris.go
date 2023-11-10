package tetris

import (
	"github.com/gopxl/pixel/pixelgl"
	"tetris/mesh"
	"tetris/polygon"
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

//func (t *Tetris) IsDownCollision() bool {
//	return t.mesh.IsDownCollision(t.current.Graph())
//}
//
//func (t *Tetris) IsLeftCollision() bool {
//	return t.mesh.IsLeftCollision(t.current.Graph())
//} //
//func (t *Tetris) IsRightCollision() bool {
//	return t.mesh.IsRightCollision(t.current.Graph())
//}

//func (t *Tetris) Current() *polygon.Polygon {
//	return t.current
//}
//
//func (t *Tetris) CurrentHitBottom() bool {
//	return t.current.HitBottom() || t.mesh.IsDownCollision(t.current.Graph())
//}

//func (t *Tetris) CurrentDown() {
//	t.current.Down()
//	if t.CurrentHitBottom() {
//		graph := t.current.Graph()
//		for i := 0; i < polygon.GraphSize; i++ {
//			for j := 0; j < polygon.GraphSize; j++ {
//				t.mesh.SetPosition(graph[i][j])
//			}
//		}
//	}
//}

func randPosition(maxX, maxY float64) (float64, float64) {
	//x := rand.Intn((int(maxX)/blockSize)-polygon.GraphSize) * blockSize
	//return float64(x), maxY - float64(blockSize*polygon.GraphSize)
	return maxX - float64(blockSize*polygon.GraphSize), maxY - float64(blockSize*polygon.GraphSize)
}
