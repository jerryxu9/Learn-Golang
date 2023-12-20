package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Learn channels in Golang")

	myChannel := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// // This will give an error: all goroutines are asleep - deadlock!
	// myChannel <- 5 // push 5 to the channel
	// fmt.Println(<-myChannel)

	wg.Add(2)

	// Receive only channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-ch
		fmt.Println(val)
		fmt.Println(isChannelOpen)
		// fmt.Println(<-myChannel) // print out the value in the channel
		wg.Done()
	}(myChannel, wg)

	// Send only channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		myChannel <- 5   // send 5 to the channel
		close(myChannel) // close the channel
		// myChannel <- 6
		wg.Done()
	}(myChannel, wg)

	wg.Wait()
}
