package main

import (
	"fmt"
	"log"
)

/*
4 stages:

1. Rain falls. The water fills up to the leak. No loss of water.
2. Rain falls. The water fills up beyond the leak. Water is lost over time.
3. Rain stops. The water still leaks. Water is lost over time.
4. Meter is read. The water level is reported.
*/

func main() {
	// L: where the leak is (mm)
	// K: rate of leak (mm/h)
	// T1: duration of rain (h)
	// T2: duration of no rain (h)
	// H: height of rain when checked (mm)
	var L, K, T1, T2, H float64
	fmt.Scanf("%f %f %f %f %f", &L, &K, &T1, &T2, &H)

	var F1, F2 float64 // min guess, max guess

	// We definitely have at least H mm of water.
	F1 += H
	F2 += H

	if H >= L {
		// If T1 is the total duration of rainfall, then we can calculate for
		// the time that the water was not leaking.
		TnotLeaking := T1 - (L / K)
		log.Println("TnotLeaking:", TnotLeaking)

		// If we are above the leak, then we might have lost some water.
		// The maximum amount of water we could have lost is the total leaked
		// from the time the rain stopped until the meter was read.
		loss := K * T1
		F2 += loss
	}

	fmt.Printf("%.6f %.6f", F1, F2)
}
