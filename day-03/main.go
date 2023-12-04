package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var schematic [][]rune
	var lineIndex int

	for scanner.Scan() {
		schematic = append(schematic, []rune(fmt.Sprintf(".%s.", scanner.Text())))
		lineIndex++
	}
	padding := [][]rune{[]rune(strings.Repeat(".", len(schematic[0])))}

	paddedSchematic := append(padding, schematic...)
	paddedSchematic = append(paddedSchematic, padding...)

	var stashValue []int

	stashRow := make([][]int, lineIndex+2)
	for i := 0; i < lineIndex+2; i++ {
		stashRow[i] = make([]int, len(schematic[0]))
	}

	currentNumberOfPieces := 1

	total := 0
	for i := 1; i < lineIndex+2; i++ {
		for j := 1; j < len(schematic[0])-1; j++ {
			digit := ""
			neighbors := make([]rune, 0)

			for {
				r := paddedSchematic[i][j]

				if !unicode.IsDigit(r) {
					stashRow[i][j] = 0
					if len(neighbors) > 0 {
						isValid := strings.ContainsAny(string(neighbors), "*#$+@!%&'()[]{}-/\\\"`><;:^~=")
						if !isValid {
							fmt.Println(digit, "\t", string(neighbors))
						} else {
							integerNumber, _ := strconv.Atoi(digit)
							fmt.Println(digit, "\t", string(neighbors))

							currentNumberOfPieces++
							stashValue = append(stashValue, integerNumber)

							total += integerNumber
						}
					}
					if r == '*' {
						stashRow[i][j] = -1
					}
					break
				}

				neighbors = append(neighbors, paddedSchematic[i-1][j-1:j+2]...)
				neighbors = append(neighbors, paddedSchematic[i][j-1:j+2]...)
				neighbors = append(neighbors, paddedSchematic[i+1][j-1:j+2]...)

				digit += string(r)
				stashRow[i][j] = currentNumberOfPieces
				j++
			}
		}
	}

	sumOfProducts := 0
	for i := 0; i < len(stashRow)-1; i++ {
		for j := 0; j < len(stashRow[0]); j++ {
			s := stashRow[i][j]
			if s < 0 {
				var smallNeighbors []int
				smallNeighbors = append(smallNeighbors, stashRow[i-1][j-1:j+2]...)
				smallNeighbors = append(smallNeighbors, stashRow[i][j-1])
				smallNeighbors = append(smallNeighbors, stashRow[i][j+1])
				smallNeighbors = append(smallNeighbors, stashRow[i+1][j-1:j+2]...)

				nonZero := 0
				product := 1
				last := -1
				for _, value := range smallNeighbors {

					if value > 0 && last != value {
						last = value
						nonZero++
						fmt.Println("Stash Value: ", value, " - ", stashValue[value-1])
						product *= stashValue[value-1]
					}

				}
				fmt.Println("Non-zeros: ", nonZero)
				if nonZero > 1 {
					sumOfProducts += product
				}
			}
		}
	}
	fmt.Println(sumOfProducts)
}
