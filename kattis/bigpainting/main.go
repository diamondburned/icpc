package main

import (
	"bufio"
	"fmt"
	"image"
	"io"
	"os"
	"strings"

	"github.com/cubicdaiya/bms"
)

func main() {
	var paintingSz, masterpieceSz image.Point
	mustSscanf(
		mustReadUntil('\n'), "%d %d %d %d",
		&paintingSz.X, &paintingSz.Y, &masterpieceSz.X, &masterpieceSz.Y)

	// painting < picture always

	painting := make([]string, paintingSz.Y)
	for i := range painting {
		painting[i] = mustReadUntil('\n')[:paintingSz.X]
	}

	masterpiece := make([]string, masterpieceSz.Y)
	for i := range masterpiece {
		masterpiece[i] = mustReadUntil('\n')[:masterpieceSz.X]
	}

	paintingSks := make([]map[rune]int, paintingSz.Y)
	for i := range paintingSks {
		paintingSks[i] = bms.BuildSkipTable(painting[i])
	}

	var count int
	searchLine(masterpiece, painting, paintingSks, func(x, y int) {
		y2 := y + len(painting)
		x2 := x + len(painting[0])

		for y1 := y; y1 < y2; y1++ {
			if masterpiece[y1][x:x2] != painting[y1-y] {
				return
			}
		}

		count++
	})

	fmt.Println(count)
}

func searchLine(str, substr []string, substrSks []map[rune]int, f func(int, int)) {
	maxY := len(str) - len(substr)
	maxX := len(str[0]) - len(substr[0])

	for y := 0; y <= maxY; y++ {
		var x int
		for {
			i := bms.SearchBySkipTable(str[y][x:], substr[0], substrSks[0])
			i := strings.Index(str[y][x:], substr[0])
			if i == -1 {
				break
			}

			x += i
			if x > maxX {
				break
			}
			f(x, y)
			x++
		}
	}
}

type BMH2D struct {
	patternMatrix []string
	dvalue        int
	dgramVSkip    map[rune]int // default len(patternMatrix)
	dgramHPos     map[rune]int // default -1
	dgramHPosLL   []int        // size of len(patternMatrix[0])
	stripLength   int
}

const dgramMaxDValue = 3

func NewBMH2D(pattern []string) {
	var bmh BMH2D
	bmh.patternMatrix = pattern
	bmh
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

var stdio = bufio.NewReader(os.Stdin)

func mustReadUntil(delim byte) string {
	s, err := stdio.ReadString(delim)
	if err != nil && err != io.EOF {
		panic(err)
	}
	return strings.TrimSuffix(s, string(delim))
}

func mustSscanf(s, f string, v ...interface{}) {
	_, err := fmt.Sscanf(s, f, v...)
	if err != nil {
		panic(err)
	}
}
