package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st gorutine sleeping")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd gorutine sleeping...")
		time.Sleep(2)
	}()

	wg.Wait()
	fmt.Println("All gorutine complete.")
}
