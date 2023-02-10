// You can edit this code!
// Click here and start typing.
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/bits"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Weight is a physical weight in scales of 3.
type Weight = int

// WeightSum is a sum value that is the sum of some weights in scales of 3.
type WeightSum struct {
	Value  int
	Combis []Weight
}

// WeightSums is the three-sum table, which is the table of all combinations of
// multiples of 3s.
type WeightSums []WeightSum

// BuildWeightSum builds the three-sum table.
func BuildWeightSum(maxNum int) WeightSums {
	var weights []Weight
	for i := 0; len(weights) == 0 || weights[len(weights)-1] < maxNum; i++ {
		pow := int(math.Pow(3, float64(i)))
		weights = append(weights, pow)
	}

	sums := make([]WeightSum, (1 << len(weights)))
	for i := range sums {
		var sum int
		var combis []Weight
		for j := 0; j < bits.Len(uint(i)); j++ {
			if (i & (1 << j)) != 0 {
				sum += weights[j]
				combis = append(combis, weights[j])
			}
		}
		sums[i] = WeightSum{sum, combis}
	}

	return sums
}

// Find finds the three sum that is the nearest to the given number. The
// returned sum will be the next sum that is greater than the given number.
func (s WeightSums) Find(n int) WeightSum {
	i := sort.Search(len(s), func(i int) bool { return s[i].Value >= n })
	return s[i]
}

// Scale is a two-sided scale.
type Scale struct{ Left, Right []Weight }

// LeftWeight returns the weight of the left side of the scale.
func (s Scale) LeftWeight() int { return sum(s.Left) }

// RightWeight returns the weight of the right side of the scale.
func (s Scale) RightWeight() int { return sum(s.Right) }

// Compare returns the difference between the left and right side of the scale.
// If the left side is heavier, it returns a positive number. If the right side
// is heavier, it returns a negative number. If both sides are equal, it returns
// 0.
func (s Scale) Compare() int {
	// TODO: implement this more efficiently.
	return s.LeftWeight() - s.RightWeight()
}

// Balance balances the scale with the given object weight and returns the
// balanced scale.
func (s Scale) Balance(sumTable WeightSums) Scale {
	for {
		difference := s.Compare()
		if difference == 0 {
			break
		}

		// If difference is larger than 0, then we would add the next weight to
		// the right side to balance the s. If difference is smaller than 0,
		// then we would add the next weight to the left side instead.
		addingSide := &s.Right
		if difference < 0 {
			addingSide = &s.Left
			difference = -difference
		}

		// Find the next sum that is greater than the difference and add its
		// weights.
		nextWeight := sumTable.Find(difference)
		*addingSide = append(*addingSide, nextWeight.Combis...)
	}

	return s
}

func sum(v []int) int {
	var sum int
	for _, i := range v {
		sum += i
	}
	return sum
}

var (
	maxInput = int(1e9)
	sumTable = BuildWeightSum(maxInput)
)

func solve(objWeight int) {
	scale := Scale{Left: []Weight{objWeight}}

	scale = scale.Balance(sumTable)

	sort.Sort(sort.Reverse(sort.IntSlice(scale.Left)))
	sort.Sort(sort.Reverse(sort.IntSlice(scale.Right)))

	fmt.Println("left pan:", fmtInts(scale.Left[1:])) // Skip the object weight.
	fmt.Println("right pan:", fmtInts(scale.Right))
}

func fmtInts(ints []int) string {
	var b strings.Builder
	for i, n := range ints {
		b.WriteString(strconv.Itoa(n))
		if i != len(ints)-1 {
			b.WriteString(" ")
		}
	}
	return b.String()
}

func main() {
	var n int
	mustSscanf(mustReadUntil('\n'), "%d", &n)

	for i := 0; i < n; i++ {
		var objWeight int
		mustSscanf(mustReadUntil('\n'), "%d", &objWeight)
		solve(objWeight)
		fmt.Println()
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
