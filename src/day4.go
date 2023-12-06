package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


func main() {
	file, err := os.Open("./input/day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	resultP1 := 0
	resultP2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		resultP1 += getP1ResultForRow(scanner.Text())
	}

	fmt.Println(resultP1)
	fmt.Println(resultP2)
}

func getP1ResultForRow(line string) int {
	tmpParts := strings.Split(line, ":")
	sets := strings.Split(tmpParts[1], "|")

	getWinnersRight := strings.Trim(strings.Replace(sets[0], "  ", " ", -1), " ")
	getNumbersRight := strings.Trim(strings.Replace(sets[1], "  ", " ", -1), " ")

	winnersStr := strings.Split(getWinnersRight, " ")
	numbersStr := strings.Split(getNumbersRight, " ")
	winners := make([]int, len(winnersStr))
	numbers := make([]int, len(numbersStr))
	for i, v := range winnersStr {
		val, err := strconv.Atoi(strings.Trim(v, " "))
		if err != nil {
			panic(err)
		}
		winners[i] = val
	}
	for i, v := range numbersStr {
		val, err := strconv.Atoi(strings.Trim(v, " "))
		if err != nil {
			panic(err)
		}
		numbers[i] = val
	}

	intersection := getIntersection(winners, numbers)
	
	result := 0
	if len(intersection) > 0 {
		result = 1
		for i, _ := range intersection {
			if i > 0 {
				result *= 2
			}
		}
	}

	return result
}

func getIntersection(a []int, b []int) []int {
	m := make(map[int]uint8)
    for _, k := range a {
        m[k] |= (1 << 0)
    }
    for _, k := range b {
        m[k] |= (1 << 1)
    }

    var inAAndB, inAButNotB, inBButNotA []int
    for k, v := range m {
        a := v&(1<<0) != 0
        b := v&(1<<1) != 0
        switch {
        case a && b:
            inAAndB = append(inAAndB, k)
        case a && !b:
            inAButNotB = append(inAButNotB, k)
        case !a && b:
            inBButNotA = append(inBButNotA, k)
        }
    }

	return inAAndB
}