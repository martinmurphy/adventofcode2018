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
	nonoverlapping := checkNonOverlap(string(data))
	fmt.Printf("Non-overlap: %v\n", nonoverlapping)

}
