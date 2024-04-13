package Redis

import (
	"fmt"
	"time"
)

func PerformIn(in time.Duration, task string) {
	at := time.Now().Add(in)
	PerformAt(at, task)
}

func PerformAt(at time.Time, task string) {
	fmt.Printf("> Scheduling SMTH \n")
}
