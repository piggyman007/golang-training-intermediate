package main

import (
	"fmt"
	"math/rand"
	"time"
)

func saying(word string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("you said: %s : %d\n", word, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case y := <-input2:
				c <- y
			case <-time.After(300 * time.Millisecond):
				fmt.Println("too slow")
				close(c) // close channel
				return
			}
		}
	}()

	return c
}

func main() {
	book := saying("book fair")
	fon := saying("fon")

	c := fanIn(book, fon)

	for i := 0; i < 10; i++ {
		// ok is whether the channel is closed?
		if data, ok := <-c; ok {
			fmt.Println("print => ", data)
		}
	}

	fmt.Println("OK")
}
