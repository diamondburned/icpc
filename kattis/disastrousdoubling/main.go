package main

import (
	"fmt"
	"math"
	"math/big"
)

var modulo = int64(math.Pow10(9)) + 7

func main() {
	var nexps int
	fmt.Scanf("%d", &nexps)

	// Represent the population bacteria as a power of modulo.
	// Meaning bacteriaReal = modulo^bacteriaMul + bacteriaRem
	bacteria := big.NewInt(1)

	for i := 0; i < nexps; i++ {
		bacteria = bacteria.Lsh(bacteria, 1)

		var used int64
		fmt.Scanf("%d", &used)

		bacteria = bacteria.Sub(bacteria, big.NewInt(used))
		if bacteria.Sign() < 0 {
			fmt.Println("error")
			return
		}
	}

	moduloBig := big.NewInt(modulo)
	fmt.Println(bacteria.Mod(bacteria, moduloBig))
}
