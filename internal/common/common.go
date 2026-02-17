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

func Origin() Position {
	return Position{0, 0}
}

type Item struct {
	Name        string
	Description string
}

type Inventory struct {
	Slots int //temp
}
