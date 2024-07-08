package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var n int
	fmt.Sscanf(mustReadUntil('\n'), "%d", &n)

	var qaly float64
	for i := 0; i < n; i++ {
		var q, y float64
		fmt.Sscanf(mustReadUntil('\n'), "%f %f", &q, &y)
		qaly += q * y
	}

	fmt.Println(qaly)
}

var stdio = bufio.NewReader(os.Stdin)

func mustReadUntil(delim byte) string {
	s, err := stdio.ReadString(delim)
	if err != nil && err != io.EOF {
		panic(err)
	}
	return s
}
