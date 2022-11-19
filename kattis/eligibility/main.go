package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var stdinBuf = bufio.NewReader(os.Stdin)

func readLine() string {
	b, err := stdinBuf.ReadSlice('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSuffix(string(b), "\n")
}

func main() {
	ncases, _ := strconv.Atoi(readLine())

	for n := 0; n < ncases; n++ {
		words := strings.Fields(readLine())

		name := words[0]
		yearBeganPostSecondary := parseYear(words[1])
		yearOfBirth := parseYear(words[2])
		courses, _ := strconv.Atoi(words[3])

		if yearBeganPostSecondary >= 2010 {
			fmt.Println(name, "eligible")
			continue
		}

		if yearOfBirth >= 1991 {
			fmt.Println(name, "eligible")
			continue
		}

		// a student who has completed 41 courses or more is considered to have
		// more than 8 semesters of full-time study.
		if courses > 40 {
			fmt.Println(name, "ineligible")
			continue
		}

		fmt.Println(name, "coach petitions")
	}
}

func parseYear(str string) int {
	parts := strings.Split(str, "/")
	year, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	return year
}
