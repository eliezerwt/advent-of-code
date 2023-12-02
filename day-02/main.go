package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	cubesAvailable := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	totalPossibleValue := 0
	setCubePower := 0

	for scanner.Scan() {

		minRequiredCubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		inputLineVector := strings.Split(scanner.Text(), ":")
		gameStateIsPossible := true

		gameId, _ := strconv.Atoi(strings.Split(inputLineVector[0], " ")[1])
		trials := strings.Split(inputLineVector[1], ";")

		for _, t := range trials {

			colors := strings.Split(t, ",")

			for _, c := range colors {
				gameSet := strings.Split(strings.Trim(c, " "), " ")
				cubeColor := gameSet[1]
				cubeAmount, _ := strconv.Atoi(gameSet[0])

				if cubeAmount > minRequiredCubes[cubeColor] {
					minRequiredCubes[cubeColor] = cubeAmount
				}

				if cubeAmount > cubesAvailable[cubeColor] {
					gameStateIsPossible = false
				}
			}
		}

		if gameStateIsPossible {
			totalPossibleValue += gameId
		}

		setCubePower += (minRequiredCubes["red"] * minRequiredCubes["blue"] * minRequiredCubes["green"])

	}

	fmt.Println("Total Possible Value: ", totalPossibleValue)
	fmt.Println("Total Set Cube Power: ", setCubePower)
}
