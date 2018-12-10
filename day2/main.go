package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	checksum := checksum(string(data))
	fmt.Printf("Checksum of input file: %v\n", checksum)
	sameChars := findIDs(string(data))
	fmt.Printf("samechars of wanted ids: %v\n", sameChars)
}
