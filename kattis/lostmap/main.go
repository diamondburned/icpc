package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

type Distance struct {
	From, To int
	Distance int
}

type Distances []Distance

func (d *Distances) AddDistance(from, to, distance int) {
	if from == to {
		return
	}
	if from > to {
		from, to = to, from
	}
	*d = append(*d, Distance{from, to, distance})
}

func main() {
	var n int
	mustSscanf(mustReadUntil('\n'), "%d", &n)

	distances := make(Distances, 0, n)
	for i := 0; i < n; i++ {
		line := strings.Fields(mustReadUntil('\n'))
		for j := 0; j < n; j++ {
			var d int
			mustSscanf(line[j], "%d", &d)
			distances.AddDistance(i, j, d)
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})

	for _, d := range distances {
		fmt.Println(d.From+1, d.To+1, "d =", d.Distance)
	}
	return
	// spew.Dump(distances)

	visited := make(map[[2]int]bool)
	roads := make([]Distance, 0, n)
	for _, distance := range distances {
		log.Println("checking", distance)
		if visited[[2]int{distance.From, distance.To}] {
			log.Println("  already visited")
			continue
		}

		log.Println("  not visited")
		visited[[2]int{distance.From, distance.To}] = true
		roads = append(roads, distance)
	}

	sort.Slice(roads, func(i, j int) bool {
		if roads[i].From == roads[j].From {
			return roads[i].To < roads[j].To
		}
		return roads[i].From < roads[j].From
	})

	for _, road := range roads {
		fmt.Println(road.From+1, road.To+1, "d =", road.Distance)
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
