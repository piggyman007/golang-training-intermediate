package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++

		fmt.Println(name, ball.hits)
		time.Sleep(500 * time.Millisecond)

		table <- ball

		if name == "fon" {
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	table := make(chan *Ball)
	go player("fon", table)
	go player("se", table)

	table <- &Ball{}

	time.Sleep(10 * time.Second)
	<-table // remove ball from table

}
