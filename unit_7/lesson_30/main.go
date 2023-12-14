package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleepyGopher(id int, c chan int) {

	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	// fmt.Println(id)
	fmt.Println("...", id, "snore zzz ...")

	c <- id
}

func main() {

	c := make(chan int)

	// c <- 99

	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}

	timerCh := time.After(2 * time.Second)

	for i := 0; i < 5; i++ {

		select {
		case gopherID := <-c:
			fmt.Println("Gopher", gopherID, "Has finished sleeping")
		case <-timerCh:
			fmt.Println("Ran out of patience!!")
			return
		}

		// c <- "hello world"
		// msg := <- c
	}
	// time.Sleep(4 * time.Second)

	//> deadlock
  //* a deadlock is a condition where one or more goroutines are block because they are waiting for something that will never happen, like this next example
  //* main goroutine is trying to get a value from channel c but this operation does not have a corresponding send operation to that channel, that's why we cannot get the value from the channel, thus producing a fatal error:
  // ! fatal error: all goroutines are asleep - deadlock!
	// func main() {
	//  c := make(chan int)
	//  <-c
	// }

}
