package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	nCasesLine := readLine()
	nCases, _ := strconv.Atoi(nCasesLine)

	var cities map[string]struct{}

	for nCase := 0; nCase < nCases; nCase++ {
		nCitiesLine := readLine()
		nCities, _ := strconv.Atoi(nCitiesLine)

		if cities == nil {
			cities = make(map[string]struct{}, nCities)
		} else {
			clear(cities)
		}

		for nCity := 0; nCity < nCities; nCity++ {
			city := readLine()
			cities[city] = struct{}{}
		}

		fmt.Println(len(cities))
	}
}

var stdinBuf = bufio.NewReader(os.Stdin)

func readLine() string {
	b, err := stdinBuf.ReadSlice('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSuffix(string(b), "\n")
}

func clear(m map[string]struct{}) {
	for k := range m {
		delete(m, k)
	}
}
