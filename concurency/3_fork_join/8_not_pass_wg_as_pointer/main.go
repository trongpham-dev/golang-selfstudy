package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go work(wg)
	wg.Wait() //dead lock due to we pass a wg as a pointer
}

func work(wg sync.WaitGroup){
	defer wg.Done()
	fmt.Println("work is done")
}