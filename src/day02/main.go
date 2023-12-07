package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var totalCubes map[string]int

func init() {
	totalCubes = make(map[string]int)
	totalCubes["red"] = 12
	totalCubes["green"] = 13
	totalCubes["blue"] = 14
}

func main() {
	sum := 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input := scanner.Text()

		// part 1
		sum += verifyGame(input)

	}
	fmt.Printf("Part 1 result: %v\n ", sum)

	file.Close()
}

func verifyGame(input string) int {
	record := strings.Split(input, ":")
	recordID := record[0]
	recordGames := record[1]

	recordIDSplit := strings.Split(recordID, " ")
	id, _ := strconv.Atoi(recordIDSplit[1])
	sets := strings.Split(recordGames, ";")
	for _, set := range sets {
		cubes := strings.Split(set, ",")
		for _, cube := range cubes {
			content := strings.Split(cube, " ")
			count, _ := strconv.Atoi(content[1])
			if totalCubes[content[2]] < count {
				return 0
			}
		}
	}
	return id
}
