package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	Data  int64
	Costs []int
}

func main() {
	stdin, _ := io.ReadAll(os.Stdin)
	lines := strings.Split(string(stdin), "\n")

	data, _ := strconv.ParseInt(lines[0], 2, 64)

	var costs []int
	for _, line := range strings.Fields(lines[1]) {
		c, _ := strconv.Atoi(line)
		costs = append(costs, c)
	}

	solve(data, costs)
}

const maxLen = 12

func solve(data int64, costs []int) {
	for _, cost := range costs {
		log.Printf("0b%b", cost)
	}
}

func gnerateSums(n int, f func([]int)) {
	switch n {
	case 1:
	case 2:
	case 3:
	case 4:
	case 5:
	case 6:
	case 7:
	case 8:
	case 9:
	case 10:
	case 11:
	case 12:
	}
}
