package main

import (
	"fmt"
	"time"
)

func task_25() {
	fmt.Println("start")
	sleep(3 * time.Second)
	fmt.Println("end")
}

func sleep(duration time.Duration) {
	<-time.After(duration)
}
