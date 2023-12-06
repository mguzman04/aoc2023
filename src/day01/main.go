package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var numbers map[string]string

func init() {
	numbers = make(map[string]string)
	// edge cases
	numbers["oneight"] = "18"
	numbers["threeight"] = "38"
	numbers["fiveight"] = "58"
	numbers["nineight"] = "98"
	numbers["eightwo"] = "82"
	numbers["sevenine"] = "79"
	numbers["twone"] = "21"
	// numbers
	numbers["one"] = "1"
	numbers["two"] = "2"
	numbers["three"] = "3"
	numbers["four"] = "4"
	numbers["five"] = "5"
	numbers["six"] = "6"
	numbers["seven"] = "7"
	numbers["eight"] = "8"
	numbers["nine"] = "9"
}

func main() {
	sum := 0
	sum2 := 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input := scanner.Text()

		// part 1
		sum += getTwoDigitNumber(input)

		// part 2
		sum2 += getRealDigits(input)
	}
	fmt.Printf("Part 1 result: %v\n ", sum)
	fmt.Printf("Part 2 result: %v\n ", sum2)

	file.Close()
}

func getTwoDigitNumber(input string) int {
	var firstDigit rune
	var lastDigit rune
	runes := []rune(input)
	for _, char := range runes {
		if isNumber(char) {
			if firstDigit == 0 {
				firstDigit = char
			}
			lastDigit = char
		}
	}
	result, _ := strconv.Atoi(string(firstDigit) + string(lastDigit))
	return result
}

func isNumber(char rune) bool {
	return unicode.IsDigit((char))
}

func findAndReplace(input string) string {
	newString := input
	for key, value := range numbers {
		newString = strings.ReplaceAll(newString, key, value)
	}
	return newString
}

func getRealDigits(input string) int {
	return getTwoDigitNumber(findAndReplace((input)))
}
