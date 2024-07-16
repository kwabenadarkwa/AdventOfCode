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
		if containsThreeVowels(word) && letterRepeats(word) && !containsAlphabets(word) {
			niceStrings++
		}
	}

	fmt.Println(niceStrings)
	fmt.Println(containsThreeVowels("aei"))

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
