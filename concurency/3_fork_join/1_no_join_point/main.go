package main

import (
	"fmt"
	"time"
)

func main(){
	go work() //fork join
	time.Sleep(100*time.Millisecond)
	fmt.Print("Done waiting, main exits!")

	//join point(but we're missing join point)
}

func work(){
	time.Sleep(500 * time.Millisecond)
	fmt.Println("printing some stuff")
}