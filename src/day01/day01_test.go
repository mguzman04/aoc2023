package main

import "testing"

func TestIsNumberZero(t *testing.T) {
	number := isNumber('0')
	if number != true {
		t.Errorf("isNumber(0) = %t; want true", number)
	}
}
func TestIsNumberNine(t *testing.T) {
	number := isNumber('9')
	if number != true {
		t.Errorf("isNumber(9) = %t; want true", number)
	}
}
func TestIsNumberFive(t *testing.T) {
	number := isNumber('5')
	if number != true {
		t.Errorf("isNumber(5) = %t; want true", number)
	}
}
func TestIsNumberLetter(t *testing.T) {
	notNumber := isNumber('A')
	if notNumber != false {
		t.Errorf("isNumber(A) = %t; want false", notNumber)
	}
}
func TestTwoDigit1(t *testing.T) {
	twelve := getTwoDigitNumber("1abc2")
	if twelve != 12 {
		t.Errorf("getTwoDigitNumber(1abc2) = %v; want 12", twelve)
	}
}

func TestTwoDigit2(t *testing.T) {
	thirtyEight := getTwoDigitNumber("pqr3stu8vwx")
	if thirtyEight != 38 {
		t.Errorf("getTwoDigitNumber(pqr3stu8vwx) = %v; want 38", thirtyEight)
	}
}
func TestTwoDigit3(t *testing.T) {
	fifteen := getTwoDigitNumber("a1b2c3d4e5f")
	if fifteen != 15 {
		t.Errorf("getTwoDigitNumber(a1b2c3d4e5f) = %v; want 15", fifteen)
	}
}
func TestTwoDigit4(t *testing.T) {
	seventySeven := getTwoDigitNumber("treb7uchet")
	if seventySeven != 77 {
		t.Errorf("getTwoDigitNumber(treb7uchet) = %v; want 77", seventySeven)
	}
}

func TestSum(t *testing.T) {
	input := [5]string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	sum := 0
	for _, text := range input {
		sum += getTwoDigitNumber(text)
	}
	if sum != 142 {
		t.Errorf("Example failed. Expected 142, got %v", sum)
	}
}

func TestRealDigits29(t *testing.T) {
	twoNine := getRealDigits("two1nine")
	if twoNine != 29 {
		t.Errorf("getRealDigits(two1nine) = %v; want 29", twoNine)
	}
}
func TestRealDigits83(t *testing.T) {
	eightThree := getRealDigits("eightwothree")
	if eightThree != 83 {
		t.Errorf("getRealDigits(eightwothree) = %v; want 83", eightThree)
	}
}
func TestRealDigits13(t *testing.T) {
	oneThree := getRealDigits("abcone2threexyz")
	if oneThree != 13 {
		t.Errorf("getRealDigits(abcone2threexyz) = %v; want 13", oneThree)
	}
}
func TestRealDigits24(t *testing.T) {
	twoFour := getRealDigits("xtwone3four")
	if twoFour != 24 {
		t.Errorf("getRealDigits(xtwone3four) = %v; want 24", twoFour)
	}
}
func TestRealDigits42(t *testing.T) {
	fourTwo := getRealDigits("4nineeightseven2")
	if fourTwo != 42 {
		t.Errorf("getRealDigits(4nineeightseven2) = %v; want 42", fourTwo)
	}
}
func TestRealDigits14(t *testing.T) {
	oneFour := getRealDigits("zoneight234")
	if oneFour != 14 {
		t.Errorf("getRealDigits(zoneight234) = %v; want 14", oneFour)
	}
}
func TestRealDigits76(t *testing.T) {
	sevenSix := getRealDigits("7pqrstsixteen")
	if sevenSix != 76 {
		t.Errorf("getRealDigits(7pqrstsixteen) = %v; want 76", sevenSix)
	}
}

func TestRealSum(t *testing.T) {
	input := [7]string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}
	sum := 0
	for _, text := range input {
		sum += getRealDigits(text)
	}
	if sum != 281 {
		t.Errorf("Example failed. Expected 281, got %v", sum)
	}
}
