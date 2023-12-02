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

	var calibration_value int

	for scanner.Scan() {
		inputText := scanner.Text()
		adjustedLine := []byte(replaceWrittenDigits(inputText))

		first := string(adjustedLine[0])
		last := string(adjustedLine[len(adjustedLine)-1])

		integer_calibration, _ := strconv.Atoi(fmt.Sprint(first, last))
		calibration_value += integer_calibration
	}

	fmt.Println(calibration_value)

}

func replaceWrittenDigits(line string) string {

	parsedLine := line

	digits := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}

	var idxValue map[int]string = make(map[int]string)

	for _, d := range digits {
		// Cursor to sweep whole word looking for multiples instances of the same number
		cursor := 0
		// Sweeps until the end of the word or when no instance was found
		for {
			idx := strings.Index(parsedLine[cursor:], d)

			if idx >= 0 {
				idxValue[cursor+idx] = parseStringToNumber(d)
				cursor += idx
			}

			if cursor >= len(line) || idx < 0 {
				break
			}

			cursor += len(d)
		}
	}

	numericString := ""
	for i := 0; i < len(line); i++ {
		if idxValue[i] != "" {
			numericString += idxValue[i]
		}
	}

	return numericString
}

func parseStringToNumber(value string) string {

	numbers := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	if len(value) > 1 {
		return numbers[value]
	}

	return value

}
