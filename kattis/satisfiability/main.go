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
		doCase()
	}
}

func doCase() {
	parts := strings.Split(readLine(), " ")

	nvars, _ := strconv.Atoi(parts[0]) // for_each(nvars) : n -> x_n
	generateBools := newBoolsGenerator(nvars)

	nclauses, _ := strconv.Atoi(parts[1])
	clausesStrs := make([]string, nclauses)
	for i := range clausesStrs {
		clausesStrs[i] = readLine()
	}

	clauses := parseClauses(clausesStrs)

generateLoop:
	for {
		vals := generateBools()
		if vals == nil {
			break
		}

		// log.Printf("vals: %v", vals)
		for _, clause := range clauses {
			result := clause.Eval(vals)
			// log.Printf("  evaluating %v = %v", clause, result)
			if !result {
				// log.Println("    clause failed")
				continue generateLoop
			}
		}

		// We found a solution!
		// log.Println("    all clauses passed")
		fmt.Println("satisfiable")
		return
	}

	// We couldn't find a solution, so print unsatisfiable.
	fmt.Println("unsatisfiable")
}

type Token struct {
	idx int16
	typ tokenType
	not bool
}

type tokenType int8

const (
	tokenTypeUnknown tokenType = iota
	tokenTypeVariable
	tokenTypeOperatorOR
)

func (t Token) val(vals []bool) bool {
	val := vals[t.idx-1]
	if t.not {
		val = !val
	}
	return val
}

func (t Token) String() string {
	switch t.typ {
	case tokenTypeVariable:
		n := fmt.Sprintf("X%d", t.idx)
		if t.not {
			n = "~" + n
		}
		return n
	case tokenTypeOperatorOR:
		return "v"
	default:
		return fmt.Sprintf("%#v", t)
	}
}

type Clause []Token

func parseClauses(clauses []string) []Clause {
	out := make([]Clause, len(clauses))
	for i, clause := range clauses {
		out[i] = make(Clause, 0, strings.Count(clause, " ")+1)
		words := strings.Split(clause, " ")
		for _, word := range words {
			switch word {
			case "v":
				out[i] = append(out[i], Token{
					typ: tokenTypeOperatorOR,
				})
			default:
				not := false
				if strings.HasPrefix(word, "~") {
					not = true
					word = word[1:]
				}

				if !strings.HasPrefix(word, "X") {
					panic("invalid variable")
				}

				word = word[1:]

				idx, err := strconv.ParseInt(word, 10, 16)
				if err != nil {
					panic("cannot parse int: " + err.Error())
				}

				out[i] = append(out[i], Token{
					idx: int16(idx),
					typ: tokenTypeVariable,
					not: not,
				})
			}
		}
	}

	return out
}

func (c Clause) Eval(vals []bool) bool {
	state := false
	for i := 0; i < len(c); i++ {
		switch token := c[i]; token.typ {
		case tokenTypeVariable:
			state = token.val(vals)
		case tokenTypeOperatorOR:
			next := c[i+1].val(vals)
			state = state || next
			// Skip the next token, since we already evaluated it.
			i++
		default:
			panic("invalid token type")
		}
	}

	return state
}

func (c Clause) String() string {
	var out strings.Builder
	out.Grow(32)
	for i, token := range c {
		if i != 0 {
			out.WriteString(" ")
		}
		out.WriteString(token.String())
	}
	return out.String()
}

func newBoolsGenerator(length int) func() []bool {
	out := make([]bool, length)
	i := 0
	max := 1 << length
	return func() []bool {
		if i >= max {
			return nil
		}

		for j := 0; j < length; j++ {
			out[j] = (i & (1 << j)) != 0
		}

		i++
		return out
	}
}

var stdinBuf = bufio.NewReader(os.Stdin)

func readLine() string {
	b, err := stdinBuf.ReadSlice('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSuffix(string(b), "\n")
}
