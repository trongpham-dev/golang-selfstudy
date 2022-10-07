package main

import (
	"fmt"
	"time"
)

func main(){
	now := time.Now()
	/*the main exit due to something no-fork-join so we dont see any printed to the screen*/
	go task1()
	go task2()
	go task3()
	go task4()
	fmt.Println("elapsed:", time.Since(now))
}

func task1(){
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task 1")
}

func task2(){
	time.Sleep(200 * time.Millisecond)
	fmt.Println("task 2")
}

func task3(){
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task 3")
}

func task4(){
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task 4")
}
