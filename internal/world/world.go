package world

import (
	c "cli-dungeon-crawler/internal/common"
	m "cli-dungeon-crawler/internal/mobs"
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

func (w *World) GenerateChunk(pt1, pt2 c.Position, ground string) {
	xRange := pt1.X - pt2.X
	yRange := pt1.Y - pt2.Y
	for x := range xRange {
		for y := range yRange {
			w.TileMap[c.Position{X: x, Y: y}] = Tile{GroundType: ground}
		}
	}

}
