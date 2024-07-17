package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

	scanner := bufio.NewScanner(file)
	var niceStrings int
	var word string

	for scanner.Scan() {
		word = scanner.Text()
		// _ = word

		if containsPairOfTwoLetters(word) && letterInBetween(word) {
			niceStrings++
		}
	}
	// fmt.Println(letterInBetween("xyxy"))

	fmt.Println(niceStrings)
	// if containsPairOfTwoLetters("ieodomkazucvgmuy") && letterInBetween("ieodomkazucvgmuy") {
	//    fmt.Println("yes")
	// }else{
	//    fmt.Println("no")
	//  }

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func containsThreeVowels(word string) bool {
	var count int
	for _, alpha := range word {
		switch alpha {
		case 'a':
			count++
		case 'e':
			count++
		case 'i':
			count++
		case 'o':
			count++
		case 'u':
			count++
		}
	}
	if count >= 3 {
		return true
	}
	return false
}

func letterRepeats(word string) bool {
	for index := range word {
		if index == 0 {
			continue
		} else if word[index] == word[index-1] {
			return true
		}
	}
	return false
}

func containsAlphabets(word string) bool {
	var twoWords string

	for index := range word {
		if index == 0 {
			continue
		} else {
			twoWords = string(word[index-1]) + string(word[index])
			switch twoWords {
			case "ab":
				return true
			case "cd":
				return true
			case "pq":
				return true
			case "xy":
				return true
			}
		}
	}
	return false
}

// it contains a pair of any two letters that appears at least twice in the string without
// overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like aaa (aa, but it overlaps).
func containsPairOfTwoLetters(word string) bool {
	var pairs [8]string
	for l := 0; l < len(word)/2; l++ {
		pairs[l] = string(word[2*l]) + string(word[2*l+1])
	}

	for index, pair := range pairs {
		for compIndex, comPair := range pairs {
			if index != compIndex && pair == comPair {
				return true
			}
		}
	}
	return false
}

func letterInBetween(word string) bool {
	lenWord := len(word)
	for index := range word {
		if index < lenWord-2 && word[index] == word[index+2] {
      // fmt.Println(word[index],word[index+2])
			return true
		}
	}
	return false
}
