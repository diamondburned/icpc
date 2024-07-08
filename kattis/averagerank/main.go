package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type player struct {
	id       int
	points   int
	sumRanks int
}

func main() {
	var nPlayers, nWeeks int
	mustSscanf(mustReadUntil('\n'), "%d %d", &nPlayers, &nWeeks)

	players := make([]player, nPlayers)
	for i := range players {
		player := &players[i]
		player.id = i + 1
	}

	ranking := make([]int, nPlayers)
	for i := range ranking {
		ranking[i] = i
	}

	rankingOfPlayer := func(playerIx int) int {
		return sort.Search(len(ranking), func(i int) bool {
			return ranking[i] == playerIx
		})
	}

	for w := 1; w <= nWeeks; w++ {
		words := strings.Fields(mustReadUntil('\n'))
		// log.Println(words)

		var nWinners int
		mustSscanf(words[0], "%d", &nWinners)

		winners := make([]int, nWinners)
		for i := range winners {
			mustSscanf(words[i+1], "%d", &winners[i])
			winners[i]--
		}

		for _, id := range winners {
			player := &players[id]
			player.points++
		}

		// Rerank only the players that have changed. Using sort.Slice will TLE,
		// so we use a custom sort function.
		for _, id := range winners {
			currentID := id
			currentRank := rankingOfPlayer(currentID)
			currentPlayer := players[currentID]

			nextRankedID := currentID
			nextRankedRank := currentRank

			// Move our cursor up to the front of the ranking until we find a
			// player with more points than us.
			for nextRankedRank > 0 {
				nextRankedRank--
				nextRankedID = ranking[nextRankedRank]
				if players[nextRankedID].points > currentPlayer.points {
					break
				}
			}

			// Shift all players between the current player and the next ranked
			// player down one rank.
			copy(ranking[nextRankedRank+1:currentRank+1], ranking[nextRankedRank:currentRank])
			// Insert the current player at the next ranked player's rank.
			ranking[nextRankedRank] = currentID
		}

		eachRank(ranking, players, func(i, rank int) {
			player := &players[i]
			player.sumRanks += rank
			// log.Println("player", player.id, "rank", rank, "sum", player.sumRanks)
		})
	}

	for _, player := range players {
		// log.Printf("%#v", player)
		fmt.Printf("%.07f\n", float64(player.sumRanks)/float64(nWeeks))
	}
}

func eachRank(ranking []int, players []player, f func(i, rank int)) {
	var lastRank int
	var lastTotal int

	for i, playerIx := range ranking {
		rank := i + 1
		player := players[playerIx]
		if i > 0 && lastTotal == player.points {
			rank = lastRank
		} else {
			lastTotal = player.points
			lastRank = rank
		}

		f(playerIx, rank)
	}
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
