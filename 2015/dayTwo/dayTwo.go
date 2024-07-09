// Day 2: I was told there would be no math
// the data is gotten in the form l*w*h
// surface area 2*l*w + 2*w*h + 2*h*l
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type leastsReturn [2]int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var totalSize int
	var totalRibbonSize int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		list := strings.SplitN(line, "x", 3)

		l, _ := strconv.Atoi(list[0])
		w, _ := strconv.Atoi(list[1])
		h, _ := strconv.Atoi(list[2])

		intList := [3]int{l, w, h}
		leastForRibbon := leastValueWithSecondLeast(intList)
		leastIndex := leastForRibbon[0]
		secondLeast := leastForRibbon[1]

		surfaceArea := 2*l*w + 2*w*h + 2*h*l
		ribbonForBox := 2*(intList[leastIndex]+intList[secondLeast]) + (l * w * h)
		arrayOfMult := [3]int{l * w, w * h, h * l}
		leastPosition := leastValueWithSecondLeast(arrayOfMult)[0]

		switch leastPosition {
		case 0:
			surfaceArea += arrayOfMult[0]
		case 1:
			surfaceArea += arrayOfMult[1]
		case 2:
			surfaceArea += arrayOfMult[2]
		}

		totalSize += surfaceArea
		totalRibbonSize += ribbonForBox
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(totalSize, totalRibbonSize)
}

func leastValueWithSecondLeast(value [3]int) leastsReturn {
	var (
		indexForMin  int
		prevMinIndex int
	)
	min := math.MaxInt64

	for index, number := range value {
		if number < min {
			min = number
			indexForMin = index
		}
	}

	min = math.MaxInt64

	for index, number := range value {
		if index == indexForMin {
			continue
		} else {
			if number < min {
				min = number
				prevMinIndex = index
			}
		}
	}
	ret := leastsReturn{indexForMin, prevMinIndex}
	return ret
}
