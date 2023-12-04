package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type numberToIndices struct {
	value string
	validIndices []string
}

func main() {
	file, err := os.Open("./input/day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	resultP1 := 0
	resultP2 := 0
	scanner := bufio.NewScanner(file)
	row := 0
	validNumberIndicesSlice := []numberToIndices{}
	validSymbolIndicesSlice :=[]string{}
	validP2SymbolIndicesSlice :=[]string{}
	for scanner.Scan() {
		numberSlice, indexList, p2IndexList := findValidNumberAndSymbolIndices(row, scanner.Text())
		validNumberIndicesSlice = append(validNumberIndicesSlice, numberSlice...)
		validSymbolIndicesSlice = append(validSymbolIndicesSlice, indexList...)
		validP2SymbolIndicesSlice = append(validP2SymbolIndicesSlice, p2IndexList...)
		row++
	}

	validNumberIndicesSliceP2 := make([]numberToIndices, len(validNumberIndicesSlice))
	copy(validNumberIndicesSliceP2, validNumberIndicesSlice)
	resultP1 = calculateP1ValidNumberSum(validNumberIndicesSlice, validSymbolIndicesSlice)
	resultP2 = calculateP2ValidNumberSum(validNumberIndicesSliceP2, validP2SymbolIndicesSlice)

	fmt.Println(resultP1)
	fmt.Println(resultP2)
}

func findValidNumberAndSymbolIndices(row int, line string) ([]numberToIndices, []string, []string) {
	re := regexp.MustCompile("[0-9]+")
	numbers := re.FindAllStringIndex(line, -1)
	numbersSlice := []numberToIndices{}
	for _, number := range numbers {
		sIdx := number[0]
		indices := getSurroundingIndices(row, sIdx, number[1] - number[0], len(line))
		value := numberToIndices{
			value: line[number[0]:number[1]],
			validIndices: indices,
		}

		numbersSlice = append(numbersSlice, value)
	}

	re = regexp.MustCompile("[^.0-9]")
	symbols := re.FindAllStringIndex(line, -1)
	symbolPositions := []string{}
	for _, index := range symbols {
		symbolPositions = append(symbolPositions, strconv.Itoa(index[0]) + "-" + strconv.Itoa(row))
	}

	re = regexp.MustCompile("[*]")
	p2Symbols := re.FindAllStringIndex(line, -1)
	p2SymbolPositions := []string{}
	for _, index := range p2Symbols {
		p2SymbolPositions = append(p2SymbolPositions, strconv.Itoa(index[0]) + "-" + strconv.Itoa(row))
	}


	return numbersSlice, symbolPositions, p2SymbolPositions
}

func getSurroundingIndices(row int, col int, numberLength int, rowLength int) []string {
	ret := []string{}
	for x := col; x < col + numberLength; x++ {
		ret = append(ret, getIndiecesAroundAPixel(row, x)...)
	}

	return ret
}

func getIndiecesAroundAPixel(row int, col int) []string {
	ret := []string{}
	ret = append(ret, strconv.Itoa(col-1) + "-" + strconv.Itoa(row-1))
	ret = append(ret, strconv.Itoa(col-0) + "-" + strconv.Itoa(row-1))
	ret = append(ret, strconv.Itoa(col+1) + "-" + strconv.Itoa(row-1))
	ret = append(ret, strconv.Itoa(col-1) + "-" + strconv.Itoa(row-0))
	ret = append(ret, strconv.Itoa(col+1) + "-" + strconv.Itoa(row-0))
	ret = append(ret, strconv.Itoa(col-1) + "-" + strconv.Itoa(row+1))
	ret = append(ret, strconv.Itoa(col-0) + "-" + strconv.Itoa(row+1))	
	ret = append(ret, strconv.Itoa(col+1) + "-" + strconv.Itoa(row+1))	
	return ret
}

func calculateP1ValidNumberSum(numberToIndicesSlice []numberToIndices, symbolIndices []string) int {
	sum := 0
	for _, symbol := range symbolIndices {
		for idx, numberToIndices := range numberToIndicesSlice {
			if contains(numberToIndices.validIndices, symbol) {
				i, err := strconv.Atoi(numberToIndices.value)
				if err != nil {
					panic(err)
				}
				sum += i
				numberToIndicesSlice[idx].validIndices = []string{}
			}
		}
	}

	return sum
} 

func calculateP2ValidNumberSum(numberToIndicesSlice []numberToIndices, symbolIndices []string) int {
	sum := 0
	nMatches := 0
	firstMatchVal := 0
	for _, symbol := range symbolIndices {
		nMatches = 0
		firstMatchVal = 0
		for idx, numberToIndices := range numberToIndicesSlice {
			if contains(numberToIndices.validIndices, symbol) {
				nMatches++
				i, err := strconv.Atoi(numberToIndices.value)
				if err != nil {
					panic(err)
				}
				if nMatches == 2 {
					sum += (i * firstMatchVal)
				} else {
					firstMatchVal = i
				}
				numberToIndicesSlice[idx].validIndices = []string{}
			}
		}
	}

	return sum
} 

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
