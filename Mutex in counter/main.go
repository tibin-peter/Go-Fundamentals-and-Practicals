package main

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // adding morethan one goroutes

	go increment(&wg)
	go increment(&wg)

	wg.Wait() // wait untill the both exicution to finish
	fmt.Println("Final counter value:", counter)

}
