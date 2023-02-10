package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	stdin, _ := io.ReadAll(os.Stdin)
	lines := strings.Split(strings.TrimSpace(string(stdin)), "\n")

	for len(lines) > 0 {
		lengths := strings.Fields(lines[0])
		lines = lines[1:]

		jackN, _ := strconv.Atoi(lengths[0])
		jillN, _ := strconv.Atoi(lengths[1])
		if jackN == 0 && jillN == 0 {
			break
		}

		jackCDs := make([]int, 0, jackN)
		jillCDs := make([]int, 0, jillN)

		for i := 0; i < jackN; i++ {
			n, _ := strconv.Atoi(lines[0])
			jackCDs = append(jackCDs, n)
			lines = lines[1:]
		}
		for i := 0; i < jillN; i++ {
			n, _ := strconv.Atoi(lines[0])
			jillCDs = append(jillCDs, n)
			lines = lines[1:]
		}

		solve(jackCDs, jillCDs)
	}
}

func solve(jackCDs, jillCDs []int) {
	var jackRanges [][2]int
	var jillRanges [][2]int

	collectRange(jackCDs, func(r [2]int) {
		jackRanges = append(jackRanges, r)
	})

	collectRange(jillCDs, func(r [2]int) {
		jillRanges = append(jillRanges, r)
	})

	// Find intersection of ranges
	intersections := intersection(jackRanges, jillRanges)

	// log.Println(jackRanges)
	// log.Println(jillRanges)
	// log.Println(intersections)

	var total int
	for _, r := range intersections {
		total += r[1] - r[0] + 1
	}

	fmt.Println(total)
}

func collectRange(input []int, f func([2]int)) {
	var start, end int
	for i, n := range input {
		if i == 0 {
			start = n
			end = n
			continue
		}

		if n == end+1 {
			end = n
			continue
		}

		f([2]int{start, end})
		start = n
		end = n
	}
	f([2]int{start, end})
}

func intersection(r1, r2 [][2]int) [][2]int {
	var result [][2]int
	var i, j int
	for i < len(r1) && j < len(r2) {
		// Left bound
		l := max(r1[i][0], r2[j][0])
		// Right bound
		r := min(r1[i][1], r2[j][1])
		// Valid intersection?
		if l <= r {
			result = append(result, [2]int{l, r})
		}
		// If right bound of r1 is smaller, increment i, else increment j.
		if r1[i][1] < r2[j][1] {
			i++
		} else {
			j++
		}
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
