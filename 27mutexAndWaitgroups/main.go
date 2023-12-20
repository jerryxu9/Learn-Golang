package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Learn Race Condition")

	wg := &sync.WaitGroup{}
	var score = []int{}
	mut := &sync.Mutex{}

	wg.Add(3)

	// Craete a couple of anonymous functions and use goroutine
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("One routine")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Two routine")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Three routine")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}
