package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait() //deadlock  due to the main method never done cause of missing wg.done
}