package main

import (
	"time"
	"math/rand"
	"fmt"
)

func Worker(in, out chan int)  {
	for {
		n := <-in
		time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
		out <- n
	}
}

func Task(ch chan <- int)  {
	i := 0
	for{
		fmt.Printf("-> Send job: %d\n", i)
		ch <- i
		i++
	}
}

func main() {
	in := make(chan int)
	out := make(chan int)
	// 手动启动work4个
	for i:=0;i<4;i++  {
		go Worker(in, out)
	}

	go Task(in)
	for n := range out {
		fmt.Printf("<- Recv job: %d\n", n)
	}
}
