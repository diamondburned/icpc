package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// possible leaps, aka viable numbers:
// - 0
// - 2
// - 20
// - 22
// - 200
// - 202
// - 220
// - 222

var max = pow(10, 18)

func main() {
	var nstr string
	fmt.Scan(&nstr)

	// assert in format 1[0]+
	zeros := strings.Count(nstr[1:], "0")
	const baselineZeroes = 4 // 10^4 = 10000
	if nstr[0] == '1' && zeros == len(nstr)-1 && zeros > baselineZeroes {
		m := crunch(pow(10, baselineZeroes))

		// our max is zeroes=4 but the actual input max is zeroes=50.
		// m is m*2 for every 0 we add.
		for i := baselineZeroes; i < zeros; i++ {
			m *= 2
		}

		fmt.Println(m)
		return
	}

	n, err := strconv.Atoi(nstr)
	if err != nil {
		panic(err)
	}

	fmt.Println(crunch(n))
}

func crunch(n int) int {
	digits := len(itoa(n))
	counted := make(map[int]struct{})
	generate := new02sGenerator(digits)
	for {
		number := generate()
		if number == -1 {
			break
		}
		if number <= n {
			counted[number] = struct{}{}
		}
	}
	return len(counted)
}

func new02sGenerator(length int) func() int {
	out := make([]bool, length)
	i := 0
	max := 1 << length
	return func() int {
		if i >= max {
			return -1
		}

		for j := 0; j < length; j++ {
			out[j] = (i & (1 << j)) != 0
		}

		i++

		var n int
		for j, b := range out {
			if b {
				n += 2 * pow(10, length-j-1)
			}
		}

		return n
	}
}

func joinDigits(digits []int) int {
	var n int
	for i, d := range digits {
		n += d * pow(10, len(digits)-i-1)
	}
	return n
}

var itoa = strconv.Itoa

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
