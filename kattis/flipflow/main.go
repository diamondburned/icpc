package main

import (
	"fmt"
)

type Hourglass struct {
	total  int
	bottom int
}

func NewHourglass(sandWeight int) Hourglass {
	return Hourglass{
		total:  sandWeight,
		bottom: sandWeight,
	}
}

func (h *Hourglass) DropSand(w int) {
	h.bottom += w
	if h.bottom > h.total {
		h.bottom = h.total
	}
}

func (h *Hourglass) Flip() {

	h.bottom = h.total - h.bottom
}

func (h Hourglass) Top() int    { return h.total - h.bottom }
func (h Hourglass) Bottom() int { return h.bottom }

func (h Hourglass) log() {
	// log.Println("how's the hourglass lookin'?")
	// log.Print("  top mass:    ", h.Top())
	// log.Print("  bottom mass: ", h.Bottom())
}

func main() {
	var t, s, n int
	fmt.Scan(&t, &s, &n)

	hourglass := NewHourglass(s)
	var lastFlip int

	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)

		hourglass.log()

		// log.Println("dropped", a-lastFlip, "grams of sand so far")
		hourglass.DropSand(a - lastFlip)

		// log.Println("flipping...")
		hourglass.Flip()
		lastFlip = a
	}

	// log.Println("we've flipped enough times")
	// log.Println("the last time we flipped was", lastFlip)
	// log.Println("it is currently", t)

	hourglass.log()

	// Assume the sand keeps falling from the last time we flipped it until the
	// current time.
	hourglass.DropSand(t - lastFlip)

	// Print out the remaining sand.
	fmt.Println(hourglass.Top())
}
