package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Learning time...")
	presentDay := time.Now()
	fmt.Println("Learning time...", presentDay.Format("01-02-2006 15:04:05 Monday")) // This has to be 01-02-2006 15:04:05 Monday
}
