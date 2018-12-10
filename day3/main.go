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

	countOverlaps := countOverlaps(string(data))
	fmt.Printf("Overlaps: %v\n", countOverlaps)
}
