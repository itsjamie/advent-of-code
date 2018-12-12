package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var totalTwos, totalThrees int

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var id string
		fmt.Sscanf(s.Text(), "%s", &id)
		twos, threes := checksum(id)
		if twos {
			totalTwos++
		}
		if threes {
			totalThrees++
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(os.Stdout, "The answer is %d", totalTwos*totalThrees)
}

func checksum(id string) (hasTwo bool, hasThree bool) {
	lOccurs := map[rune]int{}

	for _, c := range id {
		lOccurs[c]++
	}

	for _, v := range lOccurs {
		if v == 2 {
			hasTwo = true
		} else if v == 3 {
			hasThree = true
		}
	}

	return
}
