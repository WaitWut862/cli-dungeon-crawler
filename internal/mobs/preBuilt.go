package mobs

import (
	"math/rand/v2"

	c "cli-dungeon-crawler/internal/common"
)

var MobList = map[string]Mob{
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

var Wanderer Behavior = func(m *Mob) {
	switch rand.IntN(3) {
	case 0:
		m.Move()
	case 1:
		m.TurnLeft()
	case 2:
		m.TurnRight()
	}
}
