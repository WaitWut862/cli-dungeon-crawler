package mobs

import c "cli-dungeon-crawler/internal/common"

type Mob struct {
	Inventory c.Inventory
	Health    int
	Position  c.Position
	Statuses  Statuses
	Facing    c.Direction
}

type Statuses struct {
	Poisoned  bool
	Boosted   bool
	Weakened  bool
	Enraged   bool
	Fortified bool
}

func (m *Mob) TurnLeft() {
	switch m.Facing {
	case c.North:
		m.Facing = c.West
	case c.East:
		m.Facing = c.North
	case c.South:
		m.Facing = c.East
	case c.West:
		m.Facing = c.South
	}
}

func (m *Mob) TurnRight() {
	switch m.Facing {
	case c.North:
		m.Facing = c.East
	case c.East:
		m.Facing = c.South
	case c.South:
		m.Facing = c.West
	case c.West:
		m.Facing = c.North
	}
}

func (m *Mob) Move() {
	switch m.Facing {
	case c.North:
		m.Position.Y++
	case c.East:
		m.Position.X++
	case c.South:
		m.Position.Y--
	case c.West:
		m.Position.X--
	}
}

func (m *Mob) FacingString() string {
	f := m.Facing
	switch f {
	case c.North:
		return "North"
	case c.East:
		return "East"
	case c.South:
		return "South"
	case c.West:
		return "West"
	default:
		return "Direction not resolved"
	}
}

type Player Mob

func NewPlayer() *Player {
	return &Player{
		0,
		100,
		c.Origin(),
		_,
		c.North,
	}
}
