package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	SHour, SMin, SSec := startTime.Clock()
	fmt.Printf("goPomo Start: %v:%v:%v\n", SHour, SMin, SSec)
	timeToRunning := 25
	fmt.Printf("Pomodoro Running: %dmin\n", timeToRunning)

	i := 0
	for range time.Tick(1 * time.Second) {
		if i == 1500 {
			break
		}
		i++
	}

	endTime := time.Now()
	EHour, EMin, ESec := endTime.Clock()
	fmt.Printf("goPomo Stop: %v:%v:%v\n", EHour, EMin, ESec)
}
