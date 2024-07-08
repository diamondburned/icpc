package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Println(quadrant(x, y))
}

func quadrant(x, y int) int {
	switch {
	case x > 0 && y > 0:
		return 1
	case x < 0 && y > 0:
		return 2
	case x < 0 && y < 0:
		return 3
	case x > 0 && y < 0:
		return 4
	default:
		panic("invalid input")
	}
}
