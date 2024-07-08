package main

import (
	"fmt"
)

func main() {
	var l, n int
	fmt.Scan(&l, &n)

	// remain is how much paper we'll end up with assuming we use up all the
	// rolls to wipe.
	remain := l % n

	// Start with 1 roll.
	nrolls := 1

	// We're aiming for no remain. We'll need to buy enough rolls to wipe
	// without any remain.
	for remain != 0 {
		// Assume we'll always end up with remain paper left in this roll. We'll
		// do that by subtracting remain from how much paper we need per wipe
		// (n). This works, because every time we get to the next roll, we're
		// always assuming that we need (n - remain) more cm.
		n -= remain

		// Update remain.
		remain = l % n

		// We'll need to buy one more roll.
		nrolls++
	}

	fmt.Println(nrolls)
}

/*
func simulate(l, n int) {
	// Start with one toilet paper roll. We'll run through every scenario where
	// we shit and use a roll until we've reached a scenario where we've used
	// up all the rolls perfectly.
	nrolls := 1

hyperdimensionLoop:
	for {
		if nrolls > 10 {
			log.Panic("what the fuck?!")
		}

		log.Println("Starting with", nrolls, "rolls")
		rolls := make([]int, nrolls)
		for i := range rolls {
			rolls[i] = l
		}

	shittingLoop:
		for {
			// Do we need to buy more toilet paper?
			for i, roll := range rolls {
				if roll == 0 {
					// Buy this one.
					rolls[i] = l
				}
			}

			log.Println("Shitting... will be needing", n, "cm")

			// We're done shitting. Let's take a bit of toilet paper to wipe.
			// We'll need to take enough paper from our rolls.
			remain := n

			// Take from the last non-empty roll.
			for i := range rolls {
				if rolls[i] == 0 {
					log.Println("Roll", i, "is empty, grabbing the backup...")
					continue
				}

				log.Println("Taking", remain, "cm from roll", i, "which has", rolls[i], "cm left")
				rolls[i] -= remain

				if rolls[i] < 0 {
					// We ran out of paper on this roll. We'll need to take
					// more from the next roll.
					remain = -rolls[i]
					// Mark our roll as empty.
					rolls[i] = 0
				} else {
					remain = 0
				}

				log.Println("Need", remain, "cm more")

				if remain == 0 {
					log.Println("Done wiping")
					// We've taken enough paper to wipe. We're done shitting.
					continue shittingLoop
				}
			}

			// We ran out of all toilet paper rolls! Check if we still need any
			// more toilet paper.
			log.Println("Ran out of toilet paper!")

			if remain != 0 {
				log.Println("We still need", remain, "cm. Let's rewind the universe and add a roll")
				// Drats! We ran out of paper. We'll need to buy more rolls.
				// Let's rewind the universe and start shitting again with one
				// more roll.
				nrolls++
				continue hyperdimensionLoop
			}

			// We've bought enough rolls to wipe, so we have no crisis. This is
			// the universe!
			log.Println("Hooray! We've bought enough rolls to wipe. This is the universe!")
			break hyperdimensionLoop
		}
	}

	fmt.Println(nrolls)
}
*/
