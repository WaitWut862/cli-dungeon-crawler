package world

import (
	c "cli-dungeon-crawler/internal/common"
)

type World struct {
	tick     int
	tileMap  map[c.Position]Tile
	entities map[c.Position][]Entity
}

type Tile struct {
	position   c.Position
	groundType string
}

type Entity struct {
	position    c.Position
	name        string
	obstructing bool
}

func (w *World) makeWorld() {
	w.tileMap = make(map[c.Position]Tile)
	w.entities = make(map[c.Position][]Entity)
}

func (w *World) updateTick() {
	w.tick = w.tick + 1
}
