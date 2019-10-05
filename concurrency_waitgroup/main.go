package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go saying("first", &wg)
	go saying("second", &wg)

	wg.Wait()
	fmt.Println("done!")
}

func saying(word string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("you said: %s: %d\n", word, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Microsecond)
	}
}
