package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var now = time.Now()
	var wg sync.WaitGroup
	wg.Add(1) // tell how many task that we need to wait
	//fork-join
	go func(){
		defer wg.Done()
		work()
	}()

	//join point
	wg.Wait()
	fmt.Println("elapsed: ", time.Since(now))
	fmt.Print("Done waiting, main exits!")
}

func work(){
	time.Sleep(500 * time.Millisecond)
	fmt.Println("printing some stuff")
}