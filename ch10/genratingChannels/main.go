package main

import (
	"math/rand"
	"time"
	"fmt"
)

func process2(ch chan int)  {
	n := rand.Intn(3000)
	time.Sleep(time.Duration(n)* time.Millisecond)
	ch <- n
}
// chan是同步的 通信机制
func main() {
	ch := make(chan int)
	go process2(ch)

	fmt.Println("Waiting for process")
	n := <- ch
	fmt.Printf("Process took %dms\n", n)
}
