// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
	"time"
)

var w sync.WaitGroup

func ping(ch1 chan int, ch2 chan int) {
	<-ch1
	fmt.Println("Ping")
	ch2 <- 1
}
func pong(ch1 chan int, ch2 chan int) {
	<-ch2
	fmt.Println("Pong")
	ch1 <- 1
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	for i := 0; i < 10; i++ {
		go ping(ch1, ch2)
		go pong(ch1, ch2)
	}
	ch1 <- 1
	time.Sleep(time.Second * 5)
}
