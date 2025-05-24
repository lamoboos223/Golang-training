package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func printLetters() {
	for i := 'a'; i < 'j'; i++ {
		fmt.Println(string(i))
	}
}

func main() {
	done := make(chan bool)

	go func() {
		printNumbers()
		done <- true // Signal that numbers are done
	}()

	go func() {
		<-done // Wait for numbers to finish
		printLetters()
	}()

	time.Sleep(time.Second)
}
