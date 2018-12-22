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
	minutes      map[int]int
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
			if onduty.minutes == nil {
				onduty.minutes = make(map[int]int)
			}
			onduty.sleepingTime += input.t.Minute() - onduty.asleepAt

			for i := onduty.asleepAt; i < input.t.Minute(); i++ {
				onduty.minutes[i]++
			}

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
	var highestMinute int
	var highestMinuteCount int
	for _, g := range guards {
		for minute, c := range g.minutes {
			if c > highestMinuteCount {
				highestSleepGuard = g
				highestMinute = minute
				highestMinuteCount = c
			}
		}
	}

	spew.Dump(highestSleepGuard)
	spew.Dump(highestMinute)
}
