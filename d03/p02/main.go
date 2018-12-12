package main

import (
	"bufio"
	"fmt"
	"os"
)

type fabric struct {
	goodIds map[int]bool
	claims  map[claim][]int
}

func (f *fabric) add(id, x, y, w, h int) {
	if f.claims == nil {
		f.claims = make(map[claim][]int)
		f.goodIds = make(map[int]bool)
	}

	overlapped := false
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			f.claims[claim{x + i, y + j}] = append(f.claims[claim{x + i, y + j}], id)
			if len(f.claims[claim{x + i, y + j}]) > 1 {
				overlapped = true
				for _, id := range f.claims[claim{x + i, y + j}] {
					delete(f.goodIds, id)
				}
			}
		}
	}

	if !overlapped {
		f.goodIds[id] = true
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
		fabric.add(id, x, y, w, h)
	}

	for id := range fabric.goodIds {
		fmt.Fprintf(os.Stdout, "The answer is %d\n", id)
		os.Exit(0)
	}
	panic("oh no!")
}
