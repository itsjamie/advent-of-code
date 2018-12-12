package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	total := 0
	var freq int
	for s.Scan() {
		fmt.Sscanf(s.Text(), "%d", &freq)
		total += freq
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(os.Stdout, "The answer is %d", total)
}
