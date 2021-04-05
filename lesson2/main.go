package main

import (
	"fmt"

	"github.com/nj-eka/gb_go_1/lesson2/measure"
)

func main() {
	fmt.Println("Select: [1, 2, 3]")
	fmt.Println("1 - Rectangle Area.")
	fmt.Println("2 - Circle D L")
	fmt.Println("3 - Number Parts")

	var choice string
	fmt.Scan(&choice)
	switch choice {
	case "1":
		measure.RectangleS()
	case "2":
		measure.CircleDL()
	case "3":
		measure.NumberParts()
	default:
		fmt.Println("Unknown operation")
	}

}
