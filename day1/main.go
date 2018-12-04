package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	var freq int64
	freq = 0

	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	nums := strings.Split(string(dat), "\n")
	for _, v := range nums {
		i, _ := strconv.ParseInt(v, 10, 64)
		freq = freq + i
	}

	fmt.Printf("freq: %v", freq)
}
