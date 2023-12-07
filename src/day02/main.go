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
var fewestCubes map[string]int

func init() {
	totalCubes = make(map[string]int)
	totalCubes["red"] = 12
	totalCubes["green"] = 13
	totalCubes["blue"] = 14

	fewestCubes = make(map[string]int)
	fewestCubes["red"] = 0
	fewestCubes["green"] = 0
	fewestCubes["blue"] = 0
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
		sum += verifyGame(input)

		// part 2
		sum2 += minimumGame(input)
	}
	fmt.Printf("Part 1 result: %v\n ", sum)

	fmt.Printf("Part 2 result: %v\n ", sum2)

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

func minimumGame(input string) int {
	record := strings.Split(input, ":")
	recordGames := record[1]

	sets := strings.Split(recordGames, ";")
	for _, set := range sets {
		cubes := strings.Split(set, ",")
		for _, cube := range cubes {
			category := strings.Split(cube, " ")
			count, _ := strconv.Atoi(category[1])
			color := category[2]
			if fewestCubes[color] < count {
				fewestCubes[color] = count
			}
		}
	}

	return fewestCubes["green"] * fewestCubes["red"] * fewestCubes["blue"]
}
