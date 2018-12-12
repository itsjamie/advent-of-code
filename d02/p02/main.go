package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var ids []string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var id string
		fmt.Sscanf(s.Text(), "%s", &id)
		ids = append(ids, id)
	}

	for i := 0; i < len(ids); i++ {
		for j := i; j < len(ids); j++ {
			c := common(ids[i], ids[j])
			if len(c) == len(ids[i])-1 {
				fmt.Fprintf(os.Stdout, "The answer is %s", c)
				os.Exit(0)
			}
		}
	}
}

// we know strings are same length
func common(a, b string) string {
	var sb strings.Builder
	for idx := range a {
		if a[idx] == b[idx] {
			sb.WriteByte(a[idx])
		}
	}
	return sb.String()
}
