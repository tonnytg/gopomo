package main

import (
	"flag"
	"fmt"
	"os/exec"
	"time"
)

type Pomodoro struct {
	Timer time.Time
	Loop  int
	Hour  int
	Min   int
	Sec   int
}

// Count is a Method to count loops
func (p *Pomodoro) Count() {
	p.Loop = p.Loop + 1
}

// GetTime obtain exactly moment time and archive at Type Pomodoro.
// Hour, Minute and Second have a format of exactly time
func GetTime(x *Pomodoro) (Hour, Min, Sec int) {
	x.Timer = time.Now()
	x.Hour, x.Min, x.Sec = x.Timer.Clock()
	return x.Hour, x.Min, x.Sec
}

// LogicPomodoro it is a Looping to use var PomoTime for ajust Pomodoro loop
// You can use 25 minutes or custom minutes do you want.
func LogicPomodoro(x int) error {
	for i := 0; i < x; i++ {
		time.Sleep(60 * time.Second)
	}
	return nil
}

var (
	x                Pomodoro
	PomodoroTime     int = 1
	PomodoroQuantity int = 1
)

func main() {

	// PomodoroTime it is a Default time of Pomodoro.
	// It is 25 minutes with 5 minutes with pause for a break.
	PomodoroTime := flag.Int("time", 25, "Custom Time")
	PomodoroQuantity := flag.Int("repeat", 3, "Pomodoro Quantity")

	// Check some Args different of default
	flag.Parse()

	fmt.Println("Time:", *PomodoroTime)
	fmt.Println("Quantity:", *PomodoroQuantity)

	for i := 1; i <= *PomodoroQuantity; i++ {
		// GetTime Start
		GetTime(&x)
		exec.Command("say", "GoPomo start").Output()
		fmt.Printf("Start: %d:%d:%d Pomodoro: %d/%d\n", x.Hour, x.Min, x.Sec, x.Loop, *PomodoroQuantity)

		// Looping with time sleep
		err := LogicPomodoro(*PomodoroTime)
		if err != nil {
			fmt.Println("Error", err)
		}

		// GetTime End
		GetTime(&x)
		exec.Command("say", "GoPomo stop").Output()

		// Count loop + 1
		x.Count()
		fmt.Printf("End: %d:%d:%d Pomodoro: %d/%d\n", x.Hour, x.Min, x.Sec, x.Loop, *PomodoroQuantity)

		switch {
		case i%5 == 0:
			fmt.Printf("Take 15 minutes for a cofee!\nPress Enter to continue!\n")
			fmt.Scanln()
		default:
			fmt.Printf("Take 5 minutes for a cofee!\nPress Enter to continue!\n")
			fmt.Scanln()
		}
	}
}
