package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	parts := strings.Fields(readLine())
	ncases, _ := strconv.Atoi(parts[0])
	maxTime, _ := strconv.Atoi(parts[1])

	cases := strings.Fields(readLine())
	var sum int
	for n := 0; n < ncases; n++ {
		v, _ := strconv.Atoi(cases[n])
		sum += v
		if sum > maxTime {
			fmt.Println(n)
			return
		}
	}

	fmt.Println(ncases)
}

var stdinBuf = bufio.NewReader(os.Stdin)

func readLine() string {
	b, err := stdinBuf.ReadSlice('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSuffix(string(b), "\n")
}
