package player

import (
	"fmt"
	"github.com/WaitWut862/cli-dungeon-crawler/internal/common"
)

type Player struct {
	Inventory Inventory
	Health    int
	Position  common.Position
	Statuses  Statuses
	Facing    Direction
}

type Direction int

const (
	North = Direction(iota)
	East
	South
	West
)

type Statuses struct {
	Poisoned  bool
	Boosted   bool
	Weakened  bool
	Enraged   bool
	Fortified bool
}

type Inventory struct {
}

func (p *Player) TurnLeft() {
	p.Facing = p.Facing - 1
	if p.Facing < 0 {
		p.Facing = 3
	}
	fmt.Println(p.Facing)
}

func (p *Player) TurnRight() {
	p.Facing = p.Facing + 1
	if p.Facing > 3 {
		p.Facing = 0
	}
	fmt.Println(p.Facing)
}

func (p *Player) Move() {
	switch p.Facing {
	case North:
		p.Position.Y++
	case East:
		p.Position.X++
	case South:
		p.Position.Y--
	case West:
		p.Position.X--
	}
}

func (p *Player) FacingString() string {
	f := p.Facing
	switch f {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return "Direction not resolved"
	}
}
