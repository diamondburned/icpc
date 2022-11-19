package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var stdinBuf = bufio.NewReader(os.Stdin)

func main() {
	nCases, _ := strconv.Atoi(readLine())
	for n := 0; n < nCases; n++ {
		nCandidates, _ := strconv.Atoi(readLine())
		candidateVotes := make([]int, nCandidates)
		for i := 0; i < nCandidates; i++ {
			votes, _ := strconv.Atoi(readLine())
			candidateVotes[i] = votes
		}

		winner := 0
		nWinnerVotes := 0
		sumVotes := 0
		firstVote := candidateVotes[0]
		maxVote := firstVote

		for i, vote := range candidateVotes {
			sumVotes += vote

			if vote > maxVote {
				winner = i
				maxVote = vote
				nWinnerVotes = 1
				continue
			}

			if vote == maxVote {
				nWinnerVotes++
				continue
			}
		}

		// spew.Dump(candidateVotes)
		// log.Println("winner:", winner+1, maxVote)
		if nWinnerVotes > 1 {
			fmt.Println("no winner")
		} else if maxVote > sumVotes/2 {
			fmt.Println("majority winner", winner+1)
		} else {
			fmt.Println("minority winner", winner+1)
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
