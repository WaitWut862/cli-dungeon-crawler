package world

import (
	c "cli-dungeon-crawler/internal/common"
)

type World struct {
	Tick     int
	TileMap  map[c.Position]Tile
	Entities map[c.Position][]Entity
}

type Tile struct {
	Position   c.Position
	GroundType string
}

type Entity struct {
	Position    c.Position
	Name        string
	Obstructing bool
}

func (w *World) MakeWorld() {
	w.TileMap = make(map[c.Position]Tile)
	w.Entities = make(map[c.Position][]Entity)
}

func (w *World) UpdateTick() {
	w.Tick = w.Tick + 1
}
