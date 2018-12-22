package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type input struct {
	t    time.Time
	data string
}

type guard struct {
	id           int
	asleepAt     int
	sleepingTime int
}

func main() {
	inputs := []input{}
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		str := s.Text()

		t, err := time.Parse("[2006-01-02 15:04]", str[0:18])
		if err != nil {
			log.Fatal(err)
		}

		inputs = append(inputs, input{
			t:    t,
			data: str[19:],
		})
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Slice(inputs, func(i, j int) bool {
		return inputs[i].t.Before(inputs[j].t)
	})

	var onduty *guard
	guards := make(map[int]*guard, 0)
	for _, input := range inputs {
		switch input.data {
		case "falls asleep":
			onduty.asleepAt = input.t.Minute()
		case "wakes up":
			onduty.sleepingTime += input.t.Minute() - onduty.asleepAt
			onduty.asleepAt = 0
		default:
			var id int
			if _, err := fmt.Sscanf(input.data, "Guard #%d begins shift", &id); err != nil {
				log.Fatal(err)
			}

			if g, exists := guards[id]; !exists {
				g = &guard{
					id: id,
				}
				guards[id] = g
				onduty = g
			} else {
				onduty = g
			}

		}
	}
	var highestSleepGuard *guard
	for _, g := range guards {
		if highestSleepGuard == nil {
			highestSleepGuard = g
		}
		if g.sleepingTime > highestSleepGuard.sleepingTime {
			highestSleepGuard = g
		}
	}

	spew.Dump(highestSleepGuard)
}
