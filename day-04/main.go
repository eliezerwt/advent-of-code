package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	total := 0.0

	scratchCardPile := make([]int, 0)
	scratchCardIndex := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")

		hits := 0

		games := strings.Split(strings.Trim(line[1], " "), "|")
		winningNumbers := strings.Split(strings.Trim(games[0], " "), " ")
		myNumbers := strings.Split(strings.Trim(games[1], " "), " ")

		for _, v := range myNumbers {
			for _, w := range winningNumbers {
				v = strings.Replace(v, " ", "", -1)
				w = strings.Replace(w, " ", "", -1)

				if v == w && (len(v)+len(w) > 0) {
					hits++
				}
			}
		}
		scratchCardIndex++
		scratchCardPile = append(scratchCardPile, hits)

		if hits >= 1 {
			total += (math.Pow(2, float64(hits)-1))
		}

	}

	totalScratchcards := 0
	scratchCardPrize := make([]int, len(scratchCardPile))
	for i := 0; i < len(scratchCardPile); i++ {
		scratchCardPrize[i] += 1
		for k := 0; k < scratchCardPrize[i]; k++ {
			for j := 0; j < scratchCardPile[i]; j++ {
				scratchCardPrize[i+j+1]++
			}
		}
		totalScratchcards += scratchCardPrize[i]
	}

	fmt.Println(total)
	fmt.Println(totalScratchcards)
}
