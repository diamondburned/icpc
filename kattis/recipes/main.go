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
		line := mustSplit(readLine(), 3)
		r, _ := strconv.Atoi(line[0])
		p, _ := strconv.Atoi(line[1]) // src proportion
		d, _ := strconv.Atoi(line[2]) // dst proportion

		recipe := parseRecipeStdin(r)
		scale := float64(d) / float64(p)

		mainIngr := recipe.mainIngredient()
		mainIngr.Weight *= scale

		for i, ingr := range recipe {
			recipe[i].Weight = mainIngr.Weight * float64(ingr.Percentage) / 100
		}

		fmt.Println("Recipe #", n+1)
		for _, ingr := range recipe {
			fmt.Printf("%s %.1f\n", ingr.Name, ingr.Weight)
		}
		fmt.Println("----------------------------------------")
	}
}

type ingredient struct {
	Name       string
	Weight     float64
	Percentage float64
}

func parseRecipe(line string) ingredient {
	words := mustSplit(line, 3)
	weight, _ := strconv.ParseFloat(words[1], 64)
	percentage, _ := strconv.ParseFloat(words[2], 64)
	return ingredient{
		Name:       words[0],
		Weight:     weight,
		Percentage: percentage,
	}
}

type recipe []ingredient

func parseRecipeStdin(n int) recipe {
	recipes := make([]ingredient, n)
	for i := 0; i < n; i++ {
		recipes[i] = parseRecipe(readLine())
	}
	return recipes
}

func (r recipe) mainIngredient() ingredient {
	for _, ingr := range r {
		if ingr.Percentage == 100 {
			return ingr
		}
	}
	panic("No main ingredient")
}

func readLine() string {
	b, err := stdinBuf.ReadSlice('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSuffix(string(b), "\n")
}

func mustSplit(s string, n int) []string {
	split := strings.Split(s, " ")
	if len(split) != n {
		panic("Invalid number of words")
	}
	return split
}
