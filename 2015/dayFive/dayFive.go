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

		if containsPairOfTwoLetters(word) && letterInBetween(word) {
			niceStrings++
		}
	}

	fmt.Println(niceStrings)

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

func containsPairOfTwoLetters(word string) bool {
	var pairs []string

	for index := range word {
		if index <= len(word)-2 {
			pairs = append(pairs, string(word[index])+string(word[index+1]))
		}
	}

	for index, pair := range pairs {
		for compIndex, comPair := range pairs {
			if index != compIndex && index+1 != compIndex && index-1 != compIndex &&
				pair == comPair {
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
			return true
		}
	}
	return false
}
