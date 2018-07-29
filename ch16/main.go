package main

import (
	"log"
	"os"
	"fmt"
	"bufio"
)

func main() {
	xys, err := readData("ch16/data.txt")
	if err != nil{
		log.Fatalf("could not read data.txt: %v", err)
	}
	_ = xys
}

type xy struct {
	x, y float64
}

func readData(path string) ([]xy, error) {
	f, err := os.Open(path)
	if err != nil{
		return nil, err
	}

	defer f.Close()
	var xys []xy
	s := bufio.NewScanner(f)
	for s.Scan(){
		fmt.Println(s.Text())
	}
	if err := s.Err(); err != nil{
		return nil, fmt.Errorf("could not scan: %v", )
	}
	return xys, nil
}
