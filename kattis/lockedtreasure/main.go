package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ncases, _ := strconv.Atoi(readLine())
	for i := 0; i < ncases; i++ {
		parts := strings.Split(readLine(), " ")
		n, _ := strconv.Atoi(parts[0]) // n
		m, _ := strconv.Atoi(parts[1]) // m
		r := combinBinomial(n, m-1)
		fmt.Println(r)
	}
}

func combinBinomial(n, k int) int {
	// (n,k) = (n, n-k)
	if k > n/2 {
		k = n - k
	}
	b := 1
	for i := 1; i <= k; i++ {
		b = (n - k + i) * b / i
	}
	return b
}

var stdinBuf = bufio.NewReader(os.Stdin)

func readLine() string {
	b, err := stdinBuf.ReadSlice('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSuffix(string(b), "\n")
}
