package player

import (
	"fmt"
)

type Player struct {
	inventory Inventory
	health    int
	position  Position
	statuses  Statuses
	facing    Direction
}

type Direction int

const (
	north = Direction(iota)
	east
	south
	west
)

type Statuses struct {
	poisoned  bool
	boosted   bool
	weakened  bool
	enraged   bool
	fortified bool
}

type Position struct {
	x int
	y int
}

type Inventory struct {
}

func (p *Player) turnLeft() {
	p.facing = p.facing - 1
	if p.facing < 0 {
		p.facing = 3
	}
	fmt.Println(p.facing)
}

func (p *Player) turnRight() {
	p.facing = p.facing + 1
	if p.facing > 3 {
		p.facing = 0
	}
	fmt.Println(p.facing)
}

func (p *Player) move() {
	switch p.facing {
	case north:
		p.position.y++
	case east:
		p.position.x++
	case south:
		p.position.y--
	case west:
		p.position.x--
	}
}

func (p *Player) facingString() string {
	f := p.facing
	switch f {
	case north:
		return "North"
	case east:
		return "East"
	case south:
		return "South"
	case west:
		return "West"
	default:
		return "Direction not resolved"
	}
}
