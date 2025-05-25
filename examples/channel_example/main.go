package main

import (
	"fmt"
	"sync"
)

func someFunction(ch chan int) {
	ch <- 5
}

func someFunction2(ch chan int) {
	ch <- 2
}

func someFunction3(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	x := <-ch
	fmt.Println("someFunction3", x)
}

func main() {
	ch := make(chan int, 1)
	var wg sync.WaitGroup
	wg.Add(1)

	go someFunction(ch)
	go someFunction2(ch)

	x := <-ch
	fmt.Println("main:", x)
	x = <-ch
	fmt.Println("main:", x)

	ch <- 10
	go someFunction3(ch, &wg)
	wg.Wait()
}
