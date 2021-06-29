package main

// Importing fmt and time
import (
	"bufio"
	"fmt"
	"os"
)
import "time"

// Calling main
func main() {

	// Declaring t in UTC
	//t := time.Date(2020, 13, 34, 12, 45, 55, 0, time.UTC)
	startTime := time.Now()

	// Calling Clock method
	StHour, StMin, StSec := startTime.Clock()

	// Prints hours
	fmt.Printf("Pomodoro Start: %v:%v:%v\n", StHour, StMin, StSec)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {
			EndTime := time.Now()
			EndHour, EndMin, EndSec := EndTime.Clock()
			fmt.Printf("Pomodoro Start: %v:%v:%v\n", EndHour, EndMin, EndSec)
			break
		}
	}
}