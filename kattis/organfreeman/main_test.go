package main

import (
	"strconv"
	"testing"
)

func BenchmarkDigitsDiv10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var n = 1_999
		var digits int
		for n > 0 {
			n /= 10
			digits++
		}
	}
}

func BenchmarkDigitsItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = len(strconv.Itoa(1_999))
	}
}
