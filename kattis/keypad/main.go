package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var stdinBuf = bufio.NewReader(os.Stdin)

type key uint8

const (
	notPressed key = iota
	pressed
	indeterminate
)

func (s key) String() string {
	switch s {
	case notPressed:
		return "N"
	case pressed:
		return "P"
	case indeterminate:
		return "I"
	default:
		return "?"
	}
}

type keypad [][]key

func makeKeypad(nrow, ncol int) keypad {
	keypad := make([][]key, nrow)
	for i := 0; i < nrow; i++ {
		keypad[i] = make([]key, ncol)
	}
	return keypad
}

func main() {
	nCases, _ := strconv.Atoi(readLine())
	for i := 0; i < nCases; i++ {
		firstLine := strings.Fields(readLine())
		nrow, _ := strconv.Atoi(firstLine[0])
		ncol, _ := strconv.Atoi(firstLine[1])

		keypad := makeKeypad(nrow, ncol)
		for row := 0; row < nrow; row++ {
			line := readLine()
			for col := 0; col < ncol; col++ {
				switch line[col] {
				case '0':
					keypad[row][col] = notPressed
				case '1':
					keypad[row][col] = pressed
				}
			}
		}
	}
}

func readLine() string {
	b, err := stdinBuf.ReadSlice('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSuffix(string(b), "\n")
}
