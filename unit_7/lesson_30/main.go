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

	// c := make(chan int)

	// // c <- 99

	// for i := 0; i < 5; i++ {
	// 	go sleepyGopher(i, c)
	// }

	// timerCh := time.After(2 * time.Second)

	// for i := 0; i < 5; i++ {

	// 	select {
	// 	case gopherID := <-c:
	// 		fmt.Println("Gopher", gopherID, "Has finished sleeping")
	// 	case <-timerCh:
	// 		fmt.Println("Ran out of patience!!")
	// 		return
	// 	}

	// 	// c <- "hello world"
	// 	// msg := <- c
	// }
	// time.Sleep(4 * time.Second)

	//> deadlock
	//* a deadlock is a condition where one or more goroutines are block because they are waiting for something that will never happen, like this next example
	//* main goroutine is trying to get a value from channel c but this operation does not have a corresponding send operation to that channel, that's why we cannot get the value from the channel, thus producing a fatal error:
	// ! fatal error: all goroutines are asleep - deadlock!
	// func main() {
	//  c := make(chan int)
	//  <-c
	// }

	//? assembly line

	c0 := make(chan string)
	c1 := make(chan string)

	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)

	c2 := make(chan string)
	c3 := make(chan string)

	go source(c2)
	go removeIdentical(c2, c3)
	finalStage(c3)

	fmt.Println()
  
	c4 := make(chan string)
	c5 := make(chan string)
  
	go sourceGopher(c4)
	go splitSentence(c4, c5)
	printGopher(c5)
  
	fmt.Println()
	c6 := make(chan string)
	c7 := make(chan string)
	c8 := make(chan string)
	c9 := make(chan string)

	go sourceGopher(c6)
	go splitSentence(c6, c7)
	go removeIdentical(c7, c8)
	go filterGopher(c8, c9)
	printGopher(c9)

}
