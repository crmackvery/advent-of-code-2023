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
	file, err := os.Open("./input/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	resultP1 := 0
	resultP2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		resultP1 += getNumber(true, scanner.Text())
		resultP2 += getNumber(false, scanner.Text())
	}

	fmt.Println(resultP1)
	fmt.Println(resultP2)
}

func getNumber(isPart1 bool, line string) int {
	var allNumbers []string
	if isPart1 {
		allNumbers = getAllNumbers(line)
	} else {
		allNumbers = getAllP2Numbers(line)
	}

	if len(allNumbers) < 1 {
		return 0
	}

	i, err := strconv.Atoi(allNumbers[0])
	if err != nil {
        panic(err)
    }
	j, err := strconv.Atoi(allNumbers[len(allNumbers) - 1])
	if err != nil {
        panic(err)
    }

	return (i * 10) + j
}

func getAllNumbers(input string) []string {
	re := regexp.MustCompile("[0-9]")
	return re.FindAllString(input, -1)
}

func getAllP2Numbers(line string) []string {
	replacer := strings.NewReplacer("one", "o1e", "two", "t2o", "three", "t3e", "four",
		"f4r", "five", "f5e", "six", "s6x", "seven", "s7n", "eight", "e8t", "nine", "n9e")

	output := replacer.Replace(line)
	output = replacer.Replace(output)
	return getAllNumbers(output)
}
