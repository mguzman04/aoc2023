package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	sum := 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		input := scanner.Text()
		sum += getTwoDigitNumber(input)
	}
	fmt.Printf("Sum result: %v ", sum)

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
