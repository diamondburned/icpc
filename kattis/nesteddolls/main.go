package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	var ncases int
	fmt.Scanf("%d", &ncases)

	for i := 0; i < ncases; i++ {
		var ndolls int
		fmt.Scanf("%d", &ndolls)

		dolls := make([]Doll, 0, ndolls)
		for j := 0; j < ndolls; j++ {
			var w, h int
			fmt.Scanf("%d %d", &w, &h)
			dolls = append(dolls, NewDoll(w, h))
		}

		solve(dolls)
	}
}

type Doll struct {
	Width  int
	Height int
}

func NewDoll(w, h int) Doll {
	return Doll{w, h}
}

func (d Doll) LessThan(other Doll) bool {
	return d.Width < other.Width && d.Height < other.Height
}

func solve(dolls []Doll) {
	sort.Slice(dolls, func(i, j int) bool {
		d1, d2 := dolls[i], dolls[j]
		return d1.Width < d2.Width && d1.Height < d2.Height
	})

	// Precalculate in O(n) the next larger doll.
	larger := len(dolls) - 1
	for i := len(dolls) - 2; i >= 0; i-- {
		doll := &dolls[i]
		if !doll.LessThan(dolls[larger]) {
			doll.NextLarge = larger
			continue
		}

	}

	log.Println(dolls)
}
