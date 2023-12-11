package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	count := 0
	var matrix []string
	for scanner.Scan() {
		input := scanner.Text()
		count++
		split := strings.Split(input, "")
		matrix = append(matrix, split...)
	}
	fmt.Println(len(matrix))
	fmt.Println(count)
	file.Close()
}
