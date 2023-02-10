package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(elems []int, windowc int) {
	s, e := 1, 1

	deltas := make([]int, len(elems)+1)
	summer := NewSummer(elems)

	var sprev, eprev int
	slid := func() {
		// I don't know what the fuck this does. I just know that it works.
		// Sam showed it to me. Thank him.
		// TODO: what does this do?
		switch {
		case eprev != e:
			deltas[s-1]++
		case sprev != s:
			deltas[e]--
		}

		// log.Printf("for (%d, %d), freq = %v", s, e, freq)

		sprev = s
		eprev = e
	}

	slid()
	for s <= len(elems) {
		if e+1 > len(elems) || summer.Sum(s-1, e+2-1) > windowc {
			s++
		} else {
			e++
		}
		slid()
	}

	freq := make([]int, len(elems))
	freq[0] = deltas[0]
	for i := 1; i < len(freq); i++ {
		freq[i] = freq[i-1] + deltas[i]
	}

	for _, f := range freq {
		fmt.Println(f)
	}
}

// Summer does prefix sum.
type Summer struct {
	sums []int
}

// NewSummer returns a new Summer.
func NewSummer(elems []int) Summer {
	sums := make([]int, len(elems))
	for i := range elems {
		if i == 0 {
			sums[i] = elems[i]
		} else {
			sums[i] = elems[i] + sums[i-1]
		}
	}
	return Summer{sums}
}

// Sum returns the sum of elems[i:j] in O(1) time.
func (s Summer) Sum(i, j int) int {
	if i > j {
		return 0
	}
	if i == 0 {
		return s.sums[j-1]
	}
	return s.sums[j-1] - s.sums[i-1]
}

func main() {
	var nelems, windowc int
	mustSscanf(mustReadUntil('\n'), "%d %d", &nelems, &windowc)

	var elems []int
	for i := 0; i < nelems; i++ {
		var w int
		mustSscanf(mustReadUntil(' '), "%d", &w)
		elems = append(elems, w)
	}

	solve(elems, windowc)
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
