package rover

//Coordinate represents a point in space.
type Coordinate struct {
	x int
	y int
}

func (coords Coordinate) change(adder Coordinate) Coordinate {
	return Coordinate{coords.x + adder.x, coords.y + adder.y}
}
