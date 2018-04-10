package main

import "fmt"

func main() {
	rover := *NewRover(1, 1, "N")
	fmt.Println(rover)
}

//Coordinate represents a point in space.
type Coordinate struct {
	x int
	y int
}

func (coords Coordinate) changeX(value int) Coordinate {
	return Coordinate{coords.x + value, coords.y}
}
func (coords Coordinate) changeY(value int) Coordinate {
	return Coordinate{coords.x, coords.y + value}
}

// Rover encapsulate the rover class
type Rover struct {
	coords    Coordinate
	direction string
}

// NewRover constructs new rover
func NewRover(x int, y int, direction string) *Rover {
	coords := Coordinate{x, y}
	return &Rover{coords, direction}
}
func (rover Rover) valid() {
	correct := false
	for _, allowedDirection := range rover.validDirections() {
		if allowedDirection == rover.direction {
			correct = true
			break
		}
	}
	if !correct {
		panic("Incorrect direction")
	}
}

func (rover *Rover) turnLeft() {
	rover.turn(-1)
}

func (rover *Rover) turnRight() {
	rover.turn(+1)
}

func (rover *Rover) move() {
	switch rover.direction {
	case "N":
		rover.coords = rover.coords.changeY(-1)
	case "E":
		rover.coords = rover.coords.changeX(1)
	case "S":
		rover.coords = rover.coords.changeY(1)
	case "W":
		rover.coords = rover.coords.changeX(-1)
	}

}

func (rover *Rover) turn(sign int) {
	validDirections := rover.validDirections()
	nextIndex := rover.indexOfDirection(rover.direction) + sign
	if nextIndex >= 0 && nextIndex < len(validDirections) {
		rover.direction = validDirections[nextIndex]
	} else {
		rover.direction = validDirections[len(validDirections)-(sign*nextIndex)]
	}
}

func (rover Rover) validDirections() []string {
	return ([]string{"N", "E", "S", "W"})
}

func (rover Rover) indexOfDirection(direction string) int {
	for index, d := range rover.validDirections() {
		if d == direction {
			return index
		}
	}
	return -1
}
