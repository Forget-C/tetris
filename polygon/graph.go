package polygon

import (
	"golang.org/x/image/colornames"
	"image/color"
)

type (
	Shape [3][3]bool
	Graph [3][3]*Block
)

const (
	GraphSize = 3
)

type ShapeBlock struct {
	Shape Shape
	Name  string
}

// 图形定义
var (
	OBlock = ShapeBlock{
		Shape: Shape{
			{false, false, false},
			{true, true, false},
			{true, true, false},
		},
		Name: "O",
	}
	TBlock = ShapeBlock{
		Shape: Shape{
			{false, false, false},
			{true, true, true},
			{false, true, false},
		},
		Name: "T",
	}
	ZBlock = ShapeBlock{
		Shape: Shape{
			{false, false, false},
			{true, true, false},
			{false, true, true},
		},
		Name: "Z",
	}
	LBlock = ShapeBlock{
		Shape: Shape{
			{true, false, false},
			{true, false, false},
			{true, true, true},
		},
		Name: "L",
	}
	ShapeBlocks = []ShapeBlock{OBlock, TBlock, ZBlock, LBlock}
)

// 颜色定义
var Colors = []color.Color{
	colornames.Red,
	colornames.Green,
	colornames.Blue,
	colornames.Yellow,
	colornames.Purple,
}
