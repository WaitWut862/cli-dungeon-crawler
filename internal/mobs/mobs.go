package mobs

import (
	c "cli-dungeon-crawler/internal/common"
	//w "cli-dungeon-crawler/internal/world"
)

type Mob struct {
	Behavior  Behavior
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

type MobPeram func(*Mob)

type Behavior = func(m *Mob)

func NewMob(name string, options ...MobPeram) *Mob {
	mob, ok := MobList[name]
	if !ok { // todo

	}
	return &mob
}

func (m *Mob) P() c.Position {
	return m.Position
}

func (m *Mob) X() int {
	return m.Position.X
}

func (m *Mob) Y() int {
	return m.Position.Y
}

func (m *Mob) Front() c.Position {
	switch m.Facing {
	case c.North:
		return c.Pos(m.X()+1, m.Y())
	case c.East:
		return c.Pos(m.X(), m.Y()+1)
	case c.South:
		return c.Pos(m.X()-1, m.Y())
	case c.West:
		return c.Pos(m.X(), m.Y()-1)
	}
	return c.Pos(0, 0)
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
