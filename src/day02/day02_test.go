package main

import "testing"

func TestVerfiyGame1(t *testing.T) {
	// test each game
	game1 := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	game1Result := verifyGame(game1)
	// checking games
	if game1Result != 1 {
		t.Errorf("Game 1 verification failed")
	}
}

func TestVerfiyGame2(t *testing.T) {
	game2 := "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"
	game2Result := verifyGame(game2)
	if game2Result != 2 {
		t.Errorf("Game 2 verification failed")
	}
}

func TestVerfiyGame3(t *testing.T) {
	game3 := "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"
	game3Result := verifyGame(game3)
	if game3Result != 0 {
		t.Errorf("Game 3 verification failed")
	}
}

func TestVerfiyGame4(t *testing.T) {
	game4 := "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"
	game4Result := verifyGame(game4)
	if game4Result != 0 {
		t.Errorf("Game 4 verification failed")
	}
}

func TestVerfiyGame5(t *testing.T) {
	game5 := "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"
	game5Result := verifyGame(game5)
	if game5Result != 5 {
		t.Errorf("Game 5 verification failed")
	}
}

func TestVerifySum(t *testing.T) {
	games := [5]string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}

	sum := 0
	for _, game := range games {
		sum += verifyGame(game)
	}
	if sum != 8 {
		t.Errorf("Verification count failed. Got %d", sum)
	}
}

func TestVerifyMin1(t *testing.T) {
	games := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	gameResult := minimumGame(games)

	if fewestCubes["green"] != 2 {
		t.Errorf("Fewest green cubes is wrong. Should be 2, got %d", fewestCubes["green"])
	}
	if fewestCubes["red"] != 4 {
		t.Errorf("Fewest red cubes is wrong. Should be 4, got %d", fewestCubes["red"])
	}
	if fewestCubes["blue"] != 6 {
		t.Errorf("Fewest blue cubes is wrong. Should be 6, got %d", fewestCubes["blue"])
	}

	if gameResult != 48 {
		t.Errorf("Game 1 verifcation failed. Should be 48, got %d", gameResult)
	}
}
func TestVerifyMin2(t *testing.T) {
	games := "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"
	gameResult := minimumGame(games)

	if fewestCubes["green"] != 3 {
		t.Errorf("Fewest green cubes is wrong. Should be 3, got %d", fewestCubes["green"])
	}
	if fewestCubes["red"] != 1 {
		t.Errorf("Fewest red cubes is wrong. Should be 1, got %d", fewestCubes["red"])
	}
	if fewestCubes["blue"] != 4 {
		t.Errorf("Fewest blue cubes is wrong. Should be 4, got %d", fewestCubes["blue"])
	}

	if gameResult != 12 {
		t.Errorf("Game 2 verifcation failed. Should be 12, got %d", gameResult)
	}
}
func TestVerifyMin3(t *testing.T) {
	games := "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"
	gameResult := minimumGame(games)

	if fewestCubes["green"] != 13 {
		t.Errorf("Fewest green cubes is wrong. Should be 13, got %d", fewestCubes["green"])
	}
	if fewestCubes["red"] != 20 {
		t.Errorf("Fewest red cubes is wrong. Should be 20, got %d", fewestCubes["red"])
	}
	if fewestCubes["blue"] != 6 {
		t.Errorf("Fewest blue cubes is wrong. Should be 6, got %d", fewestCubes["blue"])
	}

	if gameResult != 1560 {
		t.Errorf("Game 3 verifcation failed. Should be 1560, got %d", gameResult)
	}
}
func TestVerifyMin4(t *testing.T) {
	games := "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"
	gameResult := minimumGame(games)

	if fewestCubes["green"] != 3 {
		t.Errorf("Fewest green cubes is wrong. Should be 3, got %d", fewestCubes["green"])
	}
	if fewestCubes["red"] != 14 {
		t.Errorf("Fewest red cubes is wrong. Should be 14, got %d", fewestCubes["red"])
	}
	if fewestCubes["blue"] != 15 {
		t.Errorf("Fewest blue cubes is wrong. Should be 15, got %d", fewestCubes["blue"])
	}

	if gameResult != 630 {
		t.Errorf("Game 4 verifcation failed. Should be 630, got %d", gameResult)
	}
}
func TestVerifyMin5(t *testing.T) {
	games := "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"
	gameResult := minimumGame(games)

	if fewestCubes["green"] != 3 {
		t.Errorf("Fewest green cubes is wrong. Should be 3, got %d", fewestCubes["green"])
	}
	if fewestCubes["red"] != 6 {
		t.Errorf("Fewest red cubes is wrong. Should be 6, got %d", fewestCubes["red"])
	}
	if fewestCubes["blue"] != 2 {
		t.Errorf("Fewest blue cubes is wrong. Should be 2, got %d", fewestCubes["blue"])
	}

	if gameResult != 36 {
		t.Errorf("Game 5 verifcation failed. Should be 36, got %d", gameResult)
	}
}
