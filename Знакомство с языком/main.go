package main

import (
	"fmt"
	"time"
)

func main() {
	var current_time time.Time = time.Now()
	fmt.Println(current_time.Format("02:01:2006") + " " + current_time.Format(time.Kitchen))
}
