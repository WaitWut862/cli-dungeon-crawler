package world

import (
	c "cli-dungeon-crawler/internal/common"
	// m "cli-dungeon-crawler/internal/mobs"
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

// func (w *World) GenerateChunk(pt1, pt2 c.Position, ground string) {
// 	xRange := pt2.X - pt1.X
// 	yRange := pt2.Y - pt1.Y
// 	for x := range xRange {
// 		for y := range yRange {
// 			w.TileMap[c.Position{X: x, Y: y}] = Tile{
// 				Position:   c.Position{X: x, Y: y},
// 				GroundType: ground,
// 			}
// 			w.Entities[c.Position{X: x, Y: y}] = []Entity{{
// 				Position:    c.Position{X: x, Y: y},
// 				Name:        "Wall",
// 				Obstructing: true,
// 			},
// 			}
// 		}
// 	}
// 	walker := m.NewMob("walker")
// 	for i := 0; i < (xRange*yRange)/2; {
// 		walker.Behavior(walker)
// 		if _, ok := w.Entities[walker.Position]; ok {
// 			i++
// 			delete(w.Entities, walker.Position)
// 		}
// 	}
// }

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
