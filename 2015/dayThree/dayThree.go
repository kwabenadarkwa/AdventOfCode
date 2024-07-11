// Day 3: Perfectly spherical houses in a vacuum
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type point struct {
	x            int
	y            int
	timesVisited int
	isCurrentPos bool
}

type coordinateExistsReturn struct {
	cExists bool
	index   int
}

func main() {
	var coordinates []point
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	coordinates = append(coordinates, point{0, 0, 1, true})

	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {

      position := currentPosition(coordinates)
			recentVisit := coordinates[position]
			x := recentVisit.x
			y := recentVisit.y

			coordinates[position].isCurrentPos = false

			switch char {
			case '^':
				y += 1
			case '>':
				x += 1
			case '<':
				x -= 1
			case 'v':
				y -= 1
			}

			coordinateDoesExist := coordinateExists(coordinates, x, y)

			if coordinateDoesExist.cExists == true {
				cLocation := coordinateDoesExist.index
				coordinates[cLocation].timesVisited += 1
				coordinates[cLocation].isCurrentPos = true
			} else {
				coordinates = append(coordinates, point{x, y, 1, true})
			}

		}
	}

	fmt.Println(len(coordinates))
	// fmt.Println(coordinates)
	// count := 0
	// for _, point := range coordinates{
	//   if point.timesVisited > 1{
	//     count += 1
	//   }
	// }
	// fmt.Println(count)
	// check how many
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// var coordinates []point
	// coordinates = append(coordinates, point{0, 0, 1, true})
	// coordinates = append(coordinates, point{0, 1, 1, false})
	// coordinates = append(coordinates, point{0, 2, 1, false})
	//  fmt.Println(coordinates)
	//  fmt.Println(currentPosition(coordinates))
	//  coordinates[currentPosition(coordinates)].isCurrentPos = false
	//  fmt.Println(coordinates)
}

func coordinateExists(coordinates []point, x int, y int) coordinateExistsReturn {
	returnType := coordinateExistsReturn{false, 0}
	for index, point := range coordinates {
		if point.x == x && point.y == y {
			returnType = coordinateExistsReturn{true, index}
			return returnType
		}
	}
	return returnType
}

func currentPosition(coordinates []point) int {
	for index, point := range coordinates {
		if point.isCurrentPos == true {
			return index
		}
	}
	return 0
}
