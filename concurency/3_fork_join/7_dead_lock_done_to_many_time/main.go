package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Done() // deadlock due to we do not have any action to wg..
}