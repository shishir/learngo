package main

import "fmt"
import "rover"

func main() {
	rover := *NewRover(1, 1, "N")
	fmt.Println(rover)
}
