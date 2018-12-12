package main

import (
	"bufio"
	"fmt"
	"os"
)

type fabric struct {
	claims map[claim]int
}

func (f *fabric) add(x, y, w, h int) {
	if f.claims == nil {
		f.claims = make(map[claim]int)
	}

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			f.claims[claim{x + i, y + j}]++
		}
	}
}

type claim struct {
	x int
	y int
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	fabric := fabric{}
	for s.Scan() {
		var id, x, y, w, h int
		fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		fabric.add(x, y, w, h)
	}
	var total int
	for _, c := range fabric.claims {
		if c > 1 {
			total++
		}
	}

	fmt.Printf("The answer is %d.", total)
}
