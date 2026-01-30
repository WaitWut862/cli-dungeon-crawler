package world

import (
	"github.com/WaitWut862/cli-dungeon-crawler/internal/common"
)

type World struct {
	Tick     int
	TileMap  map[common.Position]Tile
	Entities map[common.Position][]Entity
}

type Tile struct {
	position   common.Position
	groundType string
}

type Entity struct {
	position    common.Position
	name        string
	obstructing bool
}

func (w *World) MakeWorld() {
	w.TileMap = make(map[common.Position]Tile)
	w.Entities = make(map[common.Position][]Entity)
}

func (w *World) UpdateTick() {
	w.Tick = w.Tick + 1
}
