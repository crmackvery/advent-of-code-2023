package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input/day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	resultP1 := 0
	resultP2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		p1, p2 := getGameNumberIfValid(scanner.Text())
		resultP1 += p1
		resultP2 += p2
	}

	fmt.Println(resultP1)
	fmt.Println(resultP2)
}

func getGameNumberIfValid(line string) (int, int) {
	gameAndRest := strings.Split(line, ":")
	gameStr := strings.Split(gameAndRest[0], " ")[1]
	game, err := strconv.Atoi(gameStr)
	if err != nil {
        panic(err)
    }
	pulls := strings.Split(gameAndRest[1], ";")
	highCountMap := make(map[string]int)
	highCountMap["red"] = 0;
	highCountMap["green"] = 0;
	highCountMap["blue"] = 0;
	
	for _, pull := range pulls {
		nRed := getCountByColor(pull, "red")
		nGreen := getCountByColor(pull, "green")
		nBlue := getCountByColor(pull, "blue")
		if nRed > 12 || nGreen > 13 || nBlue > 14 {
			game = 0
		}
		if nRed > highCountMap["red"] {
			highCountMap["red"] = nRed
		}
		if nGreen > highCountMap["green"] {
			highCountMap["green"] = nGreen
		}
		if nBlue > highCountMap["blue"] {
			highCountMap["blue"] = nBlue
		}
	}

	pow := highCountMap["red"] * highCountMap["blue"] * highCountMap["green"]

	return game, pow
}

func getCountByColor(pull string, color string) int {
	re := regexp.MustCompile("[0-9]+ " + color)
	colorResults := re.FindAllString(pull, -1)
	if len(colorResults) < 1 {
		return 0
	}

	re = regexp.MustCompile("[0-9]+")
	number := re.FindAllString(colorResults[0], -1)
	i, err := strconv.Atoi(number[0])
	if err != nil {
        panic(err)
    }

	return i
}
