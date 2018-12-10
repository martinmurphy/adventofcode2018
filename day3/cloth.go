package main

type Claim struct {
	id                       string
	left, top, width, height int
}

func parseClaim(input string) Claim {
	c := Claim{"NOID", 0, 0, 0, 0}
	return c
}

func countOverlaps(inputFile string) int {
	return 0
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
