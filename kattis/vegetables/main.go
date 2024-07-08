package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
   2000      3000           4000, target = 1000
1  2000      3000           3000           1000
2  2000      2000      1000 3000           1000
3  2000      2000      1000 2000 1000      1000
4  1000 1000 2000      1000 2000 1000      1000
5  1000 1000 1000 1000 1000      2000 1000 1000
6  1000 1000 1000 1000 1000 1000 1000 1000 1000
*/

/*
   1000    1400, target = 500
1  1000    900     500
2  500 500 900     500
3  500 500 500 400 500
*/

func main() {
	var margin float64
	var N int
	fmt.Scanf("%f %d", &margin, &N)
	if N == 0 {
		fmt.Println(0)
		return
	}

	chunks := make(Chunks, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &chunks[i])
	}
	// log.Println(chunks)

	const maxCuts = 500

	tryCutting := func(target int) int {
		chunks := append(Chunks(nil), chunks...)
		var cuts int

		for cuts < maxCuts && !chunks.IsWithinMargin(margin) {
			maxIx := chunks.MaxIx()
			chunks.Cut(maxIx, target)
			cuts++
		}

		return cuts
	}

	min := chunks[chunks.MinIx()]
	cuts := tryCutting(min)

	sort.Search(min, func(target int) bool {
		cuts2 := tryCutting(target)
		split := cuts2 < cuts
		cuts = cuts2
		return split
	})

	fmt.Println(cuts)
}

type Chunks []int

func (c Chunks) Len() int           { return len(c) }
func (c Chunks) Less(i, j int) bool { return c[i] < c[j] }
func (c Chunks) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func (c Chunks) MinIx() int {
	min := 0
	for i, v := range c[1:] {
		if v < c[min] {
			min = i + 1
		}
	}
	return min
}

func (c Chunks) MaxIx() int {
	max := len(c) - 1
	for i := len(c) - 2; i >= 0; i-- {
		if c[i] > c[max] {
			max = i
		}
	}
	return max
}

// Cut cuts the chunk at the given index into two chunks, the first of which
// is the given size.
func (c *Chunks) Cut(at int, want int) {
	(*c)[at] -= want
	if (*c)[at] < 0 {
		panic("cutting too much")
	}
	*c = append(*c, want)
	sort.Sort(c)
}

// NextTarget finds the next target size to cut at.
func (c Chunks) NextTarget() int {
	var i int
	for {
		target := c[i] / 2
		if target > 0 {
			return target
		}
		i++
	}
	panic("no target found")
}

// IsWithinMargin returns true if the largest chunk is within the given margin
// of the smallest chunk.
func (c Chunks) IsWithinMargin(margin float64) bool {
	min := c[c.MinIx()]
	max := c[c.MaxIx()]
	return float64(min)/float64(max) >= margin
}

func (c Chunks) String() string {
	nstrings := make([]string, len(c))
	for i, v := range c {
		nstrings[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(nstrings, " ")
}
