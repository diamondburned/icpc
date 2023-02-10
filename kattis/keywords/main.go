package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	stdinb, _ := io.ReadAll(os.Stdin)
	stdin := string(stdinb)
	stdin = strings.TrimSpace(stdin)

	lines := strings.Split(stdin, "\n")
	lines = lines[1:] // ignore length

	rept := make(map[string]struct{})
	for _, line := range lines {
		log.Println(line)
		rept[sanitize(line)] = struct{}{}
	}

	fmt.Println(len(rept))
}

func sanitize(word string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsUpper(r) {
			return unicode.ToLower(r)
		}
		if r == '-' {
			return ' '
		}
		return r
	}, word)
}
