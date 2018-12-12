package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var freqs []int
	var total int
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var freq int
		fmt.Sscanf(s.Text(), "%d", &freq)
		freqs = append(freqs, freq)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	freqSet := map[int]bool{0: true}

	for {
		for _, f := range freqs {
			total += f
			if _, exists := freqSet[total]; exists {
				fmt.Fprintf(os.Stdout, "The answer is %d", total)
				os.Exit(0)
			} else {
				freqSet[total] = true
			}
		}
	}
}
