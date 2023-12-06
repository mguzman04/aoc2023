package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	hi := "Hello!"
	slice := []rune(hi)
	fmt.Println(slice)
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
