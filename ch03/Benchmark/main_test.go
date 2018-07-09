package main

import (
	"testing"
	"fmt"
)

func TestMerge(t *testing.T) {
	c := merge(asChan(1, 2, 3), asChan(4, 5, 6), asChan(7, 8, 9))
	seen := make(map[int]bool)
	for v := range c {
		if seen[v] {
			t.Errorf("saw %d at least twice", v)
		}
		seen[v] = true
	}
	for i := 1; i <= 9; i++ {
		if !seen[i] {
			t.Errorf("didn't see %d", i)
		}
	}
}

func BenchmarkMerge(b *testing.B) {
	for n := 1; n <= 1024; n *= 2 {
		chans := make([]<-chan int, n)
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for i:=range chans{
					chans[i] = asChan(0,1,2,3,4,5,6,7,8,9)
				}
				c := merge(chans...)
				for range c {
				}
			}
		})
	}
}
