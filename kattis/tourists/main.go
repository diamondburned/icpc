package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func solve(cities Cities) {
	// Requirements for traveling: the next city must be:
	//
	//    - next.number > from.number
	//    - next.number % from.number == 0

	var steps int
	walker := newDistanceWalker(cities)

	for i := range cities[:len(cities)/2] {
		from := &cities[i]
		var counter int
		then := time.Now()
		for j := from.Number * 2; j <= int32(len(cities)); j += from.Number {
			if counter++; counter%100 == 0 {
				now := time.Now()
				log.Println("walking from", from.Number, "to", j, "took", now.Sub(then))
				log.Println("  current cache size:", len(walker.cache))
				log.Println("  cache hits:", walker.hits/walker.total*100)
				log.Println("  cache sets:", walker.sets)
				then = now
			}
			next := cities.At(j)
			dist := walker.distance(from, next)
			steps += dist
		}
	}

	fmt.Println(steps)
}

type distanceWalker struct {
	cache  distanceCache
	cities Cities
	pivot  int32
	paths  sharedArray[int32]

	hits  float64
	sets  float64
	total float64
}

func newDistanceWalker(cities Cities) *distanceWalker {
	return &distanceWalker{
		cache:  make(distanceCache, len(cities)*2),
		cities: cities,
		pivot:  int32(len(cities) / 2),
		paths:  newSharedArray[int32](len(cities) / 2),
	}
}

func (w *distanceWalker) walker(from, to *City, path sharedArray[int32], dist int) int {
	stack := []*City{from}

	for len(stack) > 0 {
		from := stack[len(stack)-1]  // top
		stack = stack[:len(stack)-1] // pop

		w.total++
		newDist, ok := w.cache.distance(from, to)
		if ok {
			w.hits++
			// log.Println("cache hit: from", from.Number, "to", to.Number, "dist", newDist)
			return newDist
		}

		newDist = dist + 1

		for _, neighbor := range from.Neighbors {
			if neighbor == to.Number {
				return newDist
			}
		}

		// Opportunistically cache the distance from the previous cities
		// within path to the current city. We want to stop at the n/2th
		// city because the rest is never visited. Also, we only want to
		// cache when the distance is over, I don't know, d? This means
		// skipping the last d-1 cities.
		const minDist = 4
		if path.len() > minDist && from.Number <= w.pivot {
			// The cached path's destination is where we're currently at.
			to := from

			for i := path.len() - minDist; i >= 0; i-- {
				// The cached path's source is the city that we were in.
				from := w.cities.At(path.at(i))
				dist := path.len() + 1 - i

				// log.Println("set cache: from", from.Number, "to", to.Number, "dist", dist)
				w.sets++

				w.cache.setDistance(from, to, dist)
			}
		}

		var prev *City
		if len(path.slice) > 0 {
			prevNo := path.slice[len(path.slice)-1]
			prev = w.cities.At(prevNo)
		}

		// Register the path to this city before traversing.
		path = path.append(from.Number)
		// log.Println("new path:", path.slice)

		for _, neighborNo := range from.Neighbors {
			// Since this is an undirected graph, we're storing the number on
			// both nodes. This means when walking from x to y, we might
			// accidentally walk back to x, causing an infinite cycle.
			if prev == nil || neighborNo != prev.Number {
				neighbor := w.cities.At(neighborNo)

				dist := w.walker(neighbor, to, path, newDist)
				if dist != -1 {
					return dist
				}
			}
		}
	}

	return -1
}

func (w *distanceWalker) distance(from, to *City) int {
	dist := w.walker(from, to, w.paths.new(), 1)
	if dist == -1 {
		log.Panicf("no path from %d to %d", from.Number, to.Number)
	}

	return dist
}

type City struct {
	Number    int32
	Neighbors []int32
}

type Cities []City

func (cs Cities) At(n int32) *City {
	return &cs[n-1]
}

// ParseCities parses the input into a graph of cities, returning the flat list
// of cities.
func ParseCities(n int, edges [][2]int) Cities {
	allCities := make([]City, n)
	for i := range allCities {
		allCities[i].Number = int32(i + 1)
	}

	for _, e := range edges {
		from := &allCities[e[0]-1]
		from.Neighbors = append(from.Neighbors, int32(e[1]))
		to := &allCities[e[1]-1]
		to.Neighbors = append(to.Neighbors, int32(e[0]))
	}

	return allCities
}

func main() {
	var n int
	mustScanLinef("%d", &n)

	var edges [][2]int
	for i := 0; i < n-1; i++ {
		var from, to int
		mustScanLinef("%d %d", &from, &to)
		edges = append(edges, [2]int{from, to})
	}

	cities := ParseCities(n, edges)
	log.Println("done parsing")

	solve(cities)
}

var stdin = bufio.NewReader(os.Stdin)

func mustScanLinef(f string, v ...any) {
	line, err := stdin.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimSpace(line)
	_, err = fmt.Sscanf(line, f, v...)
	if err != nil {
		panic(err)
	}
}

// dst -> src -> dist
type distanceCache map[*City]map[*City]int

func (c distanceCache) distance(src, dst *City) (int, bool) {
	// Force deterministic ordering.
	if dst.Number < src.Number {
		src, dst = dst, src
	}
	mdst, ok := c[dst]
	if !ok {
		return 0, false
	}
	d, ok := mdst[src]
	return d, ok
}

func (c distanceCache) setDistance(x, y *City, d int) {
	// Force deterministic ordering.
	if x.Number > y.Number {
		x, y = y, x
	}
	c[[2]*City{x, y}] = d
}

// sharedArray is like a slice, except if the backing array is reallocated
// during appending, then the old slice will also use the new backing array.
type sharedArray[T any] struct {
	slice []T
	root  *[]T
}

func newSharedArray[T any](cap int) sharedArray[T] {
	backing := make([]T, 0, cap)
	return sharedArray[T]{
		slice: backing,
		root:  &backing,
	}
}

func (a sharedArray[T]) append(v ...T) sharedArray[T] {
	a.slice = append(a.slice, v...)
	if cap(a.slice) != cap(*a.root) {
		// Update the slice to the one with the bigger backing array, but
		// pretend that we didn't append anything.
		*a.root = a.slice[:len(*a.root)]
	}
	return a
}

func (a sharedArray[T]) at(i int) T {
	return a.slice[i]
}

func (a sharedArray[T]) new() sharedArray[T] {
	a.slice = a.slice[:0]
	return a
}

func (a sharedArray[T]) len() int { return len(a.slice) }
