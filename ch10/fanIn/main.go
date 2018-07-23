package main

import (
	"math/rand"
	"fmt"
)

func producer(ch chan<- int, name string)  {
	for {
		n := rand.Intn(100)
		fmt.Printf("Channel %s -> %d\n", name, n)
		ch <- n
	}
}

func consumer(ch <- chan int)  {
	for n := range ch {
		fmt.Printf("<- %d\n", n)
	}
}

// 利用select来连接 produce和consumer
func fanIn(chA, chB <- chan int, chC chan<- int)  {
	var n int
	for{
		select {
		case n=<-chA:
			chC <- n
		case n=<-chB:
			chC <- n
		}
	}
}

func main() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	go producer(chA, "A")
	go producer(chB, "B")
	go consumer(chC)
	fanIn(chA, chB, chC)
}
