package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func main() {
	stdinb, _ := io.ReadAll(os.Stdin)
	stdin := strings.TrimSpace(string(stdinb))

	var b strings.Builder
	umm := ummcode(stdin)
	for i := 0; i < len(umm); i += 7 {
		char := umm[i : i+7]
		b.WriteByte(ummbinary(char))
	}

	fmt.Println(b.String())
}

func ummbinary(ummword string) byte {
	var binary byte
	for _, b := range ummword {
		switch b {
		case 'u': // 1
			binary = (binary << 1) | 1
		case 'm': // 0
			binary = binary << 1
		}
	}
	// log.Printf("%q = %b", ummword, binary)
	return binary
}

func ummcode(input string) string {
	var ummcode strings.Builder
	words := strings.Fields(input)
	for _, word := range words {
		word = strings.Map(func(r rune) rune {
			if !unicode.IsLetter(r) {
				return -1
			}
			return r
		}, word)

		umm := strings.Map(func(r rune) rune {
			switch r {
			case 'u', 'm':
				return r
			default:
				return -1
			}
		}, word)
		if umm != word {
			continue
		}

		ummcode.WriteString(umm)
	}

	ummstr := ummcode.String()
	return ummstr
}
