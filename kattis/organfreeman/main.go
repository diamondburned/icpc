package main

import (
	"fmt"
	"log"
)

func main() {
	test := func(x int) {
		fmt.Printf("%d,%d\n", x, f(x))
	}
	testij := func(i, j int) {
		for i != j {
			test(i)
			i++
		}
	}
	// testij(1, 20)
	// testij(100, 120)
	// testij(1000, 1020)
	testij(1_999_000, 1_999_020)
	testij(1_999_000_000, 1_999_000_020)
}

var maxDigits = 10

// Basically, when we're going from 1_999_000_000 to 1_999_000_020, we're
// adding a lot of the same numbers repeatedly, especially because we're adding
// by 1.
//
// In the future, we might consider binary search since even though the
// results dip down after a bit, they're still generally increasing. Maybe we
// can use that to estimate.

var factorials = [10]int{
	1,
	1,
	2,
	6,
	24,
	120,
	720,
	5040,
	40320,
	362880,
}

// Caching n-1 (where n = |input|) for 1999000016:
//
//     [0]: (1! + 9! + 9! + 9! + 0! + 0! + 0! + 0! + 1! + 6!)
//     [1]: (1! + 9! + 9! + 9! + 0! + 0! + 0! + 0! + 1!)
//     [2]: (1! + 9! + 9! + 9! + 0! + 0! + 0! + 0!)
//     [3]: (1! + 9! + 9! + 9! + 0! + 0! + 0!)
//     [4]: (1! + 9! + 9! + 9! + 0! + 0!)
//     [5]: (1! + 9! + 9! + 9! + 0!)
//     [6]: (1! + 9! + 9! + 9!)
//     [7]: (1! + 9! + 9!)
//     [8]: (1! + 9!)
//     [9]: (1!)
//
// Caching n-1 (where n = |input|) for 199900:
//
//     [0]: (1! + 9! + 9! + 9! + 0! + 0!)
//     [1]: (1! + 9! + 9! + 9! + 0!)
//     [2]: (1! + 9! + 9! + 9!)
//     [3]: (1! + 9! + 9!)
//     [4]: (1! + 9!)
//     [5]: (1!)
//     [6]: 0
//     [7]: 0
//     [8]: 0
//     [9]: 0
//
// We'll map each row to a number. That will be the index of the array. The
// integer at that index will be the sum of the factorials of the digits in the
// row. We'll carefully invalidate the cache every multiple of 10.
type facSumCache [10]int

var powersOf10 = [10]int{
	0: 1,             // correspond to the 0th one above
	1: 10,            // correspond to the 1st one above
	2: 100,           // correspond to the 2nd one above
	3: 1_000,         // correspond to the 3rd one above
	4: 10_000,        // correspond to the 4th one above
	5: 100_000,       // correspond to the 5th one above
	6: 1_000_000,     // correspond to the 6th one above
	7: 10_000_000,    // correspond to the 7th one above
	8: 100_000_000,   // correspond to the 8th one above
	9: 1_000_000_000, // correspond to the 9th one above
}

func fgenerator() func() (x, y int) {
	var sums facSumCache
	var last int
	var next10 = powersOf10[1]
	return func() (x, y int) {
		y = 0
		x = last + 1

		// Invalidate the cache if we're at any multiple of 10. We have to be
		// smart about this: if we're moving 10 up, we only need to invalidate
		// the first item in the cache. Not only that, but we can also use the
		// n+1 sum to calculate the n sum.
		if x == next10 {
			// Invalidate the cache. We basically count the number of zeros to
			// get our array index for sums. Since everything is 0 except for
			// the leading 1, we can just count the number of digits.
			j := digits(x)
			for i := 0; i < j; i++ {
				sums[i] = 0
			}
			// Update next10 to the next multiple of 10.
			next10 *= 10
		}

		// OR

		// Seek until we find the highest power of 10 that x is divisible by.
		// We'll use that to invalidate the cache at that index and
		//
		// TODO: we can keep track of this as we generate more numbers.
		for i := len(powersOf10) - 1; i >= 0; i-- {
			if x%powersOf10[i] == 0 {
				for j := 0; j <= i; j++ {
					sums[j] = 0
				}
				// Repopulate the cache using the sums we already have.
				for j := i; j != 0; j-- {
					// We're going back up, so we can use the n+1 sum to
					// calculate the n sum.
					sums[j] = sums[j+1] - factorials[j]
				}
				break
			}
		}
		if x%10 == 0 {
			sums[0] = 0
			for i := 1; i < 10; i++ {
				sums[i] = sums[i-1] + factorials[i]
			}
		}

		for x != 0 {
			r := x % 10
			x /= 10
			log.Println("add", r, "factorial", factorials[r])
			y += factorials[r]
		}
		if last%10 == 0 {
			sums = facSumCache{}
		}
		return
	}
}

func facSum(x int) int {
	var y int
	for x != 0 {
		r := x % 10
		x /= 10
		y += factorials[r]
	}
	return y
}

func f(x int) int {
	d := digits(x)
	n := 0
	for x != 0 {
		r := x % 10
		x /= 10
		log.Println("add", r, "factorial", factorials[r])
		n += factorials[r]
	}
	return n
}

var splitDigitsBuf [10]int

func splitDigits(x int, lhsLimit int) []int {

}

func digits(x int) int {
	var digits int
	for x != 0 {
		digits++
		x /= 10
	}
	return digits
}

func g(y int) int {
	// if y < 10 {
	// 	panic("unsupported")
	// }
	if y == 1 {
		return 0
	}
	fmt.Println(y)
	return g(y-1)/y + 1
}
