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
