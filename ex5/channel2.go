package main

import (
	"fmt"
	"time"
)

func helloWithChannel(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true

}
func channel2() {
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go helloWithChannel(done)
	<-done
	fmt.Println("Main received data")
}
