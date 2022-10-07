package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Second)
		wg.Done()
		wg.Add(1) // -> dead lock due to we're reusing a wg that have not done it's task yet!
	}()
	wg.Wait()
}