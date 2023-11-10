package mesh

import (
	"fmt"
	"github.com/gopxl/pixel/pixelgl"
	"image/color"
	"tetris/polygon"
)

func NewBlock(win *pixelgl.Window, size int, color color.Color) *Block {
	return &Block{
		Color: color,
		win:   win,
		Block: polygon.NewBlock(win, color),
		size:  size,
	}
}

type Block struct {
	X     int // 网格坐标
	Y     int
	Color color.Color
	win   *pixelgl.Window
	size  int
	*polygon.Block
}

func (b *Block) MoveX(n int) {
	b.SetX(b.X + n)
}

func (b *Block) MoveY(n int) {
	b.SetY(b.Y + n)
}

func (b *Block) SetX(x int) {
	b.X = x
	b.Block.MinX = float64(x) * float64(b.size)
	b.Block.MaxX = b.Block.MinX + float64(b.size)
}

func (b *Block) SetY(y int) {
	b.Y = y
	b.Block.MinY = float64(y) * float64(b.size)
	b.Block.MaxY = b.Block.MinY + float64(b.size)
}

func (b *Block) Set(x, y int) {
	b.SetX(x)
	b.SetY(y)
}

func (b *Block) Draw() {
	b.Block.Draw()
}

func NewTetromino(win *pixelgl.Window, shape *polygon.ShapeBlock, x, y int, blockSize int, color color.Color) *Tetromino {
	t := &Tetromino{
		color:     color,
		shape:     shape,
		win:       win,
		blockSize: blockSize,
		x:         x,
		y:         y,
	}
	t.init()
	return t
}

type Tetromino struct {
	shape     *polygon.ShapeBlock
	Tetromino [3][3]*Block
	color     color.Color
	win       *pixelgl.Window
	blockSize int
	x, y      int
}

func (t *Tetromino) init() {
	for i := 2; i >= 0; i-- {
		for j := 0; j < 3; j++ {
			if t.shape.Shape[i][j] {
				block := NewBlock(t.win, t.blockSize, t.color)
				//block.Set(x+i, y+j)
				block.Set(t.x+j, t.y+2-i)
				t.Tetromino[i][j] = block
			}
		}
	}
}

func (t *Tetromino) Draw() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if t.Tetromino[i][j] != nil {
				t.Tetromino[i][j].Draw()
			}
		}
	}
}

func (t *Tetromino) SetX(x int) {
	t.x = x
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if t.Tetromino[i][j] != nil {
				t.Tetromino[i][j].SetX(x)
			}
		}
	}
}

func (t *Tetromino) SetY(y int) {
	t.y = y
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if t.Tetromino[i][j] != nil {
				t.Tetromino[i][j].SetY(y)
			}
		}
	}
}

func (t *Tetromino) MoveX(n int) {
	t.x += n
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if t.Tetromino[i][j] != nil {
				t.Tetromino[i][j].MoveX(n)
			}
		}
	}
}

func (t *Tetromino) MoveY(n int) {
	t.y += n
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if t.Tetromino[i][j] != nil {
				t.Tetromino[i][j].MoveY(n)
			}
		}
	}
}

func (t *Tetromino) String() string {
	res := "\n"
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if t.Tetromino[i][j] != nil {
				res += fmt.Sprintf("Y(%-3d,%-3d)", t.Tetromino[i][j].X, t.Tetromino[i][j].Y)
			} else {
				res += "N(000,000)"
			}
		}
		res += "\n"
	}
	return res
}

func (t *Tetromino) Rotate() {
	//tetromino := [3][3]*Block{}
	//for i := 0; i < 3; i++ {
	//	for j := 0; j < 3; j++ {
	//		tetromino[i][j] = t.Tetromino[2-j][i]
	//	}
	//}
	//
	//t.Tetromino = tetromino
	// 转置
	for i := 0; i < len(t.shape.Shape); i++ {
		for j := i + 1; j < len(t.shape.Shape[i]); j++ {
			t.shape.Shape[i][j], t.shape.Shape[j][i] = t.shape.Shape[j][i], t.shape.Shape[i][j]
		}
	}

	// 行逆序
	for i := 0; i < len(t.shape.Shape); i++ {
		for j := 0; j < len(t.shape.Shape[i])/2; j++ {
			t.shape.Shape[i][j], t.shape.Shape[i][len(t.shape.Shape[i])-1-j] = t.shape.Shape[i][len(t.shape.Shape[i])-1-j], t.shape.Shape[i][j]
		}
	}
	t.Tetromino = [3][3]*Block{}
	t.init()
}
