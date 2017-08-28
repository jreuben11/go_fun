package main

import "time"
import "fmt"

func main() {
	// Timers represent a single event in the future. - provides a channel that will be notified
	timer1 := time.NewTimer(time.Second * 2)

	// <-timer1.C blocks until timer expired
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// unlike time.Sleep can cancel timer before it expires
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
