package main

import "testing"

// abcdef contains no letters that appear exactly two or three times.
// bababc contains two a and three b, so it counts for both.
// abbcde contains two b, but no letter appears exactly three times.
// abcccd contains three c, but no letter appears exactly two times.
// aabcdd contains two a and two d, but it only counts once.
// abcdee contains two e.
// ababab contains three a and three b, but it only counts once.

// str, twos, threes
var countTests = []struct {
	str            string
	expectedTwos   int
	expectedThrees int
}{
	{"abcdef", 0, 0},
	{"bababc", 1, 1},
	{"abbcde", 1, 0},
	{"abcccd", 0, 1},
	{"aabcdd", 1, 0},
	{"abcdee", 1, 0},
	{"ababab", 0, 1},
}
var expectedChecksum = 12

var differenceData = []struct {
	id1, id2, sameChars string
	expectedDiffs       int
}{
	{"abcde", "axcye", "ace", 2},
	{"fghij", "fguij", "fgij", 1},
	{"abcde", "abcde", "abcde", 0},
}

var inputIDs = `abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz`
var expectedFoundIDs = "fgij"

func TestCounts(t *testing.T) {
	for i, data := range countTests {
		twos, threes := countID(data.str)
		if twos != data.expectedTwos || threes != data.expectedThrees {
			t.Errorf("error on row %v (%s), expected (%v, %v), actual (%v, %v)", i, data.str, data.expectedTwos, data.expectedThrees, twos, threes)
		}
	}
}

func TestChecksum(t *testing.T) {
	testData := ""
	for _, data := range countTests {
		testData += data.str + "\n"
	}
	actualChecksum := checksum(testData)
	if actualChecksum != expectedChecksum {
		t.Errorf("error on checksum, expected %v, actual %v", expectedChecksum, actualChecksum)
	}
}

func TestDifferences(t *testing.T) {
	for i, data := range differenceData {
		actualDiffs := differences(data.id1, data.id2)
		if actualDiffs != data.expectedDiffs {
			t.Errorf("error on row %v (%s, %s), expected %v, actual %v", i, data.id1, data.id2, data.expectedDiffs, actualDiffs)
		}
	}

}

func TestSameChars(t *testing.T) {
	for i, data := range differenceData {
		actualSames := sameChars(data.id1, data.id2)
		if actualSames != data.sameChars {
			t.Errorf("error on row %v (%s, %s), expected %v, actual %v", i, data.id1, data.id2, data.sameChars, actualSames)
		}
	}

}

func TestFindIDs(t *testing.T) {
	actualFoundIDs := findIDs(inputIDs)
	if actualFoundIDs != expectedFoundIDs {
		t.Errorf("error, expected %v, actual %v", expectedFoundIDs, actualFoundIDs)
	}

}
