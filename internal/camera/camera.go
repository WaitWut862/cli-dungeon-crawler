package camera

import "github.com/WaitWut862/cli-dungeon-crawler/internal/common"

type Camera struct {
	Position common.Position
	Width    int
	Height   int
}

func NewCamera(width, height int) *Camera {
	return &Camera{
		Position: common.Position{X: 0, Y: 0},
		Width:    width,
		Height:   height,
	}
}

func (c *Camera) MoveNorth() { c.Position.Y++ }
func (c *Camera) MoveSouth() { c.Position.Y-- }
func (c *Camera) MoveEast()  { c.Position.X++ }
func (c *Camera) MoveWest()  { c.Position.X-- }
