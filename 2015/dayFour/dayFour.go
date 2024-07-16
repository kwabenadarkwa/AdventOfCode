// Day 4: The ideal Stocking Suffer
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	secret := "yzbqklnj"
	var numToAppend int

	for {
		if hashCorrect(md5Hash(secret + strconv.Itoa(numToAppend))) {
			break
		} else {
			numToAppend++
			continue
		}
	}
	fmt.Println(numToAppend)
}

func md5Hash(secret string) string {
	hash := md5.Sum([]byte(secret))
	return hex.EncodeToString(hash[:])
}

func hashCorrect(hash string) bool {
	if hash[:6] == "000000" {
		return true
	}
	return false
}
