package common

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Position struct {
	X int
	Y int
}

func Pos(x, y int) Position {
	return Position{X: x, Y: y}
}

type Item struct {
	Name        string
	Description string
}

type Inventory struct {
	Capasity int
	Slots    map[int]Item
}
