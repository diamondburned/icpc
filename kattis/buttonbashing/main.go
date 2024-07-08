package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	var ncases int
	mustSscanf(mustReadUntil('\n'), "%d", &ncases)

	for i := 0; i < ncases; i++ {
		solve()
	}
}

func solve() {
	var nbuttons, t int
	mustSscanf(mustReadUntil('\n'), "%d %d", &nbuttons, &t)

	buttons := make([]int, nbuttons)
	for i, str := range strings.Fields(mustReadUntil('\n')) {
		mustSscanf(str, "%d", &buttons[i])
	}

	sort.Slice(buttons, func(i, j int) bool {
		return buttons[i] < buttons[j]
	})

	var presses int

	curr := t
	last := 0
	for {
		// log.Println("got", curr)

		button, ok := searchNearestButton(buttons, curr)
		if !ok {
			// log.Println("no button found")
			break
		}

		// log.Println("pressing button", button)

		last = curr
		curr = curr - button
		// log.Println("last:", last, "curr:", curr)

		if presses > 0 && abs(curr) > abs(last) {
			break
		}

		presses++
	}

	fmt.Println(presses, -last)
}

func searchNearestButton(buttons []int, val int) (button int, ok bool) {
	diff := int(math.MaxInt32)
	for i := len(buttons) - 1; i >= 0; i-- {
		// Ignore button 0.
		if buttons[i] == 0 {
			continue
		}
		curr := abs(val - buttons[i])
		if curr > diff {
			return
		}
		ok = true
		button = buttons[i]
		diff = curr
	}
	return
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
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
