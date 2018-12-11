package main

import "testing"

var testString = "#123 @ 3,2: 5x4"
var expectedClaim = Claim{"123", 3, 2, 5, 4}

var testClaim = Claim{"1", 1, 2, 3, 2}

var testInputFile = `#123 @ 3,2: 5x4
#123 @ 3,2: 5x4
`

func TestParseClaim(t *testing.T) {
	actualClaim := parseClaim(testString)

	if actualClaim.id != expectedClaim.id {
		t.Errorf("error parsing id, expected %v, actual %v", expectedClaim.id, actualClaim.id)
	}
	if actualClaim.left != expectedClaim.left {
		t.Errorf("error parsing left, expected %v, actual %v", expectedClaim.left, actualClaim.left)
	}
	if actualClaim.top != expectedClaim.top {
		t.Errorf("error parsing top, expected %v, actual %v", expectedClaim.top, actualClaim.top)
	}
	if actualClaim.width != expectedClaim.width {
		t.Errorf("error parsing width, expected %v, actual %v", expectedClaim.width, actualClaim.width)
	}
	if actualClaim.height != expectedClaim.height {
		t.Errorf("error parsing height, expected %v, actual %v", expectedClaim.height, actualClaim.height)
	}
}

func TestClaim(t *testing.T) {
	c := MakeCloth(10, 10)
	addClaim(c, testClaim)
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if x >= 1 && x <= 3 && y >= 2 && y <= 3 {
				if c[y][x] != 'I' {
					t.Errorf("error claimed area incorrect at x: %v, y %v", x, y)
				}
			} else {
				if c[y][x] != '.' {
					t.Errorf("error unclaimed area incorrect at x: %v, y %v", x, y)
				}
			}
		}
	}
}

func TestOverlapClaim(t *testing.T) {
	c := MakeCloth(10, 10)
	addClaim(c, testClaim)
	addClaim(c, testClaim)
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if x >= 1 && x <= 3 && y >= 2 && y <= 3 {
				if c[y][x] != '#' {
					t.Errorf("error overlap area incorrect at x: %v, y %v", x, y)
				}
			} else {
				if c[y][x] != '.' {
					t.Errorf("error unclaimed area incorrect at x: %v, y %v", x, y)
				}
			}
		}
	}
}
func TestOverlaps(t *testing.T) {
	c := MakeCloth(10, 10)
	addClaim(c, testClaim)
	addClaim(c, testClaim)
	num := countOvers(c, 10, 10)
	if num != 6 {
		t.Errorf("error overlap area incorrect expected: %v,actual %v", 6, num)
	}
}

func TestCountOverlaps(t *testing.T) {

	c := countOverlaps(testInputFile)
	if c != 20 {
		t.Errorf("error parsing inputfile, expected %v, actual %v", 20, c)
	}

}
