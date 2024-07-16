// Day 3: Perfectly spherical houses in a vacuum
// this keeps track of how many times each house is visited
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// for part 2 :
// santa starts first with the delivering but takes turns with robo santa
// so santa moves and then robo santa moves and so on and so forth
// find a way to track Santa's deliveries and then robosantas
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
	var roboCoordinates []point

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
	roboCoordinates = append(roboCoordinates, point{0, 0, 1, true})

	for scanner.Scan() {
		line := scanner.Text()
		for index, char := range line {

			position := 0
			var recentVisit point
			if index%2 == 0 {
				position = currentPosition(coordinates)
				recentVisit = coordinates[position]
				coordinates[position].isCurrentPos = false
			} else {

				position = currentPosition(roboCoordinates)
				recentVisit = roboCoordinates[position]
				roboCoordinates[position].isCurrentPos = false
			}

			x := recentVisit.x
			y := recentVisit.y

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

			var coordinateDoesExist coordinateExistsReturn
			if index%2 == 0 {
				coordinateDoesExist = coordinateExists(coordinates, x, y)
			} else {
				coordinateDoesExist = coordinateExists(roboCoordinates, x, y)
			}

			if coordinateDoesExist.cExists == true {
				cLocation := coordinateDoesExist.index

				if index%2 == 0 {
					coordinates[cLocation].timesVisited += 1
					coordinates[cLocation].isCurrentPos = true
				} else {

					roboCoordinates[cLocation].timesVisited += 1
					roboCoordinates[cLocation].isCurrentPos = true
				}
			} else {
				if index%2 == 0 {
					coordinates = append(coordinates, point{x, y, 1, true})
				} else {
					roboCoordinates = append(roboCoordinates, point{x, y, 1, true})
				}
			}

		}
	}

	fmt.Println(roboCoordinates)
	fmt.Println(coordinates)
	fmt.Println(len(coordinates) + countIfNotInCoordinates(coordinates, roboCoordinates))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
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

func countIfNotInCoordinates(coordinates []point, roboCoordinates []point) int {
	var count int
	for _, point := range roboCoordinates {
		if coordinateExists(coordinates, point.x, point.y).cExists {
			continue
		} else {
			count++
		}
	}
	return count
}
