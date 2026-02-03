package world

import (
	"math/rand"

	"github.com/WaitWut862/cli-dungeon-crawler/internal/common"
	"github.com/WaitWut862/cli-dungeon-crawler/internal/mob"
)

type World struct {
	Tick     int
	TileMap  map[common.Position]Tile
	Entities map[common.Position][]Entity
}

type Tile struct {
	Position   common.Position
	GroundType string
}

type Entity struct {
	position    common.Position
	name        string
	obstructing bool
}


func (w *World) GenerateChunk(ptA, ptB common.Position) {
	walker := mob.NewMob(nil)
	walker.Position = common.Position{X: (ptA.X / 2), Y: (ptB.Y / 2)}
	walker.Facing = common.North

	area := (ptB.X - ptA.X) * (ptB.Y - ptA.Y)

	for floorTiles := 0; floorTiles < (area / 2); floorTiles++ {
		direction := rand.Intn(5)
		switch direction {
		case 0:
			walker.Move()
		case 1:
			walker.TurnLeft()
			walker.Move()
		case 2:
			walker.TurnRight()
			walker.Move()
		case 3:
			walker.TurnLeft()
			walker.TurnLeft()
			walker.Move()
		}
		w.TileMap[walker.Position] = Tile{
			Position:   walker.Position,
			GroundType: "floor",
		}
	}
	
}


func (w *World) MakeWorld() {
	w.TileMap = make(map[common.Position]Tile)
	w.Entities = make(map[common.Position][]Entity)
}

func (w *World) UpdateTick() {
	w.Tick = w.Tick + 1
}
