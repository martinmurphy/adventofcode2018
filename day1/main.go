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
	memo := make(map[int64]bool)
	memo[freq] = true

	for {
		for i, v := range nums {
			if len(v) == 0 {
				fmt.Printf("skipping %v %v\n", i, v)
				continue // skip empty lines
			}
			val, _ := strconv.ParseInt(v, 10, 64)
			freq = freq + val
			if memo[freq] {
				fmt.Printf("already seen %v\n", freq)
				return
			}
			memo[freq] = true
		}
	}

}
