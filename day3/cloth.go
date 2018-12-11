package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Claim struct {
	id                       string
	left, top, width, height int
}

func parseClaim(input string) Claim {

	//"#123 @ 3,2: 5x4"
	re := regexp.MustCompile(`\#\s*([0-9]+)\s*\@\s*([0-9]+)\s*,\s*([0-9]+)\s*:\s*([0-9]+)\s*x\s*([0-9]+)\s*`)

	nums := re.FindAllStringSubmatch(input, -1)

	left, _ := strconv.Atoi(nums[0][2])
	top, _ := strconv.Atoi(nums[0][3])
	width, _ := strconv.Atoi(nums[0][4])
	height, _ := strconv.Atoi(nums[0][5])

	c := Claim{nums[0][1], left, top, width, height}
	return c
}

func countOvers(c [][]rune, width, height int) int {
	count := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if c[y][x] == '#' {
				count++
			}
		}
	}
	return count
}

func addClaimFile(c [][]rune, claims []string) {
	for _, v := range claims {
		if len(v) == 0 {
			fmt.Printf("skipping %v\n", v)
			continue // skip empty lines
		}
		claim := parseClaim(v)
		addClaim(c, claim)
	}
}
func overlapClaim(cloth [][]rune, c Claim) bool {
	for y := c.top; y < (c.top + c.height); y++ {
		for x := c.left; x < (c.left + c.width); x++ {
			if cloth[y][x] == '#' {
				return true
			}
		}
	}
	return false
}

func checkClaimFile(c [][]rune, claims []string) Claim {
	for _, v := range claims {
		if len(v) == 0 {
			fmt.Printf("skipping %v\n", v)
			continue // skip empty lines
		}
		claim := parseClaim(v)
		if !overlapClaim(c, claim) {
			return claim
		}
	}
	return Claim{"NOID", 0, 0, 0, 0}
}

func countOverlaps(inputFile string) int {
	c := MakeCloth(1000, 1000)
	claims := strings.Split(string(inputFile), "\n")

	addClaimFile(c, claims)
	return countOvers(c, 1000, 1000)
}

func checkNonOverlap(inputFile string) string {
	c := MakeCloth(1000, 1000)
	claims := strings.Split(string(inputFile), "\n")

	addClaimFile(c, claims)
	claim := checkClaimFile(c, claims)
	return claim.id
}

func addClaim(cloth [][]rune, c Claim) {
	for y := c.top; y < (c.top + c.height); y++ {
		for x := c.left; x < (c.left + c.width); x++ {
			if cloth[y][x] == '.' {
				cloth[y][x] = 'I'
			} else {
				cloth[y][x] = '#'
			}
		}
	}
	return
}

func MakeCloth(wid, hei int) [][]rune {
	var cloth [][]rune
	cloth = make([][]rune, hei)

	for y := 0; y < hei; y++ {
		cloth[y] = make([]rune, wid)
		for x := 0; x < wid; x++ {
			cloth[y][x] = '.'
		}
	}
	return cloth
}
