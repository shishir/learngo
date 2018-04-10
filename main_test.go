package main

import (
	"testing"
)

func TestRoverInitialization(t *testing.T) {
	rover := *NewRover(1, 1, "N")
	expectedCoords := Coordinate{1, 1}
	if rover.coords != expectedCoords && rover.direction != "N" {
		t.Errorf("rover not initialized")
	}
}

func TestRoverOnlyAcceptsLegitDirections(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Rover accepted incorrect direction")
		}
	}()
	rover := *NewRover(1, 1, "M")
	rover.valid()
}

func TestRoverCanTurnLeft(t *testing.T) {
	rover := *NewRover(1, 1, "N")
	rover.turnLeft()
	if "W" != rover.direction {
		t.Errorf("Direction invalid. Should be W was %s", rover.direction)
	}
	rover.turnLeft()
	if "S" != rover.direction {
		t.Errorf("Direction invalid. Should be S was %s", rover.direction)
	}
	rover.turnLeft()
	if "E" != rover.direction {
		t.Errorf("Direction invalid. Should be S was %s", rover.direction)
	}
	rover.turnLeft()
	if "N" != rover.direction {
		t.Errorf("Direction invalid. Should be S was %s", rover.direction)
	}
}

func TestRoverCanTurnRight(t *testing.T) {
	rover := *NewRover(1, 1, "N")
	rover.turnRight()
	if "E" != rover.direction {
		t.Errorf("Direction invalid. Should be W was %s", rover.direction)
	}
	rover.turnRight()
	if "S" != rover.direction {
		t.Errorf("Direction invalid. Should be S was %s", rover.direction)
	}
	rover.turnRight()
	if "W" != rover.direction {
		t.Errorf("Direction invalid. Should be S was %s", rover.direction)
	}
	rover.turnRight()
	if "N" != rover.direction {
		t.Errorf("Direction invalid. Should be S was %s", rover.direction)
	}
}

func TestRoverCanMove(t *testing.T) {
	rover1 := *NewRover(1, 1, "N")
	rover1.move()
	expectCoords := Coordinate{1, 0}
	if expectCoords != rover1.coords {
		t.Errorf("Invalid. expected %v, got %v", expectCoords, rover1.coords)
	}
	rover2 := NewRover(1, 1, "E")
	rover2.move()
	expectCoords = Coordinate{2, 1}
	if expectCoords != rover2.coords {
		t.Errorf("Invalid. expected %v, got %v", expectCoords, rover2.coords)
	}

	rover3 := NewRover(1, 1, "W")
	rover3.move()
	expectCoords = Coordinate{0, 1}
	if expectCoords != rover3.coords {
		t.Errorf("Invalid. expected %v, got %v", expectCoords, rover3.coords)
	}

	rover4 := NewRover(1, 1, "S")
	rover4.move()
	expectCoords = Coordinate{1, 2}
	if expectCoords != rover4.coords {
		t.Errorf("Invalid. expected %v, got %v", expectCoords, rover4.coords)
	}

}
