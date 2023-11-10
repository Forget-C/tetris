package tetris

import (
	"github.com/gopxl/pixel/pixelgl"
	"log"
	"tetris/mesh"
)

type Current struct {
	*mesh.Tetromino
	win     *pixelgl.Window
	archive *mesh.Mesh
	speed   int // 网格坐标单位
}

func (c *Current) init() {
	if c.Tetromino == nil {
		c.Tetromino = c.archive.GenTetromino()
	}
}

func (c *Current) Down(dt int) {
	if dt != 0 {
		c.Tetromino.MoveY(-c.speed * dt)
	} else {
		c.Tetromino.MoveY(-c.speed)
	}
	c.Tetromino.Draw()
	if c.archive.IsDownCollision(c.Tetromino) {
		c.archive.AddTetromino(c.Tetromino)
		log.Printf("%+v", c.archive.PositionString())
		c.Tetromino = c.archive.GenTetromino()
	}
}

func (c *Current) Left() {
	if c.archive.IsLeftCollision(c.Tetromino) {
		c.Tetromino.Draw()
		return
	}
	c.Tetromino.MoveX(-c.speed)
	c.Tetromino.Draw()
}

func (c *Current) Right() {
	if c.archive.IsRightCollision(c.Tetromino) {
		c.Tetromino.Draw()
		return
	}
	c.Tetromino.MoveX(c.speed)
	c.Tetromino.Draw()
}

func (c *Current) Rotate() {
	c.Tetromino.Rotate()
	c.Tetromino.Draw()
}

func (c *Current) Empty() bool {
	if c.Tetromino == nil {
		return true
	}
	return false
}

func (c *Current) Make() {
	c.init()
}
