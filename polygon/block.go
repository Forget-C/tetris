package polygon

import (
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/imdraw"
	"github.com/gopxl/pixel/pixelgl"
	"image/color"
)

func NewBlock(win *pixelgl.Window, color color.Color) *Block {
	return &Block{
		color: color,
		win:   win,
	}
}

// Block
// 屏幕中的方块， 坐标单位为像素
type Block struct {
	MinX, MinY, MaxX, MaxY float64
	color                  color.Color
	win                    *pixelgl.Window
}

func (b *Block) Draw() {
	imd := imdraw.New(nil)
	imd.Color = b.color
	rect := pixel.R(b.MinX, b.MinY, b.MaxX, b.MaxY)
	imd.Push(rect.Min, rect.Max)
	imd.Rectangle(0)
	imd.Draw(b.win)
}
