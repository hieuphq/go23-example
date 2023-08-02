package main

import (
	"fmt"
	"time"
)

type response struct {
	Body string
	Err  error
}

func crawlData(rsCh chan response) {
	time.Sleep(250 * time.Millisecond)
	rsCh <- response{
		Body: fmt.Sprintf("%v", time.Now().UnixMilli()),
		Err:  nil,
	}
}

func processData(rsCh chan response) {
	for rs := range rsCh {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("processed", rs)
	}
}

func channel1() {
	var a chan response
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan response, 2)
		fmt.Printf("Type of a is %T", a)

		for i := 0; i < 4; i++ {
			go crawlData(a)
		}

		// Case 1
		// rs := <-a
		// fmt.Println(rs)

		// rs = <-a
		// fmt.Println(rs)

		// Case 2
		// for rs := range a {
		// 	fmt.Println(rs)
		// }

		// Case 3

		go processData(a)

		time.Sleep(1 * time.Minute)
	}

}
