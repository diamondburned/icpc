package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// Coordinate is a 2D coordinate.
type Coordinate struct{ X, Y int }

// Direction is a direction type.
type Direction uint8

const (
	North Direction = 'N'
	East  Direction = 'E'
	South Direction = 'S'
	West  Direction = 'W'
)

var directionDeltas = map[Direction]Coordinate{
	North: {0, 1},
	East:  {1, 0},
	South: {0, -1},
	West:  {-1, 0},
}

var directionOpposites = map[Direction]Direction{
	North: South,
	East:  West,
	South: North,
	West:  East,
}

// Add returns the sum of two coordinates.
func (c Coordinate) Add(other Coordinate) Coordinate {
	return Coordinate{c.X + other.X, c.Y + other.Y}
}

// Path is a sequence of directions.
type Path []Direction

// Walk walks the path.
func (p Path) Walk(f func(Coordinate, Direction)) {
	var coord Coordinate
	for _, dir := range p {
		coord = coord.Add(directionDeltas[dir])
		f(coord, dir)
	}
}

func walkToGraph(path Path) (start, end Coordinate, edges map[Coordinate][]Coordinate) {
	edges = make(map[Coordinate][]Coordinate)

	for _, dir := range path {
		next := end.Add(directionDeltas[dir])
		edges[end] = append(edges[end], next)
		edges[next] = append(edges[next], end)
		end = next
	}

	return
}

// Pathmap is a 2D array describing the ant paths. Each cell is the direction
// that we approached it from.
type PathMap map[Coordinate][]Direction

// NewPathMap returns a new PathMap with the given bounds.
func NewPathMap(cap int) PathMap {
	return make(map[Coordinate][]Direction, cap)
}

// Trace traces the path of the ant at the given coordinate walking in the
// given direction.
func (m PathMap) Trace(c Coordinate, v Direction) {
	m[c] = append(m[c], v)
	if len(m[c]) > 1 {
		log.Printf("  WARNING: path at %v is traversed %d times (%s)", c, len(m[c]), m[c])
	}
}

// BeenAt returns true if the ant has been at the given coordinate walking
// in the given direction or the opposite.
func (m PathMap) BeenAt(c Coordinate, d Direction) bool {
	d1 := d
	d2 := directionOpposites[d]
	for _, d := range m[c] {
		if d == d1 || d == d2 {
			return true
		}
	}
	return false
}

func solve(path Path) {
	start, end, edges := walkToGraph(path)
	// spew.Dump(root)
	// log.Println("starting at", root.Coordinate, "ending at", end.Coordinate)

	mostEfficientPathLength := mostEfficientPathLength(start, end, edges)
	// log.Printf("expecting %d steps", mostEfficientPathLength)
	// log.Printf("actual %d steps", len(path))

	fmt.Println(mostEfficientPathLength)
}

func mostEfficientPathLength(start, end Coordinate, edges map[Coordinate][]Coordinate) int {
	visited := make(map[Coordinate]bool)
	visited[start] = true

	var distance int

	// Do a simple BFS.
	queue := []Coordinate{start}
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			size--
			curr := queue[0]
			queue = queue[1:]

			if curr == end {
				return distance
			}

			for _, e := range edges[curr] {
				if !visited[e] {
					visited[e] = true
					queue = append(queue, e)
				}
			}
		}
		distance++
	}

	return distance
}

func main() {
	var npaths int
	mustSscanf(mustReadUntil('\n'), "%d\n", &npaths)
	mustReadUntil('\n') // skip empty line

	for i := 0; i < npaths; i++ {
		var pathlen int
		mustSscanf(mustReadUntil('\n'), "%d\n", &pathlen)

		path := make(Path, pathlen)
		for j := range path {
			var direction string
			mustSscanf(mustReadUntil('\n'), "%s", &direction)
			path[j] = Direction(direction[0])
		}

		solve(path)
		mustReadUntil('\n')
	}
}

var stdio = bufio.NewReader(os.Stdin)

func mustReadUntil(delim byte) string {
	s, err := stdio.ReadString(delim)
	if err != nil && err != io.EOF {
		panic(err)
	}
	return s
}

func mustSscanf(s, f string, v ...interface{}) {
	_, err := fmt.Sscanf(s, f, v...)
	if err != nil {
		panic(err)
	}
}
