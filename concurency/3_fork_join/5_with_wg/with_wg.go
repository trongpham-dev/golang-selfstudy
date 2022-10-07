package main

import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(3) // if we dont have add -> main does not wait until other tasks to be done!

	go func ()  {
		defer wg.Done()
		fmt.Println("1")
	}()

	go func ()  {
		defer wg.Done()
		fmt.Println("2")
	}()

	go func ()  {
		defer wg.Done()
		fmt.Println("3")
	}()

	wg.Wait()
}