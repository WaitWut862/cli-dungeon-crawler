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

func New() *World {
	return &World{
		TileMap:  make(map[c.Position]Tile),
		Entities: make(map[c.Position][]Entity),
	}
}

func (w *World) UpdateTick() {
	w.Tick = w.Tick + 1
}

func (w *World) GenerateChunk(pt1, pt2 c.Position, ground string) {
	minX := min(pt1.X, pt2.X)
	maxX := max(pt1.X, pt2.X)
	minY := min(pt1.Y, pt2.Y)
	maxY := max(pt1.Y, pt2.Y)

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {

		}
	}
}
