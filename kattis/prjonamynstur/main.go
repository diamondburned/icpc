package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

var yarns = map[byte]int{
	'.':  20,
	'O':  10,
	'\\': 25,
	'/':  25,
	'A':  35,
	'^':  5,
	'v':  22,
}

func main() {
	var x, y int
	fmt.Scanln(&x, &y)

	canvasBytes, _ := io.ReadAll(os.Stdin)
	canvas := bytes.Split(canvasBytes, []byte("\n"))

	var cost int
	for _, r := range canvas {
		for _, c := range r {
			cost += yarns[c]
		}
	}

	fmt.Println(cost)
}
