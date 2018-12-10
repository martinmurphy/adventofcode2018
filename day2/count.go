package main

import (
	"fmt"
	"strings"
)

func countID(input string) (twos, threes int) {
	// fmt.Printf("Counting %v\n", input)
	freq := make(map[rune]int)
	for _, v := range input {
		freq[v] = freq[v] + 1
	}
	for _, v := range freq {
		if v == 2 {
			twos = 1
		}
		if v == 3 {
			threes = 1
		}

	}
	// fmt.Printf("freqs %v\n", freq)
	return twos, threes
}

func checksum(inputData string) (checksum int) {
	twos, threes := 0, 0

	// fmt.Printf("Checksumming %v\n", inputData)
	ids := strings.Split(inputData, "\n")
	for _, v := range ids {
		if len(v) == 0 {
			continue // skip empty lines
		}
		rowTwos, rowThrees := countID(v)
		twos = twos + rowTwos
		threes = threes + rowThrees
	}
	return twos * threes
}

func differences(id1, id2 string) int {
	runeID1 := []rune(id1)
	runeID2 := []rune(id2)
	diffs := 0
	for i := 0; i < len(runeID1); i++ {
		if runeID1[i] != runeID2[i] {
			diffs++
		}
	}
	return diffs
}

func sameChars(id1, id2 string) string {
	ret := ""
	runeID1 := []rune(id1)
	runeID2 := []rune(id2)
	for i := 0; i < len(runeID1); i++ {
		if runeID1[i] == runeID2[i] {
			ret = ret + string(runeID1[i])
		}
	}
	return ret
}

func findIDs(inputData string) string {
	ret := ""
	ids := strings.Split(string(inputData), "\n")
	for i, v := range ids {
		if len(v) == 0 {
			fmt.Printf("skipping %v %v\n", i, v)
			continue // skip empty lines
		}
		for j := 0; j < i; j++ {
			if differences(v, ids[j]) == 1 {
				return sameChars(v, ids[j])
			}
		}
	}

	return ret
}
