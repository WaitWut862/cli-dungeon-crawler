package character

import (
	c "cli-dungeon-crawler/internal/common"
	//w "cli-dungeon-crawler/internal/world"
)

type Character struct {
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

type CharacterPeram func(*Character)

type Behavior = func(m *Character)

func New(name string, options ...CharacterPeram) *Character {
	mob, ok := CharacterList[name]
	if !ok { // todo

	}
	return &mob
}

func (m *Character) P() c.Position {
	return m.Position
}

func (m *Character) X() int {
	return m.Position.X
}

func (m *Character) Y() int {
	return m.Position.Y
}

func (m *Character) Front() c.Position {
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

func (m *Character) TurnLeft() {
	m.Facing = (4 + m.Facing - 1) % 4
	// switch m.Facing {
	// case c.North:
	// 	m.Facing = c.West
	// case c.East:
	// 	m.Facing = c.North
	// case c.South:
	// 	m.Facing = c.East
	// case c.West:
	// 	m.Facing = c.South
	// }
}

func (m *Character) TurnRight() {
	m.Facing = (m.Facing + 1) % 4
	// switch m.Facing {
	// case c.North:
	// 	m.Facing = c.East
	// case c.East:
	// 	m.Facing = c.South
	// case c.South:
	// 	m.Facing = c.West
	// case c.West:
	// 	m.Facing = c.North
	// }
}

func (m *Character) Move() {
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

var directions = map[c.Direction]string{
	c.North: "North",
	c.East:  "East",
	c.South: "South",
	c.West:  "West",
}

func (m *Character) FacingString() string {
	if v, ok := directions[m.Facing]; ok {
		return v
	}

	return "Direction not resolved"
}
