package mesh

import (
	"fmt"
	"github.com/gopxl/pixel/pixelgl"
	"log"
	"math/rand"
	"tetris/polygon"
)

func NewMesh(win *pixelgl.Window, blockSize int) *Mesh {
	winWidth, winHeight := int(win.Bounds().Max.X), int(win.Bounds().Max.Y)
	positions := make([][]*Block, winWidth/blockSize)
	for i := 0; i < winWidth/blockSize; i++ {
		positions[i] = make([]*Block, winHeight/blockSize)
	}
	return &Mesh{
		positions: positions,
		blockSize: blockSize,
		win:       win,
	}
}

type Mesh struct {
	positions [][]*Block // 网格坐标
	blockSize int
	win       *pixelgl.Window
}

func (p *Mesh) IsDownCollision(t *Tetromino) bool {
	for i := 0; i < 3; i++ {
		for j := 2; j >= 0; j-- {
			if t.Tetromino[i][j] != nil {
				px, py := t.Tetromino[i][j].X, t.Tetromino[i][j].Y
				log.Printf("check down collision at x: %d, y: %d\n", px, py)
				if py == 0 || p.positions[px][py-1] != nil {
					return true
				}
			}
		}
	}
	return false
}

func (p *Mesh) IsLeftCollision(t *Tetromino) bool {
	for j := 0; j < 3; j++ {
		for i := 2; i >= 0; i-- {
			if t.Tetromino[i][j] != nil {
				px, py := t.Tetromino[i][j].X, t.Tetromino[i][j].Y
				if px == 0 || p.positions[px-1][py] != nil {
					return true
				}
			}
		}
	}

	return false
}

func (p *Mesh) IsRightCollision(t *Tetromino) bool {
	for i := 2; i >= 0; i-- {
		for j := 2; j >= 0; j-- {
			if t.Tetromino[i][j] != nil {
				px, py := t.Tetromino[i][j].X, t.Tetromino[i][j].Y
				if px == len(p.positions)-1 || p.positions[px+1][py] != nil {
					return true
				}
			}
		}
	}
	return false
}

func (p *Mesh) AddTetromino(t *Tetromino) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			p.AddBlock(t.Tetromino[i][j])
		}
	}
}

func (p *Mesh) RemoveTetromino(t *Tetromino) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			p.RemoveBlock(t.Tetromino[i][j])
		}
	}
}

func (p *Mesh) AddBlock(b *Block) {
	if b == nil {
		return
	}
	fmt.Printf("x %d, y %d\n", b.X, b.Y)
	p.positions[b.X][b.Y] = b
}

func (p *Mesh) RemoveBlock(b *Block) {
	p.positions[b.X][b.Y] = nil
}

func (p *Mesh) IsEmptyRow(x int) bool {
	for i := 0; i < 3; i++ {
		if p.positions[x][i] != nil {
			return false
		}
	}
	return true
}

func (p *Mesh) IsFullRow(x int) bool {
	for i := 0; i < 3; i++ {
		if p.positions[x][i] == nil {
			return false
		}
	}
	return true
}

func (p *Mesh) CanEliminates() (res []int) {
	for i := 0; i < len(p.positions); i++ {
		if p.IsFullRow(i) {
			res = append(res, i)
		}
	}
	return
}

func (p *Mesh) Eliminates(x []int) {
	if len(x) == 0 {
		return
	}
	for _, i := range x {
		for j := 0; j < len(p.positions[0]); j++ {
			p.positions[i][j] = nil
		}
	}

	start := x[len(x)-1] + 1
	if start >= len(p.positions) {
		return
	}
	for i := start; i < len(p.positions); i++ {
		for j := 0; j < len(p.positions[i]); j++ {
			p.positions[i][j].MinY -= float64(p.blockSize)
			p.positions[i][j].MaxY -= float64(p.blockSize)
			p.positions[i-start][j] = p.positions[i][j]
			p.positions[i][j] = nil
		}
	}
}

func (p *Mesh) Positions() [][]*Block {
	return p.positions
}

func (p *Mesh) PositionString() string {
	res := fmt.Sprintf("%d*%d\n", len(p.positions), len(p.positions[0]))
	for j := len(p.positions[0]) - 1; j >= 0; j-- {
		for i := 0; i < len(p.positions); i++ {
			if p.positions[i][j] != nil {
				res += "Y "
			} else {
				res += "N "
			}
		}
		res += "\n"
	}
	return res
}

func (p *Mesh) Draw() {
	for i := 0; i < len(p.positions); i++ {
		for j := 0; j < len(p.positions[i]); j++ {
			if p.positions[i][j] != nil {
				p.positions[i][j].Draw()
			}
		}
	}
}

func (p *Mesh) GenTetromino() *Tetromino {
	x, y := p.randPosition()
	log.Printf("generate tetromino, x: %d, y: %d\n", x, y)
	res := NewTetromino(
		p.win,
		&polygon.ShapeBlocks[rand.Intn(len(polygon.ShapeBlocks))],
		x, y, p.blockSize,
		polygon.Colors[rand.Intn(len(polygon.Colors))],
	)
	fmt.Println(res.String())
	return res
}

func (p *Mesh) randPosition() (int, int) {
	return rand.Intn(len(p.positions) - 1), len(p.positions[0]) - 1
}
