package main

import (
	"fmt"
	"time"
)

func main(){
	var now = time.Now()
	done := make(chan struct{})
	//fork-join
	go func(){
		work()
		done <- struct{}{}
	}()

	//join point
	<- done
	fmt.Println("elapsed: ", time.Since(now))
	fmt.Print("Done waiting, main exits!")
}

func work(){
	time.Sleep(500 * time.Millisecond)
	fmt.Println("printing some stuff")
}