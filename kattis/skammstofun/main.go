package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	stdin, _ := io.ReadAll(os.Stdin)
	words := strings.Split(string(stdin), "\n")[1]
	for _, word := range strings.Fields(words) {
		fmt.Print(word[0])
	}
	fmt.Println()
}
