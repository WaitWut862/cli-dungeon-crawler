package character

import (
	"math/rand/v2"

	c "cli-dungeon-crawler/internal/common"
)

var CharacterList = map[string]Character{
	"walker": {
		Behavior: Wanderer,
		Health:   1,
	},

	"player": {
		Health:   100,
		Position: c.Pos(0, 0),
		Facing:   c.North,
	},
}

var Wanderer Behavior = func(m *Character) {
	switch rand.IntN(3) {
	case 0:
		m.Move()
	case 1:
		m.TurnLeft()
	case 2:
		m.TurnRight()
	}
}
