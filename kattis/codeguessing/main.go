package main

import "fmt"

func main() {
	var p, q int
	fmt.Scan(&p, &q)

	ours := make([]int, 0, 2)

	var cardstr string
	fmt.Scan(&cardstr)
	for i, s := range cardstr {
		switch s {
		case 'A':
			ours = append(ours, i)
		}
	}

	ourRange := ours[1] - ours[0]
	ourDifference := q - p

	if ourRange == ourDifference {
		switch ourDifference {
		case 1:
			// aabb, baab or bbaa
			switch ours[0] {
			case 0: // aabb
				if q+2 < 10 {
					fmt.Println(q+1, q+2)
					return
				}
			case 1: // baab
				if q+1 < 10 && p-1 > 0 {
					fmt.Println(p-1, q+1)
					return
				}
			case 2: // bbaa
				if p-2 > 0 {
					fmt.Println(p-2, p-1)
					return
				}
			}
		case 2:
			// either abab or baba
			switch ours[0] {
			case 0: // abab
				if q+1 < 10 {
					fmt.Println(p+1, q+1)
					return
				}
			case 1: // baba
				if p-1 > 0 {
					fmt.Println(p-1, q-1)
					return
				}
			}
		case 3:
			// we're outer, so print inner
			fmt.Println(p+1, q-1)
			return
		}
	}

	fmt.Println(-1)
}
