package mob

import (
	"fmt"

	"github.com/WaitWut862/cli-dungeon-crawler/internal/common"
)

type Mob struct {
	Name        string
	Description string
	Inventory   *common.Inventory
	Health      int
	Position    common.Position
	Statuses    common.Statuses
	Facing      common.Direction
}

type MobParams struct {
	Name          string
	Description   string
	Health        int
	InventorySize int
}

func NewMob(params *MobParams) *Mob {
	if params == nil {
		return &Mob{
			Inventory: common.MakeInventory(0),
			Facing:    common.North,
		}
	}
	return &Mob{
		Name:        params.Name,
		Description: params.Description,
		Health:      params.Health,
		Inventory:   common.MakeInventory(params.InventorySize),
		Facing:      common.North,
	}
}

func (m *Mob) TurnLeft() {
	m.Facing = m.Facing - 1
	if m.Facing < 0 {
		m.Facing = 3
	}
	fmt.Println(m.Facing)
}

func (m *Mob) TurnRight() {
	m.Facing = m.Facing + 1
	if m.Facing > 3 {
		m.Facing = 0
	}
	fmt.Println(m.Facing)
}

func (m *Mob) Move() {
	switch m.Facing {
	case common.North:
		m.Position.Y++
	case common.East:
		m.Position.X++
	case common.South:
		m.Position.Y--
	case common.West:
		m.Position.X--
	}
}

func (m *Mob) FacingString() string {
	switch m.Facing {
	case common.North:
		return "North"
	case common.East:
		return "East"
	case common.South:
		return "South"
	case common.West:
		return "West"
	default:
		return "Direction not resolved"
	}
}

type Player struct {
	*Mob
}

func InitPlayer() *Player {
	return &Player{
		Mob: NewMob(&MobParams{
			Name:          "Player",
			Description:   "The hero of the dungeon",
			Health:        100,
			InventorySize: 10,
		}),
	}
}
