package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Dictionary struct {
	words    map[string][]string
	shortest map[string]string
}

func NewDictionary(n int) Dictionary {
	return Dictionary{
		words:    make(map[string][]string, n),
		shortest: make(map[string]string, n),
	}
}

func (d Dictionary) Add(word, alternative string) {
	d.words[word] = append(d.words[word], alternative)
	d.words[alternative] = append(d.words[alternative], word)
}

// Lookup returns the shortest alternative word for the given word. It builds
// the dictionary as it goes.
func (d Dictionary) Lookup(word string) string {
	shortestAlternative, ok := d.shortest[word]
	if ok {
		return shortestAlternative
	}

	allAlternatives := make(map[string]struct{})
	allAlternatives[word] = struct{}{}

	shortestAlternative = word

	// DFS
	stack := []string{word}
	for len(stack) > 0 {
		word := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Keep finding words similar to this one.
		alternatives, ok := d.words[word]
		if !ok {
			// No other alternatives found.
			continue
		}

		// Keep track of all alternatives words like this one.
		for _, alternative := range alternatives {
			// Don't add to lookup stack if we already checked the word.
			if _, ok := allAlternatives[alternative]; ok {
				continue
			}

			allAlternatives[alternative] = struct{}{}
			stack = append(stack, alternative)

			// Record this word if it's the shortest one yet.
			if len(alternative) < len(shortestAlternative) {
				shortestAlternative = alternative
			}
		}
	}

	// Cache our shortest words.
	for word := range allAlternatives {
		d.shortest[word] = shortestAlternative
	}

	return shortestAlternative
}

func main() {
	var n, m int
	mustSscanf(mustReadUntil('\n'), "%d %d", &n, &m)

	words := strings.Fields(mustReadUntil('\n'))

	dict := NewDictionary(m)
	for i := 0; i < m; i++ {
		var a, b string
		mustSscanf(mustReadUntil('\n'), "%s %s", &a, &b)
		dict.Add(a, b)
	}

	var totalLength int
	for _, w := range words {
		totalLength += len(dict.Lookup(w))
	}

	fmt.Println(totalLength)
}

var stdio = bufio.NewReader(os.Stdin)

func mustReadUntil(delim byte) string {
	s, err := stdio.ReadString(delim)
	if err != nil && err != io.EOF {
		panic(err)
	}
	return s
}

func mustSscanf(s, f string, v ...interface{}) {
	_, err := fmt.Sscanf(s, f, v...)
	if err != nil {
		panic(err)
	}
}
